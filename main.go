package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sync/errgroup"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorldHandler)
	run(mux)
}

func run(handler http.Handler) {
	runProduction(handler)
}

// https://stackoverflow.com/questions/37321760/how-to-set-up-lets-encrypt-for-a-go-server-application
// https://stackoverflow.com/a/40494806/17655004
func runProduction(handler http.Handler) {
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("shortener.kaplia.dev", "shortener.kaplia.dev"),
		Cache:      autocert.DirCache("/var/lib/tls-certificates/data/"),
	}

	server := &http.Server{
		Addr:    ":https",
		Handler: handler,
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
			MinVersion:     tls.VersionTLS12, // improves cert reputation score at https://www.ssllabs.com/ssltest/
		},
	}

	var g errgroup.Group

	g.Go(func() error {
		return http.ListenAndServe(":http", certManager.HTTPHandler(nil))
	})

	g.Go(func() error {
		return server.ListenAndServeTLS("", "") // Key and cert are coming from Let's Encrypt
	})

	log.Fatal(g.Wait())
}

func runDevelopment(handler http.Handler) {
	log.Fatal(http.ListenAndServe(":http", handler))
}
