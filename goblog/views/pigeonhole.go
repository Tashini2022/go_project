package views

import (
	"goblog/common"
	"goblog/service"
	"net/http"
)

func (*HTMLApi) Pigeonhole(w http.ResponseWriter, r *http.Request) {
	// // http://localhost:8080/login
	pigeonhole := common.Template.Pigeonhole
	pigeonholeRes := service.FindPostPigeonholeRes()
	pigeonhole.WriteData(w, pigeonholeRes)
}
