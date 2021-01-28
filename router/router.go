package router

import (
	"2021-Library-backend/handler"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {

	r.POST("/user", handler.User)
	r.POST("/login", handler.Login)

	//Group1 := router("/homepage") {
	//	Group1.GET("/", handler.Homepage)                             //用户主页
	//Group1.GET("/:user_id", handler.HomePage)                      //用户主页上用户的信息
	//	Group1.GET("/:user_id/mydigest", handler.MyDigest)            //显示我的发布
	//	Group1.GET("/:user_id/mydigest/digest_id", handler.digest_id) //从我的发布中查看书摘
	//	Group1.PUT("/:user_id/mydigest", handler.Ddigest)             //从我的发布删除书摘
	//	Group1.GET("/:user_id/shelf", handler.Shelf)                  //我的书架
	//	Group1.GET("/:user_id/shelf/books_id", handler.Books)         //从我的书架查看图书
	//	Group1.PUT("/:user_id/shelf/books_id", handler.Dbooks)        //从我的书架中删除书籍
	//	Group1.GET("/:user_id/info", handler.info)
	//	Group1.GET("/:user_id/info/userinfo", handler.userinfo)          //查看用户的基本信息
	//	Group1.PUT("/:user_id/info/userinfo", handler.ChangeInformation) //更改用户的昵称，头像，座右铭
	//	}

	//	Group2 := router("/Library") {
	//		Group2.POST("/searcher", handler.Searcher)         //搜索界面
	//		Group2.GET("/:books_id", handler.books_id)         //获取书籍的信息
	//		Group2.GET("/:books_id/digest_id", handler.digest) //显示该书的书摘信息
	//		Group2.POST("/:books_id", handler.adder)           //将这本书添加到我的书架里
	//	}

}
