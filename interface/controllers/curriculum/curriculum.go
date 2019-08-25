package curriculum

import (
	"tat_gogogo/infrastructure/api/handler"

	jwt "github.com/appleboy/gin-jwt/v2"

	"github.com/gin-gonic/gin"
)

/*
Controller is a function for gin to handle curriculum api
*/
func Controller(c *gin.Context) {
	targetStudentID := c.Query("targetStudentID")

	claims := jwt.ExtractClaims(c)
	studentID := claims["studentID"].(string)
	password := claims["password"].(string)

	loginHandler := handler.NewLoginHanlder(studentID, password)
	curriculumHandler := handler.NewCurriculumHandler(studentID, password, targetStudentID)

	result, err := loginHandler.Login()
	if err != nil {
		c.Status(500)
		return
	}

	if !result.GetSuccess() {
		c.JSON(401, gin.H{
			"message": result.GetData(),
		})
		return
	}

	isLoginCurriculumSuccess, err := loginHandler.LoginCurriculum()
	if err != nil {
		c.Status(500)
		return
	}

	if !isLoginCurriculumSuccess {
		c.JSON(401, gin.H{
			"message": "登入課程系統失敗",
		})
		return
	}

	curriculumResult, err := curriculumHandler.GetCurriculumResult()
	if err != nil {
		c.Status(500)
		return
	}

	c.JSON(curriculumResult.GetStatus(), curriculumResult.GetData())
}
