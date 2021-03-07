package user

import (
	"demo_server/api"
	"fmt"
	"log"
	"net/http"
)

func Greet() http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		log.Printf("req body in greet handler %v", *req)
		userID := req.FormValue("user_id")
		log.Printf("userID : %v", userID)
		if userID != "user_id_value" {
			api.Error(400, api.Response{Message: "invalid request"}, rw)
			return
		}

		name := req.FormValue("name")
		if name == "" {
			api.Error(400, api.Response{Message: "please provide name"}, rw)
			return
		}

		message := fmt.Sprintf("hello %s", name)
		api.Success(http.StatusCreated, api.Response{Message: message}, rw)
	})
}

func Welcome() http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		userID := req.FormValue("user_id")
		if userID != "user_id_value" {
			api.Error(400, api.Response{Message: "invalid request"}, rw)
			return
		}

		name := req.FormValue("name")
		if name == "" {
			api.Error(400, api.Response{Message: "please provide name"}, rw)
			return
		}

		message := fmt.Sprintf("welcome %s", name)
		api.Success(http.StatusCreated, api.Response{Message: message}, rw)
	})
}

// reverse proxy ---> main_server

// mack and jack

// use case :
//  mack and his team mentains repo which authenticate users using different ways.
//  And jack decide to use there exsting system to authenticate users but also want
//  to hide this information from external world.
//  so he decided to use revers proxy.
