package main

import "net/http"
import "html/template"
import "path/filepath"
import "log"
import "net"

var major string
var minor string

func main() {
	sm := http.NewServeMux()
	sm.HandleFunc("/", serveIndexTemplate)
	log.Println("Starting Greenweb - v" + major + "-" + minor)
	l, err := net.Listen("tcp4", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.Serve(l, sm))
	log.Println("Listening...")
}

func serveIndexTemplate(w http.ResponseWriter, r *http.Request) {
	title := "Test Greenweb RPM - v" + major + "." + minor
	log.Println(r.Method, "/")
	// Testing
	// lp := filepath.Join("templates", "index.html")
	// Live
	lp := filepath.Join("/usr/share/greenweb/", "index.html")
	tmpl, _ := template.ParseFiles(lp)
	tmpl.Execute(w, title)
}
