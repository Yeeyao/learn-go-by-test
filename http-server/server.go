package main

import (
	"fmt"
	"net/http"
)

//type Handler interface {
//	ServerHTTP(ResponserWriter, *Resquest)
//}

//func ListenAndServe(addr string, handler Handler) error {
//	return handler.error()
//}

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	fmt.Fprint(w, GetPlayerScore(player))

}

func GetPlayerScore(name string) string {

	if name == "Pepper" {
		return "20"
	}

	if name == "Floyd" {
		return "10"
	}

	return ""
}
