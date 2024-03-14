package qiniu

import (
	"github.com/spf13/viper"
)

func QiniuInit() {

	Key = Qin{
		AccessKey: viper.GetString("qin.AccessKey"),
		SecretKey: viper.GetString("qin.SecretKey"),
		Bucket:    viper.GetString("qin.Bucket"),
		Domain:    viper.GetString("qin.Domain"),
	}

	//fmt.Println(Key)
}

var Key Qin
