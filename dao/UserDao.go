package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	config "go_database/config"
	entity "go_database/entity"
	"log"
)

type UserDao struct {
	db *sql.DB
}

/**
0. sql.db 自带的有连接池 只需要保证open一次 在真正的去操作数据库(curd)时才会真正的去获取链接
1. 对于数据库的 增加 删除 修改 操作 操作完后会自动释放链接所以不需要手动关闭
2. 对于查询操作不会释放 需要 rows.Close 去释放
 */
func New()(dao *UserDao)  {
	dbConfig := config.New()

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

/**
新增用户信息
 */
func (userDao *UserDao) InserProfile(profile *entity.Profile)  {
	stmt, err := userDao.db.Prepare("INSERT userinfo SET id=? ,name=?")
	checkErr(err)
	res, err := stmt.Exec(profile.Id, profile.Name)
	checkErr(err)
	fmt.Println(res)
	//userDao.Close()
}

/**
修改用户信息
 */
func (userDao *UserDao) UpdateProfileById(profile *entity.Profile)  {
	stmt, err := userDao.db.Prepare("update userinfo set name = ? where id = ?")
	checkErr(err)

	res, err := stmt.Exec(profile.Name, profile.Id)
	checkErr(err)
	fmt.Println(res)
	//userDao.Close()
}
/**
删除用户信息
 */
func (userDao *UserDao) DeleteProfileById(id int64)  {
	stmt, err := userDao.db.Prepare("delete from userinfo where id = ?")
	checkErr(err)

	res, err := stmt.Exec(id)
	checkErr(err)
	fmt.Println(res)
	//userDao.Close()
}

func (userDao *UserDao) QueryProfileId(id int64)  []entity.Profile{
	rows, err := userDao.db.Query("select * from userinfo where id = ?", id)
	checkErr(err)

	var profiles []entity.Profile
	for rows.Next() {
		var profile entity.Profile
		err := rows.Scan(&profile.Id, &profile.Name)
		if err != nil {
			log.Println(err)
			continue
		}
		profiles = append(profiles, profile)
	}
	rows.Close()
	return profiles
}

func checkErr(err error)  {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}