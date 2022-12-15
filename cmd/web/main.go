package main

import (
	"net/http"

	"github.com/navneetshukl/golang/pkg/handlers"
)


func main() {
	http.HandleFunc("/",handlers.Home)
	http.HandleFunc("/about",handlers.About)
	http.ListenAndServe(":3000",nil)

}