package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const DefaultPort = 8080
const DefaultRedirect = "https://www.youtube.com/watch?v=dQw4w9WgXcQ"

// https://www.searchenginejournal.com/301-vs-302-redirects-seo/299843/#close
const DefaultStatusCode = http.StatusFound

var Version string

type AppConfig struct {
	// target url
	Target string
	//
	Port       int
	StatusCode int
}

// Parse flags and set defaults
func LoadConfig(name string, args []string) *AppConfig {
	flags := flag.NewFlagSet(name, flag.ContinueOnError)
	if Version != "" {
		Version = "dev"
	}

	conf := &AppConfig{}
	flags.StringVar(&conf.Target, "target", DefaultRedirect, "target URL for redirect either as param or 1st argument")
	flags.IntVar(&conf.StatusCode, "code", DefaultStatusCode, "HTTP status code used for redirects")
	flags.IntVar(&conf.Port, "port", DefaultPort, "listening server port")
	flags.Parse(args)

	arg1 := flags.Arg(0)
	if arg1 != "" && conf.Target != DefaultRedirect {
		fmt.Print("both flag and first argurment is is set, argument takes priority")
	}
	if arg1 != "" {
		conf.Target = arg1
	} else if arg1 == "help" {
		flags.Usage()
		os.Exit(2)
	} else if arg1 == "version" {
		fmt.Printf("version: %v", Version)
		os.Exit(2)
	}

	return conf
}

// Launch the redirect mux server from config
func Server(conf *AppConfig) error {
	log.Printf("Redirick will forward to: %s, listening on %d, will use status code %d.", conf.Target, conf.Port, conf.StatusCode)
	mux := http.NewServeMux()
	mux.Handle("/", http.RedirectHandler(conf.Target, conf.StatusCode))
	return http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), mux)
}

func main() {
	config := LoadConfig(os.Args[0], os.Args[1:])
	err := Server(config)
	if err != nil {
		log.Fatal(err)
	}
}
