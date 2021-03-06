package dao

import (
	"fmt"
	entity "go_database/entity"
	"log"
)

type ProfileDao struct {
	daoUtil *DaoUtil
}

/**
0. sql.db 自带的有连接池 只需要保证open一次 在真正的去操作数据库(curd)时才会真正的去获取链接
1. 对于数据库的 增加 删除 修改 操作 操作完后会自动释放链接所以不需要手动关闭
2. 对于查询操作不会释放 需要 rows.Close 去释放
*/
func NewProfileDao() (dao *ProfileDao) {
	return &ProfileDao{}
}

/**
新增用户信息
*/
func (profileDao *ProfileDao) InserProfile(profile *entity.Profile) {
	stmt, err := daoUtil.db.Prepare("INSERT userinfo SET id=? ,name=?")
	checkErr(err)
	res, err := stmt.Exec(profile.Id, profile.Name)
	checkErr(err)
	fmt.Println(res)
	//userDao.Close()
}

/**
修改用户信息
*/
func (profileDao *ProfileDao) UpdateProfileById(profile *entity.Profile) {
	stmt, err := daoUtil.db.Prepare("update userinfo set name = ? where id = ?")
	checkErr(err)

	res, err := stmt.Exec(profile.Name, profile.Id)
	checkErr(err)
	fmt.Println(res)
	//userDao.Close()
}

/**
删除用户信息
*/
func (profileDao *ProfileDao) DeleteProfileById(id int64) {
	stmt, err := daoUtil.db.Prepare("delete from userinfo where id = ?")
	checkErr(err)

	res, err := stmt.Exec(id)
	checkErr(err)
	fmt.Println(res)
	//userDao.Close()
}

func (profileDao *ProfileDao) QueryProfileId(id int64) []entity.Profile {
	rows, err := daoUtil.db.Query("select * from userinfo where id = ?", id)
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

func checkErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
