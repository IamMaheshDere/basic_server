package main

import (
	"demo_server/user"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	port := 9090
	server := negroni.Classic()

	router := mux.NewRouter()
	userRouter := router.PathPrefix("/user").Subrouter()
	server.UseHandler(router)

	userRouter.HandleFunc("/greet", user.Greet()).Methods(http.MethodGet)
	userRouter.HandleFunc("/welcome", user.Greet()).Methods(http.MethodGet)

	addr := fmt.Sprintf(":%s", strconv.Itoa(port))
	server.Run(addr)
}
