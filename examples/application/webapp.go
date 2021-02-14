package main

import (
	"fmt"
	"html/template"
	"net"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseGlob("src/templates/*.html")) //Initialize all the html templates in the templates folder
}

//Starts a local web server that listens on a random
//port, then returns the address it's listening on
func runWebapp() (address string) {
	mux := mux.NewRouter()

	mux.HandleFunc("/", indexHandler).Methods("GET")

	mux.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./src/js"))))
	mux.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./src/css"))))
	mux.PathPrefix("/asm/").Handler(http.StripPrefix("/asm/", http.FileServer(http.Dir("./src/asm"))))

	listen, err := net.Listen("tcp", "127.0.0.1:0") //Listening on ":0" will bind to localhost on a random port
	if err != nil {
		panic(fmt.Sprintf("%+v", errors.WithStack(err)))
	}

	go func() { //Serving the web server
		err := http.Serve(listen, mux)
		if err != nil {
			panic(fmt.Sprintf("%+v", errors.WithStack(err)))
		}
	}()

	return ("http://127.0.0.1:" + strconv.Itoa(listen.Addr().(*net.TCPAddr).Port))

}
