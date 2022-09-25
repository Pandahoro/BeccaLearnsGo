package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Pandahoro/BeccaLearnsGo/go-rest-api/pkg/swagger/server/restapi"

	"github.com/Pandahoro/BeccaLearnsGo/go-rest-api/pkg/swagger/server/restapi/operations"
)

func main() {
	//innit swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}
	
	api := operations.NewHelloAPIAPI(swaggerSpec)
	server := restapi.NewServer(api)

	defer func() {
		if err := server.Shutdown(); err != nil {
			//error handling
			log.Fatalln(err)
		}
	}()

	server.Port = 8080

	api.CheckHealthHandler = operations.CheckHealthHandlerFunc(Health)

	api.GetHelloUserHandler = operations.GetHelloUserHandlerFunc(GetHellowUser)

	api.GetCatNameHandler = operations.GetCatNameHandlerFunc(GetCatByName)

	//start server listening
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

	
}

//Health route returns OK
func Health(operations.CheckHealthParams) middleware.Responder {
	return operations.NewCheckHealthOK().WithPayload("OK")

}

//GetHelloUser returns Hello + your name
func GetHelloUser(user operations.GetHelloUserParams) middleware.Responder {
	return operations.NewGetHelloUserOK().WithPayload("Hello " + user.User + "!")

}

//GetCatByName returns a cat in gif
func GetCatByName(cat operations.GetCatNameParams) middleware.Responder {

	var URL string
	if cat.Name != ""{
		URL = "https://github.com/Pandahoro/cats/raw/main" + cat.Name + ".gif"
	} else {
		// by defaut return sad cat
		URL = "https://github.com/Pandahoro/cats/raw/main/SadCatto.gif"
	}
	
	response, err := http.Get(URL)
	if err != nil {
		fmt.Println("error")
	}

	return operations.NewGetCatNameOK().WithPayload(response.Body)
}