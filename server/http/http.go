package http

import (
	"go_database/controller"
	"go_database/server/router"
	"log"
	"net/http"
	"strconv"
	"time"
)

type HttpServer struct {
}

func NewServer() *HttpServer {
	return &HttpServer{}
}

func Start() {
	NewServer().Start("127.0.0.1", 8080)
}

// 开启服务
func (httpServer *HttpServer) Start(host string, port int) {

	var hostPort = host + ":" + strconv.Itoa(port)

	routerHandler := router.NewRouterHandler()

	server := &http.Server{
		Addr:        hostPort,
		Handler:     routerHandler,
		ReadTimeout: 5 * time.Second,
	}

	// 注册路由
	RegisterRouter(routerHandler)
	log.Println("[INFO] register routers info")

	// 启动服务
	log.Println("[INFO] start server at:", port)
	server.ListenAndServe()

	//http.HandleFunc("/", sayhelloName)
	//http.ListenAndServe(hostPort, nil)
}

func RegisterRouter(routerHandler *router.RouterHandler) {
	new(controller.ProfileController).RegisterRouter(routerHandler)
}
