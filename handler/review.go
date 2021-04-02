package handler

import (
	"2021-Library-backend/model"
	"log"

	"github.com/gin-gonic/gin"
)

// @Summary "查看我的发布的评论"
// @Description "获取信息"
// @Tags library
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param digest_id path string true "digest_id"
// @Param books_id path string true "books_id"
// @Success 200 {object} model.ReviewInfo
// @Failure 404 "获取失败"
// @Router /Library/{books_id}/digest/{digest_id} [get]
func GetReview(c *gin.Context) {
	id := c.Param("digest_id")
	token := c.Request.Header.Get("token")
	_, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "身份验证失败"})
		log.Println(err)
		return
	}
	Review, err2 := model.GetAllReview(id)
	if err2 != nil {
		c.JSON(404, gin.H{"message": "查找失败"})
	}
	c.JSON(200, Review)
}

// @Summary "查看评论"
// @Description "获取信息"
// @Tags library
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param digest_id path string true "digest_id"
// @Param books_id path string true "books_id"
// @Success 200 {object} model.Userinfo
// @Failure 404 "获取失败"
// @Router /Library/{books_id}/digest/{digest_id}/review [get]
func Review(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "身份验证失败"})
		log.Println(err)
		return
	}
	UserInfo, err2 := model.GetSomeThing(id)
	if err2 != nil {
		c.JSON(404, gin.H{"message": "获取失败"})
		log.Println(err2)
		return
	}
	c.JSON(200, UserInfo)
}

// @Summary "创建评论"
// @Description "创建一个新的评论"
// @Tags library
// @Accept json
// @Produce json
// @Param reviewInfo body model.ReviewInfo true "reviewInfo"
// @Param token header string true "token"
// @Param digest_id path string true "digest_id"
// @Param books_id path string true "books_id"
// @Success 200
// @Failure 404 "身份验证失败"
// @Failure 401 "创建失败"
// @Router /Library/{books_id}/digest/{digest_id}/review [put]
func CreatReview(c *gin.Context) {
	var reviewInfo model.ReviewInfo
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "身份验证失败"})
		log.Println(err)
		return
	}
	err2 := c.BindJSON(&reviewInfo)
	if err2 != nil {
		c.JSON(401, gin.H{"message": "格式错误"})
	}
	DigestId := c.Param("digest_id")
	err1 := model.CreatNewReview(id, DigestId, reviewInfo.Content)
	if err1 != nil {
		c.JSON(401, gin.H{"message": "创建失败"})
		log.Println(err)
		return
	}
	c.JSON(200, gin.H{"message": "创建成功"})
	return
}
