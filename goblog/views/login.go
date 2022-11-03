package views

import (
	"goblog/common"
	"goblog/config"
	"net/http"
)

func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {
	// // http://localhost:8080/login
	login := common.Template.Login

	login.WriteData(w, config.Cfg.Viewer)
}
