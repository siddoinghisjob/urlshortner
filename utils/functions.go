package utils

import (
	"fmt"
	"net/http"
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
	var uLink urlLink
	if err := c.BindJSON(&uLink); err != nil {
		c.IndentedJSON(http.StatusBadRequest, message{Text: "Format not correct.", Error: true})
		return
	}

	surl, err := storeInDB(uLink.URL)
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

func searchFromDB(surl string) (string, error) {
	link, err := getFromRedis(surl)
	_, notFoundErr := err.(*notFoundError)
	_, internalErr := err.(*internalError)
	if notFoundErr || internalErr {
		link, err := getFromPg(surl)
		if _, ok := err.(*notFoundError); ok {
			return link, &notFoundError{message: "not found"}
		} else if _, ok := err.(*internalError); ok {
			return link, &internalError{message: "internal error"}
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

func Healthz(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"success": true})
}
