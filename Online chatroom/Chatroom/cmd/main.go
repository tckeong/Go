package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"database/sql"

	"github.com/gorilla/websocket"
	_ "github.com/mattn/go-sqlite3"
)

type Router struct {
	r        *gin.Engine
	htmlFile []string
	isLogin  bool
	UserName string
	userWs   map[string]*websocket.Conn
	upGrader websocket.Upgrader
}

var mutex sync.Mutex

func main() {
	router := Router{
		r: gin.Default(),
		htmlFile: []string{
			"../static/html/index.html",
			"../static/html/login.html",
			"../static/html/logout.html",
			"../static/html/login_success.html",
			"../static/html/login_fail.html",
		},
		isLogin:  false,
		UserName: "cookie_value",
		userWs:   make(map[string]*websocket.Conn, 0),
		upGrader: websocket.Upgrader{
			ReadBufferSize:  2048,
			WriteBufferSize: 2048,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
	router.setup()
	router.run()
}

func (router Router) setup() {
	router.r.LoadHTMLFiles(router.htmlFile...)
	//加载静态资源，例如网页的css、js
	router.r.Static("/static", "../static/js")

	router.r.GET("/", func(c *gin.Context) {
		router.isLogin = false
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.r.GET("/index", func(c *gin.Context) {
		// 删除cookie, 设置cookie MaxAge设置为-1，表示删除cookie
		if router.isLogin {
			SetCookie(c, router.UserName, 300)
		} else {
			SetCookie(c, router.UserName, -1)
		}
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.r.GET("/logout", func(c *gin.Context) {
		// 删除cookie, 设置cookie MaxAge设置为-1，表示删除cookie
		router.isLogin = false
		SetCookie(c, router.UserName, -1)
		c.HTML(http.StatusOK, "logout.html", nil)
	})

	GetMethod("/login-page", "login.html", http.StatusOK, router.r)
	router.r.POST("/login-success", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login_success.html", nil)
	})
	router.r.POST("/login-fail", func(c *gin.Context) {
		c.HTML(http.StatusUnauthorized, "login_fail.html", nil)
	})
	router.r.POST("/login", router.login)
	router.r.GET("/socket", router.socket)
}

func (router Router) run() {
	port := ":4579"
	fmt.Println("Server start!")
	router.r.Run(port)
}

func (router *Router) login(context *gin.Context) {
	username, password := context.PostForm("username"), context.PostForm("password")
	queryPassword := DbQuery(username)
	// 访问数据库
	if queryPassword == password {
		router.isLogin = true
		router.UserName = username
		//1.路由重定向
		//2.指定重定向的URL 通过HandleContext进行重定向到login-success或login-fail
		//3.实际开发中很少在路由中使用匿名函数
		context.Request.URL.Path = "/login-success"
		router.r.HandleContext(context)
		//url重定向（坑）
		//c.Redirect(http.StatusMovedPermanently, "/login")//301 永久移动
		//c.Redirect(http.StatusPermanentRedirect, "/logout")//308 永久重定向
		context.Redirect(http.StatusTemporaryRedirect, "/login-success") //307 临时重定向
	} else {
		context.Request.URL.Path = "/login-fail"
		router.r.HandleContext(context)
		context.Redirect(http.StatusTemporaryRedirect, "/login-fail") //307 临时重定向
	}
}

func GetMethod(pathName, fileName string, status_code int, r *gin.Engine) {
	r.GET(pathName, func(context *gin.Context) {
		context.HTML(status_code, fileName, nil)
	})
}

func DbQuery(username string) string {
	db, err := sql.Open("sqlite3", "../database/userData.db")
	if err != nil {
		fmt.Println(err.Error())
		return "-1"
	}
	defer db.Close()
	queryRequest := "SELECT password FROM DATA WHERE userName = '" + username + "'"
	rows, err := db.Query(queryRequest)
	if err != nil {
		fmt.Println(err.Error())
		return "-1"
	}
	_ = rows.Next()
	var password string
	err = rows.Scan(&password)
	if err != nil {
		return "-1"
	}
	return password
}

func SetCookie(c *gin.Context, username string, time int) {
	c.SetCookie("cookie_name", username, time, "/index", GetLocalIpAddr(), false, false)
	c.SetCookie("cookie_name", username, time, "/index", "localhost", false, false)
}

func (router Router) socket(c *gin.Context) {
	go Socket(router, c)
}

func GetLocalIpAddr() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip := strings.Split(localAddr.String(), ":")[0]
	return ip
}

func Socket(router Router, c *gin.Context) {
	userName := router.UserName
	//升级get请求为webSocket协议
	ws, err := router.upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	fmt.Println("连接成功")
	fmt.Printf("username = %v\n", userName)
	(router.userWs)[userName] = ws
	defer delete(router.userWs, userName)
	defer ws.Close()
	go socketReadMessage(router, ws, &err)
	for err == nil {
		continue
	}
}

func socketWriteMessage(router Router, err *error, message []string) {
	mutex.Lock()
	defer mutex.Unlock()
	userWs := router.userWs
	message_send := make([]byte, 0)
	Touser := message[0]
	message_send = append(message_send, []byte(message[2]+" from "+message[1])...)
	if v, ok := userWs[Touser]; ok {
		v.WriteMessage(1, message_send)
	} else {
		v = userWs[message[1]]
		v.WriteMessage(1, []byte(Touser+" does not exist !"))
		return
	}

}

func socketReadMessage(router Router, ws *websocket.Conn, err *error) {
	for {
		_, message, er := ws.ReadMessage()
		if er != nil {
			*err = er
			return
		}
		message_arr := strings.Split(string(message), "|")
		MessageLog(message_arr)
		if message_arr[0] == "all" {
			go WriteMessageAll(router, err, message_arr)
		} else {
			go socketWriteMessage(router, err, message_arr)
		}
	}
}

func WriteMessageAll(router Router, err *error, message []string) {
	mutex.Lock()
	defer mutex.Unlock()
	userWs := router.userWs
	message_send := make([]byte, 0)
	message_send = append(message_send, []byte(message[2]+" from "+message[1])...)
	for _, v := range userWs {
		*err = v.WriteMessage(1, message_send)
	}
}

func MessageLog(message []string) {
	filePath := "../log/log.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("\n\n\nLog file cannot open !\n\n\n")
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("[ ")
	for i := 0; i < len(message); i++ {
		writer.WriteString(message[i])
		if i != len(message)-1 {
			writer.WriteString(", ")
		}
	}
	writer.WriteString(" ]")
	writer.WriteString("  " + GetTimeNow() + "\n")
	writer.Flush()
}

func GetTimeNow() string {
	timeNow := time.Now().Unix()
	formatTimeStr := time.Unix(timeNow, 0).Format("2006/01/02 15:04:05")
	return formatTimeStr
}
