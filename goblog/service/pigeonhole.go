package service

import (
	"goblog/config"
	"goblog/dao"
	"goblog/models"
)

func FindPostPigeonholeRes() models.PigeonholeRes {
	// 查询所有的文章，进行月份整理
	posts, _ := dao.GetAllPostPage()
	pigaonholeMap := make(map[string][]models.Post)
	for _, post := range posts {
		month := post.CreateAt.Format("2006-01")
		pigaonholeMap[month] = append(pigaonholeMap[month], post)
	}

	// 查询所有的文章

	categorys, _ := dao.GetAllCategory()
	return models.PigeonholeRes{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Categorys:    categorys,
		Lines:        pigaonholeMap,
	}
}
