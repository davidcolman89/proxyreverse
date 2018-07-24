package services

import "net/http"

type Proxy interface {
	ReverseProxy(w http.ResponseWriter, r *http.Request) (error)
}
