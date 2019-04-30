package main

import (
	"go_database/dao"
)

func main()  {
	userDao := dao.New()
	userDao.InserProfile(29, "test")
}