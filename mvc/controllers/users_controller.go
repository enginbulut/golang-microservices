package controllers

import (
	"encoding/json"
	"github.com/enginbulut/golang-microservices/mvc/services"
	"github.com/enginbulut/golang-microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userIdParam := req.URL.Query().Get("user_id")

	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue,_ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		resp.WriteHeader(apiErr.StatusCode)
		jsonValue,_ := json.Marshal(apiErr)
		resp.Write(jsonValue)
		//handle the error and return to client
		return
	}

	jsonValue,_ := json.Marshal(user)
	resp.Write(jsonValue)
}
