package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

var Rdb *redis.Pool

func init() {
	Rdb = &redis.Pool{
		MaxIdle:     Viper.GetInt("redis.MaxIdleConn"),
		MaxActive:   Viper.GetInt("redis.MaxActiveConn"),
		IdleTimeout: Viper.GetDuration("redis.MaxConnIdleTimeout") * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", Viper.GetString("redis.Host"))
			if err != nil {
				log.Fatalf("redis init error: %v\n", err)
				return nil, err
			}

			if Viper.GetString("redis.Password") != "" {
				if _, err := c.Do("AUTH", Viper.GetString("redis.Password")); err != nil {
					c.Close()
					log.Fatalf("redis auth error: %v\n", err)
					return nil, err
				}
			}

			if _, err := c.Do("SELECT", Viper.GetString("redis.DbNum")); err != nil {
				c.Close()
				log.Fatalf("redis select db error: %v\n", err)
				return nil, err
			}

			if _, err := c.Do("PING"); err != nil {
				log.Fatalf("redis ping error: %v\n", err)
				return nil, err
			}

			return c, nil
		},
	}
}

func SetJson(k string, v interface{}, expiration int) (reply interface{}, err error) {
	conn := Rdb.Get()
	defer conn.Close()

	str, _ := json.Marshal(v)
	return conn.Do("SET", k, str, "EX", expiration)
}

func GetJson(k string, v interface{}) bool {
	conn := Rdb.Get()
	defer conn.Close()

	cache, err := conn.Do("GET", k)
	if err != nil {
		log.Printf("redis error: %v\n", err)
		return false
	}

	if cache != nil && cache != "" {
		json.Unmarshal(cache.([]byte), v)
		return true
	}

	return false
}

type User struct {
	ID   int64  `redis:"id"`
	Name string `redis:"name"`
}

// 添加struct类型的值（需要注意的是结构体字段是可导出的字段名称，并且使用了字段标签 redis。）
func StructAdd() {
	u1 := User{
		ID:   1,
		Name: "name1",
	}
	conn := Rdb.Get()
	defer conn.Close()
	replyStruct, err := conn.Do("HMSET", redis.Args{}.Add("hkey1").AddFlat(&u1)...)
	if err != nil {
		fmt.Println("struct err: ", err)
	}
	// 最后存在redis里的样子如下(用struct的方式，redisgo会解析成key-value形式)：
	// id 1
	// name name1
	fmt.Println(replyStruct) // OK
}

// structValues 把hash读取到struct
func StructValues() {
	conn := Rdb.Get()
	defer conn.Close()
	v, err := redis.Values(conn.Do("HGETALL", "hkey1"))
	if err != nil {
		fmt.Println("redis.Values() err: ", err)
	}

	// redis.ScanStruct()
	u2 := new(User)
	if err := redis.ScanStruct(v, u2); err != nil {
		fmt.Println("redis.ScanStruct() err: ", err)
	}
	fmt.Printf("%+v\n", u2) // &{ID:1 Name:name1}
}


