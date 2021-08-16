package service

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HomeService struct {
	log *zap.Logger
}

func NewHome(log *zap.Logger) *HomeService {
	return &HomeService{
		log: log,
	}
}

func (h *HomeService) Index(ctx *gin.Context) {
	h.log.Info("welcome Home page")
}
