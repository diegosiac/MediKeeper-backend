package medicine

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Handler struct {
	service *Service
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		service: NewService(db),
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
}
