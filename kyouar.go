/* kyouar
*
* A tiny endpoint to create a QR code from a supplied URL. Useful for runtime usage, e.g:
*
*  <img src='https://kyouar.herokuapp.com/u/http://facebook.com'> <!-- QR code image to facebook.com -->
*
* Alternatively, calling /b/<url> will return a base64 encoded response, suitable for an img src tag.
*
* URL component can include http(s):// or just a domain.
*
 */

package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"

	qrcode "github.com/skip2/go-qrcode"

	"github.com/gorilla/mux"
)

func landing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, friend. Example usage: %s", (r.Host + "/u/www.google.com"))
}

func buildQR(url string) ([]byte, error) {
	var png []byte
	png, err := qrcode.Encode(url, qrcode.Medium, 256)
	return png, err
}

func handleURL(w http.ResponseWriter, r *http.Request) {

	mode := string(r.URL.String()[1]) // 'u' for an image, 'b' for base64

	defer func() { // Handle panic.
		if recover() != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("[!] Panic! Unable to build QR.")
		}
	}()

	iurl := r.URL.String()[3:] // Hack to allow ? in URL
	remoteIP := r.RemoteAddr
	remoteIP, _, err := net.SplitHostPort(remoteIP)
	if err != nil {
		remoteIP = r.RemoteAddr
	}

	// URL rejiggering. Mux doesn't like forward slashes, sometimes. http:// gets converted to http:/ but https:// remains intact, usually.
	if len(iurl) < 5 || iurl[0:4] != "http" {
		iurl = "http://" + iurl
	}
	if len(iurl) < 8 {
		log.Printf("[!] [%s] Error parsing URL: %s", remoteIP, iurl)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid URL!"))
		return
	}

	if iurl[0:6] == "http:/" && iurl[6:7] != "/" {
		iurl = "http://" + iurl[6:]
	} else if iurl[0:7] == "https:/" && iurl[7:8] != "/" {
		iurl = "https://" + iurl[7:]
	}
	_, err = url.ParseRequestURI(iurl)
	if err != nil {
		log.Printf("[!] [%s] Error parsing URL: %s", remoteIP, iurl)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid URL!"))
		return
	}

	// Good URL, let's create a QR code
	log.Printf("[+] [%s] Request to build URL: %s", remoteIP, iurl)
	imageBytes, err := buildQR(iurl)
	if err != nil {
		log.Printf("[!] [%s] Error parsing URL: %s", remoteIP, iurl)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to create QR code"))
	} else {
		if mode == "u" { //image
			w.Write(imageBytes)
		} else if mode == "b" {
			b64data := base64.StdEncoding.EncodeToString(imageBytes)
			//b64html := "data:image/png;base64, " + b64data
			w.Write([]byte("data:image/png;base64, " + b64data))
		} else {
			//Should be unreachable
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid URL!"))
		}
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Invalid URL!"))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("[W] $PORT not set, using default of 8000")
		port = "8000"
	}
	log.Printf("[+] Started kyouar. Listening on port %s\n", port)

	r := mux.NewRouter()
	r.SkipClean(true) //Mux doesn't like double forward slashes
	r.HandleFunc("/", landing).Methods("GET")
	r.HandleFunc("/u/{url:[a-z,A-Z,0-9,\\-,_,\\/,\\.,:,\\/\\/]*}", handleURL).Methods("GET") //Mux doesn't like forward slashes
	r.HandleFunc("/b/{url:[a-z,A-Z,0-9,\\-,_,\\/,\\.,:,\\/\\/]*}", handleURL).Methods("GET") //Same endpoint, but to return QR in a base64 HTML img form
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	http.ListenAndServe(":"+port, r)
}
