package handler

import (
	"github.com/gin-contrib/cors"
)

func (h *Handler) setupCors() {
	h.gin.Use(cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type"},
		AllowOrigins: h.corsOrigins,
	}))
}
