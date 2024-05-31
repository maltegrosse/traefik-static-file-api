package main

import (
	"log"
	"net/http"
	"os"
)

var Config config

func main() {
	apiPort := getEnv("API_PORT", "8080")
	filename := getEnv("FILENAME", "services.yml")
	router := getEnv("ROUTER", "http")
	domain := getEnv("DOMAIN", ".example.com")
	service := getEnv("SERVICE", "svc@docker")
	certResolver := getEnv("CERT_RESOLVER", "le")
	entryPoint := getEnv("ENTRY_POINT", "websecure")
	Config = config{
		apiPort:      apiPort,
		filename:     filename,
		router:       router,
		domain:       domain,
		service:      service,
		certResolver: certResolver,
		entryPoint:   entryPoint,
	}
	http.HandleFunc("/", handleRequest)
	log.Println("Starting server on port", Config.apiPort)
	http.ListenAndServe(":"+Config.apiPort, nil)

}
func getEnv(name string, def string) string {
	if len(os.Getenv(name)) > 0 {
		return os.Getenv(name)
	} else {
		return def
	}
}
