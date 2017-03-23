package api

import (
	"log"

	"github.com/plimble/errors"
	"gopkg.in/kataras/iris.v6"
)

type HFunc func(*iris.Context) error

func W(h HFunc) iris.HandlerFunc {
	return func(ctx *iris.Context) {
		if err := h(ctx); err != nil {
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
		} else {
			ctx.SetStatusCode(200)
		}
	}
}
