package controllers

import "net/http"

func (pm *PrivateMessager) HTTPServe() http.Handler {
	return http.HandlerFunc(pm.PrivateMessager)
}

func (pm *PrivateMessager) SetMethods() []string {
	return []string{"GET"}
}

func (pm *PrivateMessager) EndPoint() string {
	return "/private"
}


func (pm *PrivateMessager) PrivateMessager(w http.ResponseWriter, r *http.Request) {
	
}