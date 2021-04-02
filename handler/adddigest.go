package handler

import (
	"strconv"
	"time"

	"2021-Library-backend/model"

	"github.com/gin-gonic/gin"
)

// @Summary 添加书摘
// @Description 将创建的新书摘加入数据库中
// @Tags digest
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param summaryInfo body model.SummaryInfo true "summaryInfo"
// @Success 200 {object} []model.Summary "创建成功"
// @Failure 400 "编辑错误"
// @Router /digest [post]
func AddDigest(c *gin.Context) {

	token := c.Request.Header.Get("token")
	_, err0 := model.VerifyToken(token)
	if err0 != nil {
		c.JSON(404, gin.H{"message": "认证失败"})
		return
	}

	var summary model.Summary
	var summaryInfo model.SummaryInfo
	var book model.Book

	user_id := c.Query("user_id")
	class_id := c.DefaultQuery("class_id", "0")

	summary.User_id, _ = strconv.ParseInt(user_id, 10, 64)
	summary.Class_id, _ = strconv.ParseInt(class_id, 10, 64)

	//println(summary.User_id)

	err := c.BindJSON(&summaryInfo)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "编辑错误",
		})
		return
	}

	summary.Title = summaryInfo.Title
	summary.Chapter = summaryInfo.Chapter
	summary.Summary_information = summaryInfo.Summary_information
	summary.Thought = summaryInfo.Thought
	summary.Public = summaryInfo.Public

	model.DB.Where("book_name = ?", summary.Title).First(&book)
	/*
		if book.Book_id == 0 {
			book.Book_id = 1
		}
	*/

	summary.Book_id = book.Book_id

	summary.Date = time.Now().Format("2006-01-02 15:04:05")

	model.DB.Create(&summary)

	c.JSON(200, gin.H{
		"message": "创建成功",
		"data":    summary,
	})

}
