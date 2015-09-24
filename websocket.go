package websocket

import (
	"github.com/googollee/go-socket.io"

	. "bitbucket.org/pqstudio/go-webutils/logger"
)

var (
	Server      *socketio.Server
	connections map[socketio.Socket]bool
)

func Init() {
	var err error
	connections = make(map[socketio.Socket]bool)
	Server, err = socketio.NewServer(nil)

	if err != nil {
		Log.Info("Socket.io fatal: %+v", err)
	}

	Server.On("connection", func(so socketio.Socket) {
		connections[so] = true
		Log.Info("Socket.io connection")
		so.Join("app")

		so.On("disconnection", func() {
			delete(connections, so)
			Log.Info("Socket.io disconnection")
		})
	})

	Server.On("error", func(so socketio.Socket, err error) {
		Log.Info("Socket.io error: %+v", err)
	})
}

func SendMessage(msg string) {
	for conn, _ := range connections {
		conn.Emit("update", msg)
	}
}
