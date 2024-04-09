package routes

import (
	"encoding/json"
	"fill-asset-process-prototype/models"
	"fill-asset-process-prototype/services"
	"net/http"
)

func HandlerUserGet(w http.ResponseWriter, r *http.Request) {
	users := services.GetUsers()
	isSave := services.SaveUsersCache(users)
	resp := models.Response{}
	if isSave {
		resp = models.Response{
			Code:    200,
			Message: "Saved successfully",
		}
	} else {
		resp = models.Response{
			Code:    200,
			Message: "Unsaved",
		}
	}

	json.NewEncoder(w).Encode(&resp)
}
