package service

import (
	"errors"
	"goblog/dao"
	"goblog/models"
	"goblog/util"
)

func Login(userName, passwd string) (*models.LoginRes, error) {
	passwd = util.Md5Crypt(passwd, "TaShini")
	user := dao.GetUser(userName, passwd)
	if user == nil {
		return nil, errors.New("账号或者密码不正确！")
	}
	uid := user.Uid
	// 生成token，jwt技术
	token, err := util.Award(&uid)
	if err != nil {
		return nil, errors.New("token 生成失败！")
	}
	userInfo := models.UserInfo{
		Uid:      user.Uid,
		UserName: user.UserName,
		Avatar:   user.Avatar,
	}
	lr := &models.LoginRes{
		Token:    token,
		UserInfo: userInfo,
	}
	return lr, nil
}
