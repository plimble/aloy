package mw

import (
	"log"

	"github.com/plimble/aloy/services/aloy"
	"github.com/plimble/errors"
	"gopkg.in/kataras/iris.v6"
)

type handlerFunc func(*iris.Context) error
type HMWFunc func(*iris.Context) error
type HFunc func(*iris.Context, aloy.UsecaseInterface) error

func HMW(h handlerFunc) iris.HandlerFunc {
	return func(ctx *iris.Context) {
		if err := h(ctx); err != nil {
			status, werr := errors.ErrorStatus(err)
			if status == 500 {
				log.Printf("%v", err)
			}

			ctx.JSON(status, iris.Map{
				"error": iris.Map{
					"code":    status,
					"message": werr.Error(),
				},
			})
		}
	}
}

func H(h HFunc, uc aloy.UsecaseInterface) iris.HandlerFunc {
	return func(ctx *iris.Context) {
		if err := h(ctx, uc); err != nil {
			status, werr := errors.ErrorStatus(err)
			if status == 500 {
				log.Printf("%+v", err)
			}

			ctx.JSON(status, iris.Map{
				"error": iris.Map{
					"code":    status,
					"message": werr.Error(),
				},
			})
		}
	}
}
