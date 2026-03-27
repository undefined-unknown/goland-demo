package controller

import "net/http"

type TemplateController struct{}

func NewTemplateController() *TemplateController {
	return &TemplateController{}
}

func (c *TemplateController) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("template list"))
}

func (c *TemplateController) Favorite(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("template favorite"))
}