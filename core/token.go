package core

import (
	"fmt"
	"guizizhan/pkg/md5"
	"time"
)

func GenerateToken() string {
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := md5.MD5Encode(str)
	return temp
}
