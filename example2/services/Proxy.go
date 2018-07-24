package services

import (
	"github.com/davidcolman89/proxyreverse/example2/repositories"
	"github.com/davidcolman89/proxyreverse/example2/utils"
	"net/http"
	"fmt"
	"log"
)


const defaultApi = "https://api.mercadolibre.com/"
const localApi = "http://localhost:8888/people"


type proxyService struct {
	ProxyRepo repositories.Proxy
}


func NewProxyService(repo repositories.Proxy) Proxy{
	return proxyService{repo}
}

func (s proxyService) ReverseProxy(w http.ResponseWriter, r *http.Request) (error) {

	ip := utils.GetIp()
	destinyPath := r.URL.Path
	target := defaultApi + destinyPath

	utils.Statistics(ip, destinyPath)

	body, err := s.ProxyRepo.Call(target)

	log.Println("Return response from target")
	fmt.Fprintf(w, string(body), destinyPath)

	return err
}