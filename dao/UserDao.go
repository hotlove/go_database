package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	config "go_database/config"
	entity "go_database/entity"
)

type UserDao struct {
	db *sql.DB
}

func New()(dao *UserDao)  {
	dbConfig := config.New()

	fmt.Println(dbConfig)

	var driver = dbConfig.Username + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + dbConfig.Port + ")/" + dbConfig.Dbname + "?charset=utf8"

	dbConnect, err := sql.Open("mysql", driver)

	checkErr(err)

	return &UserDao{
		db: dbConnect,
	}

}

func (userDao *UserDao) Close()  {
	userDao.db.Close()
}

func (userDao *UserDao) InserProfile(profile *entity.Profile)  {
	stmt, err := userDao.db.Prepare("INSERT userinfo SET id=? ,name=?")
	checkErr(err)
	res, err := stmt.Exec(profile.Id, profile.Name)
	checkErr(err)
	fmt.Println(res)

	userDao.Close()
}

func checkErr(err error)  {
	if err != nil {
		panic(err)
	}
}