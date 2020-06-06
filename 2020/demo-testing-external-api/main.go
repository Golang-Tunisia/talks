package main

import (
	"Meetup/api"
	"Meetup/api/external_api"
	"Meetup/vars/api_vars"
	"encoding/json"
	"github.com/flannel-dev-lab/cyclops"
	"github.com/flannel-dev-lab/cyclops/response"
	"github.com/flannel-dev-lab/cyclops/router"
	"net/http"
)

type Handler struct {
	Api api.ExternalApi
}

func main() {
	api := external_api.New("http://localhost:8080")
	handler := Handler{Api: api}

	routes := router.New()

	routes.Post("/user", handler.CreateUser)

	cyclops.StartServer(":80", routes)

}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user api_vars.User

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.ErrorResponse(http.StatusBadRequest, err, err.Error(), w, false, nil)
		return
	}

	id, err := h.Api.CreateUser(user)
	if err != nil {
		response.ErrorResponse(http.StatusUnprocessableEntity, err, err.Error(), w, false, nil)
		return
	}

	response.SuccessResponse(http.StatusOK, w, map[string]string{"id": id})
	return
}
