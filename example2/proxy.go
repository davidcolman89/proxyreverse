package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"net/url"
	"net/http/httputil"
	"github.com/davidcolman89/proxyreverse/example2/utils"
)

const defaultRemote = "https://api.mercadolibre.com/"
const defaultLocalRemote = "http://localhost:8888/people"

func main() {
	//http.HandleFunc("/", automaticProxy)
	http.HandleFunc("/", manualProxy)
	http.HandleFunc("/reverseProxy", automaticProxy)

	fmt.Println("Server Listen on Localhost:9999")
	http.ListenAndServe(":9999", nil)
}

func manualProxy(w http.ResponseWriter, r *http.Request) {


	ip := utils.GetIp()
	destinyPath := r.URL.Path
	utils.Statistics(ip, destinyPath)

	remote := defaultRemote + destinyPath
	resp, err := http.Get(remote)

	if err != nil {
		fmt.Println("Error:  ",err)
	}

	body, err := ioutil.ReadAll(resp.Body)


	fmt.Fprintf(w, string(body), r.URL.Path)

}


func automaticProxy(w http.ResponseWriter, r *http.Request) {


	ip := utils.GetIp()
	destinyPath := r.URL.Query().Get("p")
	utils.Statistics(ip, destinyPath)

	remote := defaultRemote + destinyPath

	fmt.Println(remote)

	url, _ := url.Parse(remote)

	fmt.Println(url)

	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(w, r)

}
