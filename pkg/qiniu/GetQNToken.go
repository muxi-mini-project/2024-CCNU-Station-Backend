package qiniu

import (
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

//func UploadFileToQiniu(localFilePath string) (string, error) {
//	mac := qbox.NewMac(utils.Key.AccessKey, utils.Key.SecretKey)
//	cfg := storage.Config{
//		Zone:          &storage.ZoneHuanan,
//		UseCdnDomains: false,
//		UseHTTPS:      false,
//	}
//
//	uploader := storage.NewFormUploader(&cfg)
//	putPolicy := storage.PutPolicy{
//		Scope: utils.Key.Bucket,
//	}
//	token := putPolicy.UploadToken(mac)
//	ret := storage.PutRet{}
//	remoteFileName := "captcha/" + time.Now().String() + path.Base(localFilePath)
//	err := uploader.PutFile(context.Background(), &ret, token, remoteFileName, localFilePath, nil)
//	if err != nil {
//		return "", err
//	}
//	return utils.Key.Domain + "/" + ret.Key, nil
//}

func GetQNToken() string {
	var maxInt uint64 = 1 << 32
	putPolicy := storage.PutPolicy{
		Scope:   Key.Bucket,
		Expires: maxInt,
	}
	mac := qbox.NewMac(Key.AccessKey, Key.SecretKey)
	QNToken := putPolicy.UploadToken(mac)
	return QNToken
}
func GenerateURL(key string) string {
	URL := fmt.Sprintf("%s/%s", Key.Domain, key)
	return URL
}
