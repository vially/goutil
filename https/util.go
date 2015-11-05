package httpsutil

import (
	"log"
	"net/http"
	"os"
)

var DefaultSSLChecker = SSLEnabled

// RunHTTPSConditionally executes ListenAndServe when the SSL_ENABLED environment variable is false. Otherwise it
// executes ListenAndServeTLS
func RunHTTPSConditionally(addr, certFile, keyFile string, handler http.Handler) {
	l := log.New(os.Stdout, "[http] ", 0)
	l.Printf("listening on %s", addr)

	if DefaultSSLChecker() {
		l.Fatal(http.ListenAndServeTLS(addr, certFile, keyFile, handler))
	} else {
		l.Fatal(http.ListenAndServe(addr, handler))
	}
}

func SSLEnabled() bool {
	return os.Getenv("SSL_ENABLED") != "0" && os.Getenv("SSL_ENABLED") != "false"
}
