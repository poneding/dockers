package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	pathBase string
	port     int
	version  string
)

func init() {
	pathBase = os.Getenv("HELLO_APP_PATH_BASE")
	if pathBase != "" && !strings.HasPrefix(pathBase, "/") {
		pathBase = "/" + pathBase
	}
	if port, _ = strconv.Atoi(os.Getenv("HELLO_APP_PORT")); port == 0 {
		port = 80
	}

	version = os.Getenv("HELLO_APP_VERSION")
	if version == "" {
		version = "unknown"
	}
}

func headers(w http.ResponseWriter, r *http.Request) {
	var headers string
	for k, v := range r.Header {
		headers += fmt.Sprintf("%v: %v\n", k, v)
	}
	log.Println("GET [200] /headers")
	log.Println("Headers \n" + headers)
	fmt.Fprintf(w, "Headers \n"+headers)
}

func greet(w http.ResponseWriter, r *http.Request) {
	host, _ := os.Hostname()
	if host == "" {
		host = "-"
	}
	log.Println("GET [200] " + pathBase)
	fmt.Fprintf(w, "Hello World! \nTime now is: %v\nServer: %s\nVersion: %s\n", time.Now().Format(time.RFC3339), host, version)
}

// Simulate a long running request
func elapsed(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	time.Sleep(2 * time.Second)
	log.Println("GET [200] /elapsed")
	fmt.Fprintf(w, "Time elapsed: %v\n", time.Since(start))
}

func readMySettings(w http.ResponseWriter, r *http.Request) {
	var conf = struct {
		App      string
		Author   string
		UserName string
		Password string
	}{}
	// var conf configT
	app, err := os.ReadFile("./mysettings/app.json")
	if err == nil {
		json.Unmarshal(app, &conf)
	}
	secret, err := os.ReadFile("./mysettings/secret.json")
	if err == nil {
		json.Unmarshal(secret, &conf)
	}
	res, _ := json.MarshalIndent(conf, "", "  ")
	fmt.Fprintf(w, "%s", res)
}

func main() {
	rootPath := pathBase
	if pathBase == "" {
		rootPath = "/"
	}

	http.HandleFunc(pathBase+"/headers", headers)
	http.HandleFunc(pathBase+"/elapsed", elapsed)
	http.HandleFunc(pathBase+"/settings", readMySettings)
	http.HandleFunc(rootPath, greet)
	log.Printf("hello-app serve at 0.0.0.0:%d", port)
	go func() {
		for {
			log.Println("hello-app server is running.")
			time.Sleep(1 * time.Second)
		}
	}()

	log.Println()
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
