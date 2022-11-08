package route

import (
	"msgoblog/api"
	"msgoblog/views"
	"net/http"
)

func Router() {
	//1. 页面  views 2. api 数据（json） 3. 静态资源
	//1.homepage
	http.HandleFunc("/", views.HTML.Index)
	//2.demofunc
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	//3.
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
