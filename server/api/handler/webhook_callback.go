package handler

import (
	"io/ioutil"

	"github.com/plimble/errors"
	"gopkg.in/kataras/iris.v6"
)

func (h *Handler) WebhookCallback(ctx *iris.Context) error {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return errors.WithStack(err)
	}

	err = h.aloy.WebhookCallback(ctx.Request.Host, body)

	return errors.WithStack(err)
}
