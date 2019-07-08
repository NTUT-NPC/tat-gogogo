package login

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/cookieJar"
	"net/url"
	"tat_gogogo/consts"
	"log"
	"bytes"
	"encoding/json"
)

type Result struct {
	success bool
	status int
	message string
}

type LoginController struct {
	studentID string
	password string
}

func HandleLogin(c *gin.Context) {
	studentID := c.PostForm("studentId")
	password := c.PostForm("password")
	
	controller := LoginController{studentID: studentID, password: password}

	result := controller.handleRequest()

	c.JSON(result.status, gin.H{
		"success": result.success,
		"message": result.message,
	})
}

func (controller *LoginController) newClient() (*http.Client, *http.Request) {
	cookieJar, _ := cookiejar.New(nil)
	
	client := &http.Client{
		Jar: cookieJar,
	}

	data := 	url.Values{
		"forceMobile": {"mobile"},
		"mpassword": {controller.password}, 
		"muid": {controller.studentID},
	}

	req, err := http.NewRequest("POST", consts.Login, bytes.NewBufferString(data.Encode()))

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", consts.IndexPage)
	req.Header.Set("User-Agent", "Direk Android App")

	return client, req
}

func (controller *LoginController) handleRequest() (Result) {
	client, req := controller.newClient()
	resp, err := client.Do(req)
	
	if err != nil {
		log.Fatalln(err)
		return Result{success: false, status: 401}
	}

	defer resp.Body.Close()

	var data map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&data)

	statusCode := 200
	isSuccess := data["success"].(bool)
	message := "登入成功"
	if !isSuccess {
		statusCode = 401
		message = "帳號或密碼錯誤，請重新輸入。"
	}

	return Result{success: isSuccess, status: statusCode, message: message}
}
