package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"net/url"
	"log"
	"net/http/httputil"
)

func main() {

	http.HandleFunc("/proxy", manualProxy)
	http.HandleFunc("/automaticProxy", automaticProxy)

	http.ListenAndServe(":9999", nil)
}

func manualProxy(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("https://api.mercadolibre.com/categories/MLA97994")

	if err != nil {
		fmt.Println("Error:  ",err)
	}

	body, err := ioutil.ReadAll(resp.Body)


	fmt.Fprintf(w, string(body), r.URL.Path)

}


var cmd Cmd
var srv http.Server

func StartServer(bind string, remote string)  {
	log.Printf("Listening on %s, forwarding to %s", bind, remote)
	h := &handle{reverseProxy: remote}
	srv.Addr = bind
	srv.Handler = h
	//go func() {
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
	//}()
}

func automaticProxy(w http.ResponseWriter, r *http.Request) {


	//s := "https://api.mercadolibre.com/categories/MLA97994"
	s := "http://localhost:8888/"
	remote,err := url.Parse(s)

	if err != nil {
		log.Fatalln(err)
	}

	//http.Handle("http://localhost:9999", httputil.NewSingleHostReverseProxy(remote))

	StartServer()


}

