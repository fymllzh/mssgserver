package api

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"math/rand"
	"mssgserver/utils"
	"net/http"
	gintemplate "github.com/foolin/gin-template"

)

type loginForm struct {
	Email    string `json:"email" form:"email" binding:"required,email,max=60"`
	Password string `json:"password" form:"password" binding:"required,alphanum,min=6,max=12"`
}
type Logins struct {
	Id             int    `json:"id" db:"id"`
	Name           string `json:"name" db:"name"`
	Phone          string `json:"phone" db:"phone"`
	Status         int    `json:"status" db:"status"`
	Email          string `json:"email" db:"email"`
	Password       string `json:"password" db:"passwd"`
	LoginIp        string `json:"login_ip" db:"login_ip"`
	LoginTime      int64  `json:"login_time" db:"login_time"`
	LoginCount     int    `json:"login_count" db:"login_count"`
	LoginFailCount int    `json:"login_fail_count" db:"login_fail_count"`
	Salt           string `json:"salt" db:"salt"`
}


func updateLoginTime(id int, loginip string) {
	sql := "update ct_user set login_time = ?, login_ip = ?,login_fail_count = 0  where id = ?"
	_, err := utils.DB.Exec(sql, utils.GetUnix(), loginip, id)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func updateLoginError(id int, account string) {
	if id > 0 {
		sql := "update ct_user set login_fail_count = login_fail_count + 1, login_time = ? where id = ?"
		_, err := utils.DB.Exec(sql, utils.GetUnix(), id)
		if err != nil {
			fmt.Println(err.Error())
		}
	} else if account != "" {
		fmt.Println("")
	}

}
func LoginAuth(c *gin.Context) {
	utils.Logger.Info("login start")

	var form loginForm
	if err := c.ShouldBind(&form); err != nil {
		//c.Redirect(http.StatusFound, fmt.Sprintf("/admin/login?account=%s&msg=输入正确的账号和密码", form.Email))
		fmt.Println("绑定参数错误", err.Error())
		return
	}

	//登录
	var userinfo Logins
	sql := "select id,name,phone,status,email,passwd,login_ip,login_time,login_count,login_fail_count,salt from ct_user where email = ?"
	err := utils.DB.Get(&userinfo, sql, form.Email)
	if err != nil {
		fmt.Println("账号查询错误", err.Error())
		go updateLoginError(0, form.Email)
		//c.Redirect(http.StatusFound, fmt.Sprintf("/admin/login?account=%s&msg=账号或密码错误", form.Email))
		return
	}
	//校验失败次数
	loginFailStatus := utils.Viper.GetBool("base.login-fail")
	if loginFailStatus {
		loginFailWaitTime := utils.Viper.GetInt64("base.login-fail-wait-time")
		loginFailCount := utils.Viper.GetInt("base.login-fail-count")
		loginDate := utils.Unix2Date(userinfo.LoginTime)
		if loginDate == utils.GetDate() {
			if userinfo.LoginFailCount > loginFailCount {
				if utils.GetUnix()-userinfo.LoginTime < loginFailWaitTime {
					//c.Redirect(http.StatusFound, fmt.Sprintf("/admin/login?account=%s&msg=禁止访问%d分钟", form.Email, loginFailWaitTime/60))
					fmt.Printf("account=%s&msg=禁止访问%d分钟",form.Email,loginFailWaitTime/60)
					return
				}
			}
		}
	}
	//校验密码
	if userinfo.Password != utils.Md5(form.Password+userinfo.Salt) {
		fmt.Println("密码错误",userinfo.Password,form.Password,userinfo.Salt,utils.Md5(form.Password+userinfo.Salt))
		go updateLoginError(userinfo.Id, "")
		//c.Redirect(http.StatusFound, fmt.Sprintf("/admin/login?account=%s&msg=账号或密码错误", form.Email))
		return
	}
	//设置登录状态
	session := sessions.Default(c)
	session.Set("login_id", fmt.Sprintf("%d", userinfo.Id))
	session.Set("admin_name", userinfo.Email)
	err = session.Save()
	if err != nil {
		//c.Redirect(http.StatusFound, fmt.Sprintf("/admin/login?account=%s&msg=登录失败，请联系管理员", form.Email))
		fmt.Println(err)
		return
	}
	//更新登录信息
	go updateLoginTime(userinfo.Id, utils.GetRequestIP(c))

	c.Redirect(http.StatusFound, "/api/set/create")
}

func Login(c *gin.Context) {
	gintemplate.HTML(c, http.StatusOK, "logins", gin.H{
		"title":   "测试CT 任务管理平台",
		"bgurl":   fmt.Sprintf("/statics/images/rand/pic_%d.jpg", rand.Intn(6)),
	})
}

func (l *Logins) Items(id int) (logins []Logins, err error) {
	sql := "select id,name,phone,status,email,passwd,login_ip,login_time,login_count,login_fail_count,salt from ct_user where id = ? "
	err = utils.DB.Select(&logins, sql, id)
	return
}
