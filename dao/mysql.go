// @Title        mysql
// @Description  安装 mysql
// @Author       Eachin
// @Date         2021/4/2 10:16 下午

package dao

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"oauth/setting"
	"time"
)

var db *gorm.DB

func SetupMySQL(config *setting.MySQLConfig) (err error) {
	dns := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.DB,
	)
	if db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dns,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{}); err != nil {
		return err
	}

	var sqlDB *sql.DB
	if sqlDB, err = db.DB(); err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return
}
