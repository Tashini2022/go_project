package api

import (
	"goblog/common"
	"goblog/service"
	"net/http"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {
	// 接受用户名和密码，返回json数据
	param := common.GetRequestJsonParam(r)
	userName := param["username"].(string)
	passwd := param["passwd"].(string)
	loginRes, err := service.Login(userName, passwd)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, loginRes)
}
