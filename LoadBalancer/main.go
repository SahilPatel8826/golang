package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type Server interface {
	Address() string
	IsAlive() bool
	Serve(rw http.ResponseWriter, req *http.Request)
}

type simpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}

func newSimpleServer(addr string) *simpleServer {
	serverUrl, err := url.Parse(addr)
	handleErr(err)

	return &simpleServer{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}

}

type LoadBalancer struct {
	port              string
	roundRobinCounter int
	servers           []Server
}

func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:              port,
		roundRobinCounter: 0,
		servers:           servers,
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("error:%v\n", err)
		os.Exit(1)
	}
}

func (s *simpleServer) Address() string {
	return s.addr
}
func (s *simpleServer) IsAlive() bool {
	return true
}
func (s *simpleServer) Serve(rw http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(rw, req)
}

func (lb *LoadBalancer) getNextAvailableServer() Server {

	server := lb.servers[lb.roundRobinCounter%len(lb.servers)]
	for !server.IsAlive() {
		lb.roundRobinCounter++
		server = lb.servers[lb.roundRobinCounter%len(lb.servers)]
	}
	lb.roundRobinCounter++
	return server
}

func (lb *LoadBalancer) ServeProxy(rw http.ResponseWriter, req *http.Request) {
	targetServer := lb.getNextAvailableServer()
	fmt.Printf("forwarding request to address %q\n", targetServer.Address())
	targetServer.Serve(rw, req)
}

func main() {
	server := []Server{
		newSimpleServer("https://www.google.com"),
		newSimpleServer("https://www.facebook.com"),
		newSimpleServer("https://www.wikipedia.com"),
	}
	lb := NewLoadBalancer("8000", server)
	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.ServeProxy(rw, req)
	}
	http.HandleFunc("/", handleRedirect)

	fmt.Printf("Load Balancer started at :%s\n", lb.port)
	http.ListenAndServe(":"+lb.port, nil)
}
