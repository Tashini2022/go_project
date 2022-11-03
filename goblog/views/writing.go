package views

import (
	"goblog/common"
	"goblog/service"
	"net/http"
)

func (*HTMLApi) Writing(w http.ResponseWriter, r *http.Request) {
	// // http://localhost:8080/writing
	writing := common.Template.Writing
	wr := service.Writing()
	writing.WriteData(w, wr)
}
