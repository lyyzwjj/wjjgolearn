package service

import (
	"fmt"
	"github.com/lyyzwjj/wjjgolearn/02liwenzhou/113blog/dao/db"
	"github.com/lyyzwjj/wjjgolearn/02liwenzhou/113blog/model"
)

func GetAllCategoryList() (categoryList []*model.Category, err error) {
	categoryList, err = db.GetAllCategoryList()
	if err != nil {
		fmt.Printf("get category list failed, err:%v\n", err)
		return
	}
	return
}
