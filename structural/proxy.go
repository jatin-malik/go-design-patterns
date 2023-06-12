package structural

import "fmt"

// Proxy is a structural design pattern that provides an object that acts as a substitute for a real service object
// used by a client. A proxy receives client requests, does some work (access control, caching, etc.)
// and then passes the request to a service object.

// Nginx can act as a proxy server for your application server.

type Server interface {
	handleRequest(string, string) (int, string)
}

type AppServer struct {
}

func (app *AppServer) handleRequest(url, method string) (int, string) {
	// Only two paths we are handling here
	if method == "GET" && url == "/health/check" {
		return 200, "OK"
	}
	if method == "POST" && url == "/create/user" {
		return 201, "Created"
	}
	return 404, "Not Found"
}

type Nginx struct {
	app                *AppServer
	maxRequestsAllowed int
	rateLimiter        map[string]int
}

func (server *Nginx) handleRequest(url, method string) (int, string) {
	// Does rate limiting per url
	if !server.isAllowed(url) {
		return 429, "Rate Limited"
	}
	// Delegates work to real appServer
	return server.app.handleRequest(url, method)
}

func (server *Nginx) isAllowed(url string) bool {
	if _, ok := server.rateLimiter[url]; !ok {
		server.rateLimiter[url] = 1
	}
	if server.rateLimiter[url] > server.maxRequestsAllowed {
		return false
	}
	server.rateLimiter[url] += 1
	return true
}

func getNginx() *Nginx {
	return &Nginx{
		app:                &AppServer{},
		maxRequestsAllowed: 2,
		rateLimiter:        make(map[string]int),
	}
}

func RunProxyDemo() {
	// Client works with nginx proxy only
	n := getNginx()

	healthCheck := "/health/check"
	createUser := "/create/user"

	code, resp := n.handleRequest(healthCheck, "GET")
	fmt.Printf("Url %s -> Got status %d and response %s\n", healthCheck, code, resp)
	code, resp = n.handleRequest(healthCheck, "GET")
	fmt.Printf("Url %s -> Got status %d and response %s\n", healthCheck, code, resp)
	code, resp = n.handleRequest(healthCheck, "GET")
	fmt.Printf("Url %s -> Got status %d and response %s\n", healthCheck, code, resp)
	code, resp = n.handleRequest(createUser, "POST")
	fmt.Printf("Url %s -> Got status %d and response %s\n", createUser, code, resp)
	code, resp = n.handleRequest(createUser, "POST")
	fmt.Printf("Url %s -> Got status %d and response %s\n", createUser, code, resp)
}
