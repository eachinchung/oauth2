// @Title        config
// @Description  配置表
// @Author       Eachin
// @Date         2021/4/4 8:35 下午

package mysql

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Configs 配置表
type Configs struct {
	gorm.Model
	ConfigKey string         `gorm:"uniqueIndex:key_version;column:config_key;type:varchar(32);not null"`   // 配置名称
	Version   uint8          `gorm:"uniqueIndex:key_version;column:version;type:tinyint unsigned;not null"` // 版本
	Config    datatypes.JSON `gorm:"column:config;type:json;not null"`                                      // 配置
	Status    uint8          `gorm:"column:status;type:tinyint unsigned;not null;default:0"`                // 状态
}
