package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/saanai/util-sys/config"
)

const (
	DEBUG = "debug"
)

func init() {

	// time zone settings
	timeZone := "Asia/Tokyo"
	location, err := time.LoadLocation(timeZone)
	if err != nil {
		log.WithError(err).Fatal("load location failed.")
	}
	time.Local = location
}

func main() {
	// makeCertificate()

	// load configuration
	conf, err := config.NewConfig()
	if err != nil {
		log.WithError(err).Fatal("configuration is not valid")
	}

	// log level settings
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	if conf.App.LogLevel == DEBUG {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	mux := http.NewServeMux()

	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", errorHandler)

	mux.HandleFunc("/helloworld", helloHandler)
	mux.HandleFunc("/printHeader", printRequestHandler)

	//mux.HandleFunc("/login", login)
	//mux.HandleFunc("/logout", logout)
	//mux.HandleFunc("/signup_account", signupAccount)
	//mux.HandleFunc("/authenticate", authenticate)

	//mux.HandleFunc("/thread/new", newThread)
	//mux.HandleFunc("/thread/create", createThread)
	//mux.HandleFunc("/thread/post", postThread)
	//mux.HandleFunc("/thread/read", readThread)
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServeTLS("cert.pem", "key.pem")

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!, %s!", r.URL.Path[1:])
}

func printRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r)
}

// 自己証明書作成関数
func makeCertificate() {
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serilaNumber, _ := rand.Int(rand.Reader, max)
	subject := pkix.Name{
		Organization:       []string{"saanai inc."},
		OrganizationalUnit: []string{"Books"},
		CommonName:         "Go Web Programming",
	}

	template := x509.Certificate{
		SerialNumber: serilaNumber,
		Subject:      subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
	}
	pk, _ := rsa.GenerateKey(rand.Reader, 2048)

	certBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)
	certFile, _ := os.Create("cert.pem")
	pem.Encode(certFile, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes})
	certFile.Close()

	keyFile, _ := os.Create("key.pem")
	pem.Encode(keyFile, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	keyFile.Close()
}
