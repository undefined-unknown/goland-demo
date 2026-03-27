package controller

import (
	"encoding/json"
	"net/http"
	"template-backend/internal/service"
)

type TemplateController struct{
	service *service.TemplateService
}

func NewTemplateController() *TemplateController {
	return &TemplateController{
		service: service.NewTemplateService(),
	}
}

func (c *TemplateController) List(w http.ResponseWriter, r *http.Request) {
	data := c.service.GetTemplateList()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (c *TemplateController) Favorite(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("template favorite"))
}