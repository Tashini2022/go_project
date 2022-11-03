package api

import (
	"errors"
	"goblog/common"
	"goblog/dao"
	"goblog/models"
	"goblog/service"
	"goblog/util"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (*Api) GetPost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path // /api/v1/post/xx
	pidStr := strings.TrimPrefix(path, "/api/v1/post/")
	pId, err := strconv.Atoi(pidStr)
	if err != nil {
		common.Error(w, errors.New("文章请求地址不存在！"))
		return
	}
	post, err := dao.GetPostByPid(pId)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, post)
}

func (*Api) SearchPost(w http.ResponseWriter, r *http.Request) {
	// http://localhost:8080/api/v1/post/search?val=xxxxx
	_ = r.ParseForm()
	condition := r.Form.Get("val")
	searchRes := service.SearchPost(condition)
	common.Success(w, searchRes)
}

func (*Api) SaveAndUpdatePost(w http.ResponseWriter, r *http.Request) {
	// 获取用户id 判断用户是否登录
	token := r.Header.Get("Authorization")
	_, claim, err := util.ParseToken(token)
	if err != nil {
		common.Error(w, errors.New("账户已过期"))
		return
	}
	uId := claim.Uid

	method := r.Method
	switch method {
	case http.MethodPost:
		// Post == Save
		param := common.GetRequestJsonParam(r)
		cId, _ := strconv.Atoi(param["categoryId"].(string))
		content := param["content"].(string)
		markdown := param["markdown"].(string)
		slug := param["slug"].(string)
		title := param["title"].(string)
		postType := param["type"].(float64)
		pType := int(postType)
		post := &models.Post{
			Pid:        -1,
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: cId,
			UserId:     uId,
			ViewCount:  0,
			Type:       pType,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.SavePost(post)
		common.Success(w, post)
	case http.MethodPut:
		// Put == Update
		param := common.GetRequestJsonParam(r)
		cId, _ := strconv.Atoi(param["categoryId"].(string))
		content := param["content"].(string)
		markdown := param["markdown"].(string)
		slug := param["slug"].(string)
		title := param["title"].(string)
		postType := param["type"].(float64)
		pidFloat := param["pid"].(float64)
		pid := int(pidFloat)
		pType := int(postType)
		post := &models.Post{
			Pid:        pid,
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: cId,
			UserId:     uId,
			ViewCount:  0,
			Type:       pType,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.UpdatePost(post)
		common.Success(w, post)
	}

}
