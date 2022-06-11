package ws

import (
	"context"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	WriteBufferSize: 1024,
	ReadBufferSize:  1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Pool struct {
	conn *websocket.Conn
	data chan string
}

func (receiver *Pool) Upgrade(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	receiver.conn = conn
}

func (receiver *Pool) Send() chan<- string {
	if receiver.data == nil {
		receiver.data = make(chan string)
	}
	return receiver.data
}

func (receiver *Pool) Run(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case d, ok := <-receiver.data:
				if !ok {
					return
				}
				if receiver.conn != nil {
					receiver.conn.WriteMessage(websocket.TextMessage, []byte(d))
				}
			}
		}
	}()
}

func New(ctx context.Context) *Pool {
	p := Pool{}
	p.Run(ctx)
	return &p
}
