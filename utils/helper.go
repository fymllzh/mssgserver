package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/toolkits/net"
)

//获得UUID
func GenUUID() string {
	u, _ := uuid.NewRandom()
	return u.String()
}

//获取本机IP
func GetLocalIp() (myip string) {
	ips, err := net.IntranetIP()
	if err != nil {
		fmt.Printf("get local ip error:%s", err)
		panic(err)
	}

	var localip string
	for _, ip := range ips {
		localip = ip
		break
	}
	return localip
}

//获取ip
func GetRequestIP(c *gin.Context) string {
	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	return reqIP
}

//获取随机值（毫秒数）
func GenRandMillSec() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return int(r.Float32() * 10 * 1000)
}

func Map2Url(m map[string]interface{}) string {
	url := ""
	for k, v := range m {
		url += "&" + k + "=" + any(v)
	}

	return url[1:]
}

func Map2Url2(m map[string]string) string {
	url := ""
	for k, v := range m {
		url += "&" + k + "=" + v
	}

	return url[1:]
}
func URange(min int, max int) (s []string) {

	for i := min; i <= max; i++ {
		s = append(s, strconv.Itoa(i))
	}
	return
}

func any(value interface{}) string {
	switch value.(type) {
	case string:
		str1, _ := value.(string)
		return str1
	case int32:
		str2, _ := value.(int32)
		return strconv.Itoa(int(str2))
	case int64:
		str3, _ := value.(int64)
		return strconv.Itoa(int(str3))
	case int:
		str4, _ := value.(int)
		return strconv.Itoa(str4)
	case float64:
		str5, _ := value.(float64)
		return strconv.Itoa(int(str5))
	case float32:
		str6, _ := value.(float32)
		return strconv.Itoa(int(str6))
	default:
		return ""
	}
}

func Str2Int(value string) int {
	n, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return n
}
func Md5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

func Hmac(key, data string) string {
	hmac := hmac.New(md5.New, []byte(key))
	hmac.Write([]byte(data))
	return hex.EncodeToString(hmac.Sum([]byte("")))
}

func Sha1(data string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum([]byte("")))
}

func Repeated(target []string) (ret []string) {
	sort.Strings(target)

	targetLen := len(target)
	for i := 0; i < targetLen; i++ {
		if (i > 0 && target[i-1] == target[i]) || len(target[i]) == 0 {
			continue
		}
		ret = append(ret, target[i])
	}
	return
}

func ActionId2Name(id string) string {
	names := map[string]string{
		"1": "任务上线",
		"2": "任务下线",
		"3": "任务暂停",
		"4": "任务恢复",
		"5": "立即执行",
	}
	return names[id]
}
