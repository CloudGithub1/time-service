package main

import (
	"encoding/json"
	"net"
	"net/http"
	"strings"
	"time"
	"log"
)

type Resp struct {
	Timestamp string `json:"timestamp"`
	IP        string `json:"ip"`
}

func clientIP(r *http.Request) string {
	// Prefer X-Forwarded-For (comma separated), then X-Real-IP, then RemoteAddr
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		parts := strings.Split(xff, ",")
		return strings.TrimSpace(parts[0])
	}
	if xr := r.Header.Get("X-Real-Ip"); xr != "" {
		return xr
	}
	// fallback to RemoteAddr (may contain port)
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

func handler(w http.ResponseWriter, r *http.Request) {
	resp := Resp{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		IP:        clientIP(r),
	}
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(resp)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("time-service starting on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
