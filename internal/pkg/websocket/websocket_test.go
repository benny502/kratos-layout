package websocket

import (
	"fmt"
	"net/http"
	"testing"

	wb "github.com/gorilla/websocket"
)

func TestWebsocketClient(t *testing.T) {
	dialer := wb.Dialer{}
	header := http.Header{}
	header.Add("Cookie", "vhash=49198581d9e1ca4427cadad14799176d; token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjI0NDMxNDMsInN1YiI6InRlc3QiLCJuYmYiOjE2NjIzNzExNDMsImF1ZCI6ImV2ZXJ5IiwiaWF0IjoxNjYyMzcxMTQzLCJqdGkiOiIyOWU0NmViNDczNTVjNzQzNDkwMDk3YWRhYmI4MTcwYSIsImlzcyI6ImZsZXhiaSIsInN0YXR1cyI6MSwiZGF0YSI6eyJpc0dvZCI6MCwidWlkIjo0fX0.0LMVps5Q198Gns2yYyXp7fofjIYqZTc__HCT_y0MaYY")
	connect, _, err := dialer.Dial("wss://flexbi.hylink.com/websocket/", header)

	if err != nil {
		t.Log(err)
		return
	}
	defer connect.Close()

	err = connect.WriteMessage(wb.TextMessage, []byte("PING"))
	if err != nil {
		t.Log(err)
	}

	messageType, messageData, err := connect.ReadMessage()
	if err != nil {
		t.Log(err)
	}

	switch messageType {
	case wb.TextMessage:
		fmt.Println(string(messageData))
	}

}
