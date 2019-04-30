package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	config "go_database/config"
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

func (userDao *UserDao) InserProfile(params ...interface{})  {
	fmt.Println(params)
	stmt, err := userDao.db.Prepare("INSERT userinfo SET id=? ,name=?")
	checkErr(err)
	res, err := stmt.Exec(params...)
	checkErr(err)
	fmt.Println(res)

	userDao.Close()
}

func checkErr(err error)  {
	if err != nil {
		panic(err)
	}
}