package services

import "net/http"

type Service interface {
	ConfigureEndpoints()
	InitService()
}

type Controller interface {
	HTTPServe() http.Handler
	SetMethods() []string
	EndPoint() string
}
