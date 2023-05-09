package http

import (
	"Server/application/api"
	"Server/application/core/room"
	"Server/application/db"
	"Server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
)

func StartHTTPServer(port int) (*gin.Engine, error) {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	//r.LoadHTMLGlob("application/http/templates/*")
	r.GET("/", func(c *gin.Context) {
		//c.HTML(http.StatusOK, "index.html", gin.H{
		//	"title": "My Website",
		//})
	})
	r.GET("/register", func(c *gin.Context) {
		//c.HTML(http.StatusOK, "register.html", gin.H{
		//	"title": "My Website",
		//})
	})
	r.GET("/login", func(c *gin.Context) {
		//c.HTML(http.StatusOK, "login.html", gin.H{
		//	"title": "My Website",
		//})
	})
	r.GET("/password", func(c *gin.Context) {
		//c.HTML(http.StatusOK, "password.html", gin.H{
		//	"title": "My Website",
		//})
	})
	r.GET("/file/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		local := "application/file/" + filename

		// 打开本地文件
		file, err := os.Open(local)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer file.Close()

		// 将文件发送给客户端
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
		c.Header("Content-Type", "application/octet-stream")
		c.File(filename)
	})
	r.POST("/register", Register)
	r.POST("/login", Login)
	r.POST("/password", ChangePassword)
	err := r.Run(fmt.Sprint(":", port))
	if err != nil {
		return nil, err
	}
	return r, nil
}
func Register(c *gin.Context) {
	// 获取用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 将用户名和加密后的密码存储到MongoDB

	err := db.Mongomanager.Register(username, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Account alreaday exists"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Register Success"})
	}

}

// 登录接口
func Login(c *gin.Context) {
	// 获取用户名和密码
	username := c.PostForm("username")

	// 查询MongoDB中是否存在该用户信息
	info, err := db.Mongomanager.FindUserInfoByUsername(username)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	} else {
		Key := utils.GlobalObject.SecretKey
		secertKey, _ := api.GenerateToken("123", "windows", 123, c.ClientIP(), []byte(Key))
		c.JSON(http.StatusOK, gin.H{
			"uid":      info.PID,
			"username": username,
			"token":    secertKey,
		})
		go room.RoomMgrObj.AddWaitingPlayer(info.PID, username, secertKey)
	}
}
func ChangePassword(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	newpassword := c.PostForm("newpassword")
	hashedPasswordOld, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashedPasswordNew, err := bcrypt.GenerateFromPassword([]byte(newpassword), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	err = db.Mongomanager.UpdatePassword(username, string(hashedPasswordOld), string(hashedPasswordNew))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "faild"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "okokokook"})
	}
}
