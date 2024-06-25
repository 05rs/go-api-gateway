package main

import (
	"fmt"
	"go-api-gateway/config"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
)

const (
	// Todo will get from env.
	ENV                   = "dev"
	SERVICE_HOST          = ""
	SERVICE_PORT          = "8080"
	HTTP_RESPONSE_TIMEOUT = 5
)

func main() {
	router := mux.NewRouter()

	// Load configuration
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	cache := initializeCache(config)

	// Register services and their routes
	registerServices(router, config, cache)

	// Serve API gateway
	addr := SERVICE_HOST + ":" + SERVICE_PORT
	log.Printf("Starting [%s] server on %s", ENV, addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

func initializeCache(config *config.Config) *cache.Cache {
	// Create a new cache with the configured window and cleanup interval
	return cache.New(5*time.Minute, config.RateLimitWindow)
}

// Register services and their routes based on configuration
func registerServices(router *mux.Router, config *config.Config, cache *cache.Cache) {
	for serviceName, serviceDetails := range config.Services {
		log.Printf("Registering service: %s", serviceName)
		log.Printf(" Base URL: %s", serviceDetails.BaseURL)
		for _, route := range serviceDetails.Routes {
			// log.Printf("Path: %s", route.Path)
			log.Printf("  path: /%s%s", serviceName, route.Path)
			// Handle each route using a closure to capture configuration and route specifics
			router.HandleFunc(fmt.Sprintf("/%s%s", serviceName, route.Path), makeProxyHandler(serviceName, &route, config, cache))
		}
	}
}

// Creates a handler function to proxy requests to the appropriate backend service
func makeProxyHandler(serviceName string, route *config.Route, config *config.Config, cache *cache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !authenticate(r) {
			log.Printf("Unauthorized request from %s", r.RemoteAddr)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		log.Printf("Proxying request to service: %s", serviceName)
		proxyRequest(w, r, serviceName, route, config, cache)
	}
}

// Proxy the incoming request to the backend service
func proxyRequest(w http.ResponseWriter, r *http.Request, serviceName string, route *config.Route, config *config.Config, cache *cache.Cache) {
	serviceDetails, ok := config.Services[serviceName]
	if !ok {
		log.Printf("Service '%s' not found", serviceName)
		http.Error(w, fmt.Sprintf("Service '%s' not found", serviceName), http.StatusNotFound)
		return
	}

	// Check for rate limit
	if isRateLimited(serviceName, r.RemoteAddr, config.RateLimitWindow, config.RateLimitCount, cache) {
		log.Printf("Request from %s exceeded rate limit for service: %s", r.RemoteAddr, serviceName)
		http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
		return
	}

	target := serviceDetails.BaseURL
	targetURL, err := url.Parse(target)
	if err != nil {
		panic(err)
	}

	log.Printf("Proxying request to: %s", targetURL)
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.ServeHTTP(w, r)
}

// Dummy authentication function (replace with actual logic)
func authenticate(r *http.Request) bool {
	// Simulate a successful check for any token
	return true
}

// Check if a client IP has exceeded the rate limit for a service
func isRateLimited(serviceName string, ip string, rateLimitWindow time.Duration, rateLimitCount int, cache *cache.Cache) bool {
	cacheKey := fmt.Sprintf("rateLimit-%s-%s", serviceName, ip)
	countIface, ok := cache.Get(cacheKey)
	if !ok {
		// Initialize count for new client
		cache.Set(cacheKey, 1, rateLimitWindow)
		return false
	}
	count := countIface.(int)

	if count < rateLimitCount {
		cache.Set(cacheKey, count+1, rateLimitWindow)
		return false

	}
	return true

}
