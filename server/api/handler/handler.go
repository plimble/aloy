package handler

import (
	"github.com/plimble/aloy/aloy"
)

type Handler struct {
	aloy aloy.Aloy
}

func New(aloy aloy.Aloy) *Handler {
	return &Handler{aloy}
}
