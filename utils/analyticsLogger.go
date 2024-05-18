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
	fmt.Println(date)
	var uid int
	err := l.db.QueryRow("SELECT id from urls where link = $1", url).Scan(&uid)
	if err != nil {
		fmt.Println(err)
		return
	}
	l.buffer = append(l.buffer, analyticsStore{uid: uid, name: name, date: date})
}

func (l *analyticsLogger) flush() {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	for _, v := range l.buffer {
		err := setClicksInPG(int8(v.uid), v.name, v.date)
		if err != nil {
			log.Fatal(fmt.Sprintln("Error registering data."))
		}
	}

	l.buffer = nil

}
