package handler

import(
	"strconv"
	"time"

	"study/model"

	"github.com/gin-gonic/gin"
)

func AddDigest(c *gin.Context){
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
		"data": summary,
	})

}
