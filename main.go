package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// 创建路由处理函数
	handler := func(w http.ResponseWriter, r *http.Request) {
		// 在此处编写处理请求的逻辑
		fmt.Fprintln(w, "欢迎访问Dior官网")
	}

	// 将路由处理函数注册到路由器
	r.HandleFunc("/", handler)

	// 启动HTTP服务器并监听指定端口
	log.Fatal(http.ListenAndServe(":8080", r))
}
