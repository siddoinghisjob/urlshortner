package utils

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var aLogger *analyticsLogger

func Logging(ticker *time.Ticker) {
	aLogger = setupLogger(DB)

	go func() {
		for {
			aLogger.flush()
			<-ticker.C
		}
	}()
}

func Post(c *gin.Context) {
	var url urlLink
	if err := c.BindJSON(&url); err != nil {
		c.IndentedJSON(http.StatusBadRequest, message{Text: "Format not correct.", Error: true})
		return
	}
	surl, err := storeInDB(url.URL)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, message{Text: "Error. Could not process.", Error: true})
		return
	}
	c.IndentedJSON(http.StatusCreated, message{Text: surl, Error: false})

}

func Get(c *gin.Context) {
	url, err := searchFromDB(c.Param("id"))

	countryName := c.DefaultQuery("c", "NA")

	if _, ok := err.(*internalError); ok {
		c.IndentedJSON(http.StatusInternalServerError, message{Text: "Internal Error", Error: true})
		return
	}
	if _, ok := err.(*notFoundError); ok {
		c.IndentedJSON(http.StatusNotFound, message{Text: url, Error: true})
		return
	}
	y, m, d := time.Now().Date()
	aLogger.store(url, countryName, fmt.Sprintf("%s-%s-%s", strconv.Itoa(y), m, strconv.Itoa(d)))
	c.IndentedJSON(http.StatusOK, message{Text: url, Error: false})
}

func getId(url string) int {
	re := regexp.MustCompile(`[A-Za-z]{2}(\d+)[A-Za-z]{2}`)

	matches := re.FindStringSubmatch(url)

	if len(matches) > 0 {
		integerStr := matches[1]

		number, err := strconv.Atoi(integerStr)
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			return 0
		}

		return number
	}

	return 0
}

func Analytics(c *gin.Context) {
	id := getId(c.Param("id"))

	totalVisitors := getTotalVis(id)
	countryData := getTotalCountry(id)
	dateData := getTotalDate(id)

	c.JSON(http.StatusOK, gin.H{
		"total_visitors": totalVisitors,
		"country_data":   countryData,
		"date_data":      dateData,
	})
}

func searchFromDB(surl string) (string, error) {
	link, err := getFromRedis(surl)
	_, notFoundErr := err.(*notFoundError)
	_, internalErr := err.(*internalError)
	if notFoundErr || internalErr {
		link, err := getFromPg(surl)
		if _, ok := err.(*notFoundError); ok {
			return link, &notFoundError{}
		} else if _, ok := err.(*internalError); ok {
			return link, &internalError{}
		}
		return link, nil
	} else if err == nil {
		return link, nil
	}
	return link, &internalError{}
}

func storeInDB(link string) (string, error) {
	surl, err := getFromRedis(link)

	if err == nil {
		return surl, nil
	}

	if _, ok := err.(*notFoundError); ok {
		surl, err := setIntoPg(link)
		if err != nil {
			return "", &internalError{}
		}
		err = setIntoRedis(link, surl)
		if err != nil {
			return "", &internalError{}
		}
		return surl, nil
	}

	return "", &internalError{}
}
