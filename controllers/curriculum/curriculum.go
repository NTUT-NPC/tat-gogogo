package curriculum

import (
	"tat_gogogo/crawler/curriculum"

	"log"

	"github.com/gin-gonic/gin"
)

/*
Controller handles curriculum
it will get all the years and semesters
the default target student will be self
*/
func Controller(c *gin.Context) {
	studentID := c.PostForm("studentID")
	password := c.PostForm("password")
	targetStudentID := c.PostForm("targetStudentID")

	result, err := curriculum.GetCurriculums(studentID, password, targetStudentID)
	if err != nil {
		log.Panicln(err)
		c.Status(500)
		return
	}

	if result.Status != 200 {
		c.JSON(result.Status, gin.H{
			"message": result.Data,
		})
		return
	}

	c.JSON(result.Status, result.Data)
}
