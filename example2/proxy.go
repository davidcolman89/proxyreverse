package main

import (
	"net/http"
	"fmt"
	"net/url"
	"net/http/httputil"
	"github.com/davidcolman89/proxyreverse/example2/utils"
	"github.com/davidcolman89/proxyreverse/example2/repositories"
	"github.com/davidcolman89/proxyreverse/example2/services"
)

func main() {
	//http.HandleFunc("/", automaticProxy)
	http.HandleFunc("/", manualProxy)
	http.HandleFunc("/reverseProxy", automaticProxy)

	fmt.Println("Server Listen on Localhost:9999")
	fmt.Println("/ 				--> manual proxy")
	fmt.Println("/reverseProxy 	--> automatic proxy")

	http.ListenAndServe(":9999", nil)
}

func manualProxy(w http.ResponseWriter, r *http.Request) {

	repo := repositories.NewProxyRepo()
	service := services.NewProxyService(repo)

	err:= service.ReverseProxy(w , r )

	if err != nil {
		fmt.Println("Error:  ",err)
	}

}


func automaticProxy(w http.ResponseWriter, r *http.Request) {


	ip := utils.GetIp()
	destinyPath := r.URL.Query().Get("p")
	utils.Statistics(ip, destinyPath)

	remote := "http://localhost:8888" + destinyPath

	fmt.Println(remote)

	url, _ := url.Parse(remote)

	fmt.Println(url)

	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(w, r)

}
