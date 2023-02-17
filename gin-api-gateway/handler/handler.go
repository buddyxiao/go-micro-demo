package handler

import (
	rand "rand-service/proto"
	user "user-service/proto"
)

type APIHandler struct {
	randClient rand.RandService
	userClient user.UserService
}

func GetAPIHandler(rand rand.RandService, user user.UserService) *APIHandler {
	return &APIHandler{
		randClient: rand,
		userClient: user,
	}
}
