package router

import (
	"goblog/api"
	"goblog/views"
	"net/http"
)

// Router 路由
func Router() {
	// 	1 页面views	2 数据Api	3 静态资源
	http.HandleFunc("/", views.HTML.Index)
	// http://localhost:8080/c/1	1 参数(分类id)
	http.HandleFunc("/c/", views.HTML.Category)
	// http://localhost:8080/login
	http.HandleFunc("/login", views.HTML.Login)
	// http://localhost:8080/p/7.html
	http.HandleFunc("/p/", views.HTML.Detail)
	// http://localhost:8080/writing
	http.HandleFunc("/writing", views.HTML.Writing)
	// http://localhost:8080/pigeonhole
	http.HandleFunc("/pigeonhole", views.HTML.Pigeonhole)
	// http://localhost:8080/api/v1/post
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/post/", api.API.GetPost)
	// http://localhost:8080/api/v1/post/search?val=xxxxx
	http.HandleFunc("/api/v1/post/search", api.API.SearchPost)
	http.HandleFunc("/api/v1/qiniu/token/", api.API.QiniuToken)
	// http://localhost:8080/api/v1/login
	http.HandleFunc("/api/v1/login", api.API.Login)
	// 加载静态资源
	h1 := http.FileServer(http.Dir("public/resource/"))
	http.Handle("/resource/", http.StripPrefix("/resource/", h1))
}
