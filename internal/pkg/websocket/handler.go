package websocket

import (
	"net/http"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/gorilla/websocket"
)

type Websocket struct {
	upgrader *websocket.Upgrader
	log      *log.Helper
}

func (w *Websocket) Handler(writer http.ResponseWriter, request *http.Request) {

	conn, err := w.upgrader.Upgrade(writer, request, nil)
	if err != nil {
		w.log.Error("upgrade:", err)
	}
	defer conn.Close()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			w.log.Error("read message:", err)
			break
		}
		w.log.Debugf("recv: %s", message)

		msg := string(message)
		switch msg {
		case "PING":
		default:
		}
		err = conn.WriteMessage(mt, message)
		if err != nil {
			w.log.Error("write:", err)
			break
		}
	}

	if err != nil {
		w.log.Error("close:", err)
	}
}

func NewWebsocket(logger log.Logger) *Websocket {
	return &Websocket{
		upgrader: &websocket.Upgrader{},
		log:      log.NewHelper(logger),
	}
}

var ProviderSet = wire.NewSet(NewWebsocket)
