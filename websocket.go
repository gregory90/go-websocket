package websocket

import (
	"github.com/googollee/go-socket.io"
	"gopkg.in/gin-gonic/gin.v1"

	. "github.com/gregory90/go-webutils/logger"
)

var (
	Server      *socketio.Server
	connections map[socketio.Socket]bool
)

func GinHandler(c *gin.Context) {

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

	Server.ServeHTTP(c.Writer, c.Request)
}

func Init() {
	var err error
	Server, err = socketio.NewServer(nil)

	connections = make(map[socketio.Socket]bool)

	if err != nil {
		Log.Info("Socket.io fatal: %+v", err)
	}
}

func SendMessage(msg string) {
	for conn, _ := range connections {
		conn.Emit("app", msg)
	}
}
