package handler

import(
	"study/model"
	"github.com/gin-gonic/gin"
)

func ViewLibraryRanking(c *gin.Context){
	var books []model.Book
	var booksShow []model.Book
	num := 5 //排行上的书目数

	model.DB.Find(&books)

	for i := 0; i < len(books); i++ {
		for j := len(books)-1; j > i; j-- {
			if books[j].Click_sum > books[j-1].Click_sum {
				books[j], books[j-1] = books[j-1], books[j]
			}
		}

		if num > 0 {
			num--
		} else {
			break
		}

		booksShow = append(booksShow, books[i])
	}

	c.JSON(200, gin.H{
		"message": "获取成功",
		"data": booksShow,
	})

}
