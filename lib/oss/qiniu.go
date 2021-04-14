// @Title        qiniu
// @Description	 七牛云 对象存储
// @Author       Eachin
// @Date         2021/4/7 9:14 下午

package oss

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"oauth/constant/exception"
	"oauth/logger"
	"oauth/setting"
)

// ByteUploader 字节数组上传
func ByteUploader(
	ctx *gin.Context,
	bucket, key, name string,
	data []byte,
) (string, error) {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(setting.Conf.QiniuConfig.AccessKey, setting.Conf.QiniuConfig.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": name,
		},
	}

	dataLen := int64(len(data))
	err := formUploader.Put(context.Background(), &ret, upToken, key+name, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		log := logger.NewSentry(ctx)
		log.Error("上传文件错误", err)
		return ret.Key, exception.ErrorServerBusy
	}
	return ret.Key, nil
}
