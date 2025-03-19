package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return r.Host == "localhost:8080"
	},
}

type Handler struct {
	logger *logrus.Logger
}

func NewHandler(logger *logrus.Logger) *Handler {
	return &Handler{logger: logger}
}

func (h *Handler) getMsg(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest,
			fmt.Errorf("couldn't make websocket connection: %w", err).Error(), h.logger)
		return
	}
	h.logger.Infof("New websocket connection established")
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest,
				fmt.Errorf("couldn't read message: %w", err).Error(), h.logger)
			return
		}
		h.logger.Infof("Received message: %s", msg)

		response := "server responses: received " + string(msg)

		err = conn.WriteMessage(websocket.TextMessage, []byte(response))
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError,
				fmt.Errorf("couldn't write message: %w", err).Error(), h.logger)
			return
		}
		h.logger.Infof("Sent response to client: %s", response)
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/send_message", h.getMsg)

	return router
}
