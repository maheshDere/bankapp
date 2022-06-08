package server

import (
	"fmt"
	"strconv"

	"bankapp/config"
	"bankapp/utils"

	"github.com/urfave/negroni"
)

func StartAPIServer() {
	port := config.AppPort()
	server := negroni.Classic()

	dependencies, err := initDependencies()
	if err != nil {
		panic(err)
	}

	router := initRouter(dependencies)
	server.UseHandler(router)
	jwt, err := utils.Create("shubhamvyas@gmail.com", 1, 1)
	if err != nil {
		panic(err)
	}
	// fmt.Println(jwt)

	payload, err := utils.Validate(jwt)
	fmt.Println(payload, err)
	addr := fmt.Sprintf(":%s", strconv.Itoa(port))
	server.Run(addr)
}
