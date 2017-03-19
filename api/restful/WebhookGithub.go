package restful

import (
	"io/ioutil"

	"github.com/plimble/aloy/aloy"
	"github.com/plimble/errors"
	"gopkg.in/kataras/iris.v6"
)

func WebhookGithub(ctx *iris.Context, uc aloy.UsecaseInterface) error {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return errors.WithStack(err)
	}

	if err = uc.GetGithubWebhook(&aloy.GetGithubWebhookInput{Payload: body}); err != nil {
		ctx.SetStatusCode(500)
	} else {
		ctx.SetStatusCode(200)
	}

	return err
}
