package handler

import(
	"strconv"

	"study/model"

	"github.com/gin-gonic/gin"
)

func AddDigestClass(c *gin.Context){
	var summaryClass model.SummaryClass
	var summaryClassName model.SummaryClassName

	user_id := c.Param("user_id")
	summaryClass.User_id, _  = strconv.ParseInt(user_id, 10, 64)

	err := c.BindJSON(&summaryClassName)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "编辑错误",
		})
		return
	}

	summaryClass.Class_name = summaryClassName.Name

	model.DB.Create(&summaryClass)

	c.JSON(200, gin.H{
		"message": "添加成功",
		"data": summaryClass,
	})

}
