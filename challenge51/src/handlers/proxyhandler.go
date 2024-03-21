package handlers

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func HandleRequestAndRedirect(w http.ResponseWriter, r *http.Request) {
	clientIP := r.RemoteAddr

	ipAddress := clientIP
	if strings.HasPrefix(clientIP, "[") && strings.Contains(clientIP, "]") {
		ipAddress = strings.TrimPrefix(clientIP, "[")
		ipAddress = strings.Split(ipAddress, "]")[0]
	}

	// Add X-Forwarded-For header with the client's IP address
	r.Header.Add("X-Forwarded-For", ipAddress)

	//check the host is in the forbidden-hosts.txt
	isForbidden, err := IsHostForbidden(r.URL.Host)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if isForbidden {
		http.Error(w, "Website not allowed "+r.URL.Host, http.StatusForbidden)
		return
	}

	targetURL, err := url.Parse(r.URL.Scheme + "://" + r.URL.Host)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// Serve the request using the reverse proxy
	w.Header().Set("X-Forwarded-Proto", r.URL.Scheme)
	w.Header().Set("X-Forwarded-Host", r.URL.Host)

	proxy.ServeHTTP(w, r)
}
