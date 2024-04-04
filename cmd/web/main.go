package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
)

func main() {

	addr := flag.String("addr", ":4000", "network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, color.GreenString("INFO\t"), log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, color.RedString("ERROR\t"), log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	mux.HandleFunc("/home", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	infoLog.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)
}
