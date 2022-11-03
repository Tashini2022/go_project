package views

import (
	"errors"
	"goblog/common"
	"goblog/service"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Detail(w http.ResponseWriter, r *http.Request) {
	//// http://localhost:8080/p/7.html
	detail := common.Template.Detail
	// 获取路劲参数
	path := r.URL.Path
	pidStr := strings.TrimPrefix(path, "/p/")
	pidStr = strings.TrimSuffix(pidStr, ".html")
	pId, err := strconv.Atoi(pidStr)
	if err != nil {
		detail.WriteError(w, errors.New("文章请求地址不存在！"))
		return
	}
	postRes, err := service.GetPostDetail(pId)
	if err != nil {
		detail.WriteError(w, errors.New("文章获取出错！"))
		return
	}
	detail.WriteData(w, postRes)
}
