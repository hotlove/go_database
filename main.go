package main

import (
	"go_database/dao"
	"go_database/entity"
)

func main() {
	userDao := dao.New()

	profile := new(entity.Profile)
	profile.Id = 30
	profile.Name = "test30"

	userDao.InserProfile(profile)
}
