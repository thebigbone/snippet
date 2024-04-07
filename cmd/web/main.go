package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
)

// dependency injection for using custom log across the app
type applicaton struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "network address")
	custom_log := flag.String("log", "", "custom file path for log files")
	flag.Parse()

	// for clear log messages in the console
	infoLog := log.New(os.Stdout, color.GreenString("INFO\t"), log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, color.RedString("ERROR\t"), log.Ldate|log.Ltime|log.Lshortfile)

	// application instance
	app := &applicaton{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	if *custom_log != "" {
		f, err := os.OpenFile(*custom_log, os.O_RDWR|os.O_CREATE, 0666)

		if err != nil {
			errorLog.Fatal(err)
		}

		defer f.Close()
		infoLog.SetOutput(f)
		errorLog.SetOutput(f)
	}

	// custom server config
	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("starting the server at %s", *addr)
	err := server.ListenAndServe()
	errorLog.Fatal(err)
}
