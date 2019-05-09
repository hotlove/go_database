package entity

import "time"

type Profile struct {
	Id          int64     // 人员id
	Name        string    // 人员名称
	Password    string    // 密码
	Mobile      string    // 手机号码
	Email       string    // 邮件地址
	AreaCode    string    // 地区
	AreaInfo    string    // 详细地址
	CreatedTime time.Time // 创建时间
	Deleted     int64     // 是否删除
	Avatar      string    // 头像地址
}
