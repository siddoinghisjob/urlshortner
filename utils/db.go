package utils

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
)

func (err *notFoundError) Error() string {
	return fmt.Sprintf("error : %s", err.message)
}

func (err *internalError) Error() string {
	return fmt.Sprintf("error : %s", err.message)
}

const letters string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func hashHelper(h *[]byte) {
	for i := range *h {
		(*h)[i] = letters[rand.Intn(len(letters))]
	}
}

func hash(h int) string {
	hashedVal1 := make([]byte, 2)
	hashedVal2 := make([]byte, 2)
	hashHelper(&hashedVal1)
	hashHelper(&hashedVal2)
	id := strconv.Itoa(h)
	hashedVal := fmt.Sprintf("%s%s%s", string(hashedVal1), id, string(hashedVal2))
	return hashedVal
}

func setIntoPg(link string) (string, error) {
	var surl string
	var id int
	err := DB.QueryRow("SELECT id FROM urls WHERE link = $1", link).Scan(&id)
	if err == sql.ErrNoRows {
		err = DB.QueryRow("INSERT INTO urls (link) VALUES ($1) returning id", link).Scan(&id)
		if err != nil {
			fmt.Print("Error 43")
			return "", &internalError{}
		}
		surl = hash(id)

		_, err = DB.Exec("INSERT INTO surls (uid, surl) VALUES ($1,$2)", id, surl)

		if err != nil {
			fmt.Print("Error 51")
			return "", &internalError{}
		}
		return surl, nil
	}
	if err != nil {
		fmt.Print("Error 57")
		return "", &internalError{}
	}
	err = DB.QueryRow("SELECT surl FROM surls WHERE uid = $1", id).Scan(&surl)
	if err != nil {
		fmt.Print("Error 62")
		return "", &internalError{}
	}
	return surl, nil
}

func getFromPg(link string) (string, error) {
	var id int
	err := DB.QueryRow("SELECT id FROM urls WHERE link = $1", link).Scan(&id)
	if err == sql.ErrNoRows {
		return "", &notFoundError{"Not Found"}
	}
	if err != nil {
		return "", &internalError{}
	}
	var surl string
	err = DB.QueryRow("SELECT surl FROM surls WHERE uid = $1", id).Scan(&surl)
	if err == nil {
		return surl, nil
	}
	return "", &internalError{"Internal Error"}
}

func setClicksInPG(url string, name string, date string) (string, error) {
	cid, err := setCountryInPG(name)

	if err != nil {
		return "", &internalError{}
	}
	var uid int
	var surl string

	err = DB.QueryRow("SELECT id FROM urls WHERE link = $1", url).Scan(&uid)
	if err == sql.ErrNoRows {
		return "", nil
	}

	if err != nil {
		return "", &internalError{}
	}
	err = DB.QueryRow("SELECT surl FROM surls WHERE uid = $1", uid).Scan(&surl)

	if err != nil {
		return "", &internalError{}
	}

	_, err = DB.Exec("INSERT INTO clicks (uid, cid, date) VALUES ($1,$2,$3) returning id", uid, cid, date)

	if err != nil {
		return "", &internalError{}
	}

	return surl, nil
}

func setCountryInPG(name string) (int8, error) {
	var id int8
	err := DB.QueryRow("Select id from country where name = $1", name).Scan(&id)
	if err == nil {
		return id, nil
	}

	if err != sql.ErrNoRows {
		return 0, &internalError{}
	}

	err = DB.QueryRow("INSERT INTO country (name) VALUES ($1) returning id", name).Scan(&id)

	if err != nil {
		return 0, &internalError{}
	}

	return id, nil
}

func getTotalVis(surl string) (int, int) {
	var totalVisitors, id int

	err := DB.QueryRow("SELECT uid FROM surls WHERE surl = $1", surl).Scan(&id)
	if err != nil {
		return -1, -1
	}

	err = DB.QueryRow("SELECT COUNT(*) FROM clicks WHERE uid = $1", id).Scan(&totalVisitors)
	if err != nil {
		return -1, -1
	}
	return totalVisitors, id
}

func getTotalCountry(id int) (countryData []analyticsData) {
	rows, err := DB.Query("SELECT c.name, COUNT(v.uid) AS visitors_count FROM clicks v JOIN country c ON v.cid = c.id WHERE uid = $1 GROUP BY c.name ORDER BY visitors_count DESC LIMIT 3", id)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var countryName string
		var visitorsCount int
		err = rows.Scan(&countryName, &visitorsCount)
		if err != nil {
			return
		}
		countryData = append(countryData, analyticsData{Name: countryName, Data: visitorsCount})
	}
	return
}

func getTotalDate(id int) (dateData []analyticsData) {
	rows, err := DB.Query("SELECT date, COUNT(*) AS visitors_count FROM clicks WHERE uid = $1 GROUP BY date ORDER BY visitors_count DESC LIMIT 5", id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var date string
		var visitorsCount int
		err = rows.Scan(&date, &visitorsCount)
		if err != nil {
			return
		}
		date = date[:len(date)-10]
		dateData = append(dateData, analyticsData{Name: date, Data: visitorsCount})
	}
	return
}
