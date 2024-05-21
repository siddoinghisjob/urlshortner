package utils

import (
	"database/sql"
	"fmt"
	"log"
)

func setupLogger(db *sql.DB) *analyticsLogger {
	return &analyticsLogger{
		db: db,
	}
}

func (l *analyticsLogger) store(url, name, date string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.buffer = append(l.buffer, analyticsStore{url: url, name: name, date: date})
}

func (l *analyticsLogger) flush() {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	for _, v := range l.buffer {
		surl, err := setClicksInPG(v.url, v.name, v.date)
		if err != nil {
			log.Fatal(fmt.Sprintln("Error registering data."))
			continue
		}

		sendData(surl, nil)
	}

	l.buffer = nil
}
