package main

import "net/http"

func main() {
	p("ChitChat", version(), "started at", config.Address)
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	//
	// 所有路由模式在这里匹配定义在另一个文件中的处理器函数
	//
	mux.HandleFunc("/", index)
	// 错误处理
	mux.HandleFunc("/err", err)

}
