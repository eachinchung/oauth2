// @Title        init
// @Description
// @Author       Eachin
// @Date         2021/4/7 10:27 下午

package sonyFlake

import (
	"github.com/gin-gonic/gin"
	"github.com/sony/sonyflake"
	"oauth/constant/exception"
	"oauth/logger"
	"time"
)

var sonyFlake *sonyflake.Sonyflake

// Init 初始化sonyFlake
func Init(startTime string, machineId uint16) (err error) {
	sonyMachineID := machineId
	var st time.Time
	if st, err = time.Parse("2006-01-02", startTime); err != nil {
		return err
	}
	settings := sonyflake.Settings{
		StartTime: st,
		MachineID: func() (uint16, error) {
			return sonyMachineID, nil
		},
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	return
}

// GenID 生成ID
func GenID(ctx *gin.Context) (uint64, error) {
	id, err := sonyFlake.NextID()
	if err != nil {
		log := logger.NewSentry(ctx)
		log.Error("生成 ID 失败", err)
		return 0, exception.ErrorServerBusy
	}
	return id, nil
}
