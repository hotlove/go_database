package main

import (
	"fmt"
	"go_database/dao"
	_ "go_database/entity"
)

func main() {
	profileDao := dao.NewProfileDao()

	// 新增操作
	//profile := new(entity.Profile)
	//profile.Id = 30
	//profile.Name = "test30"
	//userDao.InserProfile(profile)
	//
	//// 修改操作
	//profile2 := new(entity.Profile)
	//profile2.Id = 30
	//profile2.Name = "test30update"
	//userDao.UpdateProfileById(profile2)

	profiles := profileDao.QueryProfileId(30)
	for _, profile := range profiles {
		fmt.Println("profile is", profile)
	}
}
