package telegram

import (
	"context"
	"fmt"
	"github.com/farwydi/stickfinbot/pkg/domain"
	"github.com/farwydi/stickfinbot/pkg/usecase/general"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"os"
	"time"
)

func NewTelegramEndpoint(gapp *general.General) domain.Endpoint {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()

	r.POST("/hook", func(c *gin.Context) {
		var update tgbotapi.Update
		if err := c.ShouldBindJSON(&update); err != nil {
			err := c.Error(err)
			c.JSON(http.StatusBadRequest, err.JSON())
			return
		}

		ctx, cancel := context.WithTimeout(c.Request.Context(), 8*time.Second)
		defer cancel()

		if err := gapp.Proc(ctx, toUpdate(update)); err != nil {
			err := c.Error(err)
			c.JSON(http.StatusInternalServerError, err.JSON())
			return
		}

		c.Status(http.StatusOK)
	})

	return &tgEndpoint{
		srv: &http.Server{
			Addr:    fmt.Sprintf(":%s", port),
			Handler: r,
		},
	}
}

type tgEndpoint struct {
	srv *http.Server
}

func (ep *tgEndpoint) Run() error {
	// service connections
	if err := ep.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (ep *tgEndpoint) Stop() {
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := ep.srv.Shutdown(ctx); err != nil {
		log.Printf("Fail shutdown: %v", err)
		return
	}
}
