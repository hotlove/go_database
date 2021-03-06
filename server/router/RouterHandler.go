package router

import (
	"fmt"
	"net/http"
)

type RouterHandler struct{}

func NewRouterHandler() *RouterHandler {
	return &RouterHandler{}
}

var mux = make(map[string]func(w http.ResponseWriter, r *http.Request))

func (routerHandler *RouterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	if fun, ok := mux[r.URL.Path]; ok {
		fun(w, r)
		return
	}

	//静态资源
	//if strings.HasPrefix(r.URL.Path,constant.STATIC_BAES_PATH){
	//	if fun, ok := mux[constant.STATIC_BAES_PATH]; ok {
	//		fun(w, r)
	//		return
	//	}
	//}
	http.Error(w, "error URL:"+r.URL.String(), http.StatusBadRequest)
}

func (p *RouterHandler) Router(relativePath string, handler func(http.ResponseWriter, *http.Request)) {
	mux[relativePath] = handler
}
