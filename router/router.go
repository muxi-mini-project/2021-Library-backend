package router

import (
	"2021-Library-backend/handler"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	//r := gin.Default()
	r.POST("/user", handler.User)
	r.POST("/login", handler.Login)
	//library书城
	v2 := r.Group("/library")
	{

		//书城主界面
		v2.GET("/", handler.ViewLibrary)

		//书城排行榜
		v2.GET("/ranking", handler.ViewLibraryRanking)

		//书城分类
		v2.GET("/classes", handler.ViewLibraryClasses)

		//分类下的书籍
		v2.GET("/class/:class_id", handler.ViewLibraryClass)
	}

	//digest书摘
	v3 := r.Group("/digest")
	{

		//书摘主界面
		v3.GET("/mysummary/:user_id", handler.ViewDigest)

		//选择是否公开 //参数summary_id
		v3.PUT("/mysummary/:user_id", handler.IfPublic)

		//搜索书摘
		v3.POST("/mysummary/:user_id", handler.SearchDigest)

		//删除书摘
		v3.DELETE("/mysummary/:user_id/delete/:digest_id", handler.DeleteDigestTwo)

		//个人书摘内容
		v3.GET("/person/:summary_id", handler.ViewPersonalDigest)

		//编辑书摘
		v3.PUT("/person/:summary_id", handler.EditDigest)

		//创作书摘 //参数user_id 和 class_id （默认为0）
		v3.POST("/", handler.AddDigest)

		//书摘分类
		v3.GET("/mysummary/:user_id/classes", handler.ViewDigestClasses)

		//分类下的书摘
		v3.GET("/mysummary/:user_id/classes/:class_id", handler.ViewClassDigest)

		//编辑书摘分类界面
		v3.GET("/mysummary/:user_id/classes_edit", handler.EditDigestClasses)

		//添加书摘分类
		v3.POST("/mysummary/:user_id/classes_add", handler.AddDigestClass)

		//删除书摘类别 //参数class_id
		v3.DELETE("/mysummary/:user_id/classes_edit", handler.DeleteDigestClass)
	}
	Group1 := r.Group("/homepage")
	{
		Group1.GET("/", handler.HomePage)                        //用户主页上用户的信息
		Group1.GET("/mydigest", handler.MyDigest)                //显示我的发布
		Group1.GET("/mydigest/:digest_id", handler.DigestId)     //从我的发布中查看书摘
		Group1.PUT("/mydigest/:digest_id", handler.DeleteDigest) //从我的发布删除书摘
		Group1.GET("/shelf", handler.Shelf)                      //我的书架
		Group1.GET("/shelf/:books_id", handler.BookId)           //从我的书架查看图书
		Group1.PUT("/shelf/:books_id", handler.Deletebooks)      //从我的书架中删除书籍
		Group1.GET("/info", handler.Userinfo)                    //查看用户的基本信息
		Group1.PUT("/info", handler.ChangeInformation)           //更改用户的昵称，头像，座右铭
	}

	//安卓提供	Group1.GET("/:user_id/info", handler.info)
	Group2 := r.Group("/Library")
	{
		Group2.POST("/searcher", handler.Searcher)                             //搜索界面
		Group2.GET("/:books_id", handler.BookId2)                              //获取书籍的信息
		Group2.GET("/:books_id/digest", handler.Digest)                        //显示该书的书摘信息
		Group2.GET("/:books_id/digest/:digest_id", handler.GetReview)          //显示该书摘下的评论
		Group2.POST("/addbook/:books_id", handler.AddBook)                     //将这本书添加到我的书架里
		Group2.GET("/:books_id/digest/:digest_id/review", handler.Review)      //进入写评论的页面
		Group2.PUT("/:books_id/digest/:digest_id/review", handler.CreatReview) //一个创建评论的功能
	}
	//删除这个路由	Group1.GET("/", handler.Homepage)

}
