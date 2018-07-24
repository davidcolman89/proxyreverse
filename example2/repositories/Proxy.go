package repositories

import (
	"net/http"
	"io/ioutil"
	"log"
)

type proxyRepo struct {

}

func NewProxyRepo() Proxy{
	return proxyRepo{}
}

func (r proxyRepo) Call(target string)  ([]byte, error){

	log.Println("Get to target: ", target)
	//TODO QUE PASA SI VIENE POR POST....  
	resp, err := http.Get(target)


	if err != nil {
		return nil, err
	}

	log.Println("Read Body")
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil

}