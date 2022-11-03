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

// Index 访问首页，需要将首页涉及的模板涉进行解析
func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	// 页面上涉及到的所有数据，必须按照前端要求定义
	// 数据库查询
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取出错！", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员！"))
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	// 每页显示的数量
	pageSize := 10
	path := r.URL.Path
	slug := strings.TrimPrefix(path, "/")
	hr, err := service.GetAllIndexInfo(slug, page, pageSize)
	if err != nil {
		log.Println("获取数据出错：", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员！"))
	}
	index.WriteData(w, hr)
}
