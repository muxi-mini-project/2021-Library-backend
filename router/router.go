package router

import(
	"study/handler"

	"github.com/gin-gonic/gin"
)

func Router(){
	r := gin.Default()

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

		//删除书摘 //参数summary_id
		v3.DELETE("/mysummary/:user_id/delete", handler.DeleteDigest)

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

	r.Run()
}
