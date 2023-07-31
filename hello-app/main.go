package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	pathBase string
	port     string
	version  string
)

func init() {
	pathBase = os.Getenv("HELLO_APP_PATH_BASE")
	if pathBase != "" && !strings.HasPrefix(pathBase, "/") {
		pathBase = "/" + pathBase
	}
	port = os.Getenv("HELLO_APP_PORT")
	if port == "" {
		port = "80"
	}
	version = os.Getenv("HELLO_APP_VERSION")
	if version == "" {
		version = "v2"
	}
}

func headers(w http.ResponseWriter, r *http.Request) {
	var headers string
	for k, v := range r.Header {
		headers += fmt.Sprintf("%v: %v\n", k, v)
	}
	fmt.Fprintf(w, "Headers \n"+headers)
}

func greet(w http.ResponseWriter, r *http.Request) {
	host, _ := os.Hostname()
	if host == "" {
		host = "-"
	}
	log.Println("GET [200] /")
	fmt.Fprintf(w, "Hello World! \nTime now is: %v\nServer: %s\n Version: %s\n", time.Now().Format(time.RFC3339), host, version)
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
		Version  string
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
	http.HandleFunc(pathBase+"/configuration/mysettings", readMySettings)
	http.HandleFunc(rootPath, greet)
	log.Println("hello-app server started.")
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
