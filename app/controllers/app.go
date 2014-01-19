package controllers

import "github.com/robfig/revel"

type App struct {
	*revel.Controller
}

type StatusStruct struct {
	Message string `json:"message"`
}

func (c App) Index() revel.Result {
	return c.RenderJson(StatusStruct{"ok"})
}
