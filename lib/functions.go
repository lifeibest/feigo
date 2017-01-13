package lib

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"math/rand"

	"strconv"
	"time"
)

//create md5 string
func Strtomd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}

//password hash function
func Pwdhash(str string) string {
	return Strtomd5(str)
}

func StringsToJson(str string) string {
	rs := []rune(str)
	jsons := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			jsons += string(r)
		} else {
			jsons += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}

	return jsons
}

//随机IP
func Randip() string {
	var ip string
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 1; i <= 4; i++ {
		if i == 4 {
			ip += strconv.Itoa(r.Intn(100))
		} else {
			ip += strconv.Itoa(r.Intn(100)) + "."
		}
	}
	return ip
}

//过滤html,截取字符串
func SubHTML2str(s string, start, length int) string {
	s = beego.HTML2str(s)
	s = beego.Substr(s, start, length)
	return s
}
