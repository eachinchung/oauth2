// @Title        wechat_users
// @Description  微信表储存用户 openid、unionid
// @Author       Eachin
// @Date         2021/4/3 12:33 上午

package mysql

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// WechatUsers 用户微信表
type WechatUsers struct {
	gorm.Model
	UserID     uint64         `gorm:"unique;column:user_id;type:bigint unsigned;not null"`    // 用户ID
	OpenID     string         `gorm:"unique;column:open_id;type:varchar(32)"`                 // 微信 openid
	UnionID    string         `gorm:"unique;column:union_id;type:varchar(32)"`                // 微信 unionid
	Status     uint8          `gorm:"column:status;type:tinyint unsigned;not null;default:0"` // 状态
	JSONExtent datatypes.JSON `gorm:"column:json_extent;type:json"`                           // json拓展
}
