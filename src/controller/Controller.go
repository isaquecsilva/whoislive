package controller

import (
	"html/template"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/isaquecsilva/whoislive/src/service"
	util "github.com/isaquecsilva/whoislive/src/util/template"
	"github.com/isaquecsilva/whoislive/src/view"
)

type Controller struct {
	service service.LiveServiceInterface
}

func (c *Controller) LiveChannels(ctx *fiber.Ctx) error {
	channels, err := c.service.GetLiveChannels()
	if err != nil {
		ctx.Status(http.StatusInternalServerError).Set("Content-Type", "text/plain")
		return ctx.SendString("An error has ocurred on the server. Please, try again later. \n\nError Message: " + err.Error())
	}

	ctx.Status(http.StatusOK).Set("Content-Type", "text/html")
	tmpl, err := template.ParseFiles(util.GetTemplatesFilePath("livestreamchannels.templ"))
	return tmpl.Execute(ctx.Response().BodyWriter(), view.NewChannelsView("Live Channels", channels))
}

func NewController(service service.LiveServiceInterface) *Controller {
	return &Controller{
		service,
	}
}