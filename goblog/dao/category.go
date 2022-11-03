package dao

import (
	"goblog/models"
	"log"
)

func CountGetAllPost() (count int) {
	row := DB.QueryRow("select count(1) from blog_post")
	_ = row.Scan(&count)
	return count
}

func CountGetAllPostBySlug(slug string) (count int) {
	row := DB.QueryRow("select count(1) from blog_post where slug = ?", slug)
	_ = row.Scan(&count)
	return count
}

func CountGetAllPostByCategoryId(cId int) (count int) {
	row := DB.QueryRow("select count(1) from blog_post where category_id = ?", cId)
	_ = row.Scan(&count)
	return count
}

func GetCategoryNameById(cId int) string {
	row := DB.QueryRow("select name from blog_category where cid = ?", cId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var categoryName string
	row.Scan(&categoryName)
	return categoryName
}

func GetAllCategory() ([]models.Category, error) {
	rows, err := DB.Query("select * from blog_category")
	if err != nil {
		log.Println("GetAllCategory 查询出错:", err)
		return nil, err
	}
	var categorys []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)
		if err != nil {
			log.Println("GetAllCategory 取值出错:", err)
			return nil, err
		}
		categorys = append(categorys, category)
	}
	return categorys, nil
}
