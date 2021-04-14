// @Title        users
// @Description  用户表
// @Author       Eachin
// @Date         2021/4/2 10:30 下午

package mysql

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Users 用户表
type Users struct {
	gorm.Model
	UserID     uint64         `gorm:"unique;column:user_id;type:bigint unsigned;not null"`       // 用户ID
	Phone      string         `gorm:"unique;column:phone;type:varchar(32)"`                      // 手机号码
	Country    uint16         `gorm:"column:country;type:smallint unsigned;not null;default:86"` // 国家号（中国86）
	Username   string         `gorm:"index:username;column:username;type:varchar(32);not null"`  // 用户姓名
	Avatar     string         `gorm:"column:avatar;type:varchar(256);not null"`                  // 用户头像
	Status     uint8          `gorm:"column:status;type:tinyint unsigned;not null;default:0"`    // 状态
	JSONExtent datatypes.JSON `gorm:"column:json_extent;type:json"`                              // json拓展
}
