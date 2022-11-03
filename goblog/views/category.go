package views

import (
	"errors"
	"goblog/common"
	"goblog/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Category(w http.ResponseWriter, r *http.Request) {
	categoryTemplate := common.Template.Category
	// // http://localhost:8080/c/1	1 参数(分类id)
	path := r.URL.Path
	cidStr := strings.TrimPrefix(path, "/c/")
	cId, err := strconv.Atoi(cidStr)
	if err != nil {
		categoryTemplate.WriteError(w, errors.New("请求路径不匹配"))
		return
	}

	if err := r.ParseForm(); err != nil {
		log.Println("表单获取出错！", err)
		categoryTemplate.WriteError(w, errors.New("系统错误，请联系管理员！"))
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	// 每页显示的数量
	pageSize := 10
	CategoryResponse, err := service.GetPostByCategoryId(cId, page, pageSize)
	if err != nil {
		categoryTemplate.WriteError(w, err)
		return
	}
	categoryTemplate.WriteData(w, CategoryResponse)
}
