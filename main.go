package main

import (
	_ "go_database/entity"
	"go_database/server/http"
)

func main() {

	//var dev string
	//flag.StringVar(&dev, "dev", "dev", "context profile")
	//fmt.Println(dev)

	//profileDao := dao.NewProfileDao()

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

	//profiles := profileDao.QueryProfileId(30)
	//for _, profile := range profiles {
	//	fmt.Println("profile is", profile)
	//}

	http.Start()

	//http.NewServer().Start("127.0.0.1", 8080)

}
