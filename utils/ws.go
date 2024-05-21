package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var connections = struct {
	sync.RWMutex
	list map[int][]*websocket.Conn
}{list: make(map[int][]*websocket.Conn)}

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

func sendData(id int, ws []*websocket.Conn) {
	if ws == nil {
		tmp, ok := connections.list[id]
		if !ok {
			return
		}
		ws = tmp
	}
	totalVisitors := getTotalVis(id)
	countryData := getTotalCountry(id)
	dateData := getTotalDate(id)

	data := make(map[string]any)
	data["total"] = totalVisitors
	data["country"] = countryData
	data["date"] = dateData

	jdata, err := json.Marshal(data)
	if err != nil {
		return
	}
	for _, w := range ws {
		if err := w.WriteMessage(websocket.TextMessage, jdata); err != nil {
			log.Println("WriteMessage error on initial data:", err)
			return
		}
	}
}

func Analytics(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	url := c.Param("id")
	id := getId(url)
	go func() {
		sendData(id, []*websocket.Conn{ws})
	}()

	connections.Lock()
	connections.list[id] = append(connections.list[id], ws)
	connections.Unlock()
}

func removeConn(curr *websocket.Conn) {
	connections.Lock()
	defer connections.Unlock()
	for id, conn := range connections.list {
		var flg bool = false
		for i, ws := range conn {
			if ws == curr {
				conn = append(conn[:i], conn[i+1:]...)
				flg = true
				break
			}
		}
		if flg {
			if len(conn) == 0 {
				delete(connections.list, id)
			}
			break
		}
	}
}
