package router

import(
	"study/handler"

	"github.com/gin-gonic/gin"
)

func Router(){
	r := gin.Default()

	//library书城
	/*v2 := r.Group("/library")
	{

		//书城主界面
		v2.GET("/", handler.ViewLibrary)

		//书城排行榜
		v2.GET("/ranking", handler.ViewLibraryRanking)

		//分类下的书籍
		v2.GET("/:class_id", handler.ViewLibraryClass)
	}*/

	//digest书摘
	v3 := r.Group("/digest")
	{

		//书摘主界面
		v3.GET("/:user_id/mysummary", handler.ViewDigest)

		//选择是否公开 //参数summary_id
		v3.PUT("/:user_id/mysummary", handler.IfPublic)

		//搜索书摘  //参数search
		v3.POST("/:user_id/mysummary", handler.SearchDigest)

		//个人书摘内容
		/*v3.GET("/:user_id/:digest_id", handler.ViewPersonalDigest)

		//编辑书摘
		v3.PUT("/:user_id/:digest_id", handler.EditDigest)

		//创作书摘
		v3.POST("/", handler.AddDigest)

		//书摘分类
		v3.GET("/:user_id/mysummary/classes", handler.ViewDigestClasses)

		//分类下的书摘
		v3.GET("/:user_id/mysummary/:class_id", handler.ViewClassDigest)

		//编辑书摘分类界面
		v3.GET("/:user_id/mysummary/classes/edit", handler.EditDigestClasses)

		//添加书摘分类
		v3.POST("/:user_id/mysummary/classes/edit", handler.AddDigestClass)

		//删除书摘类别
		v3.DELETE("/:user_id/mysummary/classes/edit", handler.DeleteDigestClass)*/
	}

	r.Run()
}
