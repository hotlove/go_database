package dao

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	config "go_database/config"
	"sync"
)

type DaoUtil struct {
	db *sql.DB
}

var once sync.Once
var daoUtil *DaoUtil

func NewDaoUtil() (daoUtil *DaoUtil) {

	// 保证只会执行一次 db保证只会open一次 db已经实现连接池
	once.Do(func() {
		dbConfig := config.New()

		var driver = dbConfig.Username + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + dbConfig.Port + ")/" + dbConfig.Dbname + "?charset=utf8"

		db, err := sql.Open("mysql", driver)

		checkErr(err)

		// 提前定义了变量 但是还没有分配空间 这里new进行分配空间

		// 以下两种方式都可以
		//daoUtil = &DaoUtil{
		//	db: db,
		//}

		daoUtil = new(DaoUtil)
		daoUtil.db = db
	})

	return daoUtil
}