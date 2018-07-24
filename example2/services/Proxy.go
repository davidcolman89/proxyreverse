package services

import (
	"github.com/davidcolman89/proxyreverse/example2/repositories"
	"github.com/davidcolman89/proxyreverse/example2/utils"
	"net/http"
	"fmt"
	"io/ioutil"
)


const defaultApi = "https://api.mercadolibre.com/"
const defaultLocalRemote = "http://localhost:8888/people"


type proxyService struct {
	ProxyRepo repositories.Proxy
}


func NewProxyService(repo repositories.Proxy) Proxy{
	return proxyService{repo}
}

func (s proxyService) ReverseProxy(w http.ResponseWriter, r *http.Request) (error) {

	ip := utils.GetIp()
	destinyPath := r.URL.Path

	utils.Statistics(ip, destinyPath)

	target := defaultApi + destinyPath

	resp, err := http.Get(target)

	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	fmt.Fprintf(w, string(body), destinyPath)

	return nil
}