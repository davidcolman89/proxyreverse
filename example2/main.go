package main

import (
	"net/http"
	"fmt"
	"github.com/davidcolman89/proxyreverse/example2/repositories"
	"github.com/davidcolman89/proxyreverse/example2/services"
	"github.com/gomodule/redigo/redis"
	"log"
)

func main() {

	http.HandleFunc("/", manualProxy)

	fmt.Println("Server Listen on Localhost:9999")
	fmt.Println("/ 				--> manual proxy")

	http.ListenAndServe(":9999", nil)
}

func manualProxy(w http.ResponseWriter, r *http.Request) {

	c, err := redis.Dial("tcp", ":6379")

	if err != nil {
		fmt.Println("Error:  ",err)
	}

	defer c.Close()

	repo := repositories.NewProxyRepo()
	service := services.NewProxyService(repo)

	err = service.ReverseProxy(w , r )

	if err != nil {
		fmt.Println("Error:  ",err)
	}

}