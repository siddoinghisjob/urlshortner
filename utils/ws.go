package utils

import (
	"encoding/json"
	"log"
	"net/http"
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
	list map[string][]*websocket.Conn
}{list: make(map[string][]*websocket.Conn)}

func sendWS(v []byte, ws []*websocket.Conn) {
	connections.RLock()
	defer connections.RUnlock()
	for _, w := range ws {
		if err := w.WriteMessage(websocket.TextMessage, v); err != nil {
			log.Println("WriteMessage error on initial data:", err)
			w.Close()
			go removeConn(w)
		}
	}
}

func sendData(url string, ws []*websocket.Conn) {
	if ws == nil {
		tmp, ok := connections.list[url]
		if !ok {
			return
		}
		ws = tmp
	}
	totalVisitors, id := getTotalVis(url)
	if totalVisitors == -1 {
		mp := make(map[string]bool)
		mp["error"] = true
		jdata, err := json.Marshal(mp)
		if err != nil {
			return
		}
		sendWS(jdata, ws)
		return
	}

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

	sendWS(jdata, ws)
}

func Analytics(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	url := c.Param("id")

	go func() {
		sendData(url, []*websocket.Conn{ws})
	}()

	connections.Lock()
	connections.list[url] = append(connections.list[url], ws)
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
