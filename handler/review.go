package handler

import (
	"2021-Library-backend/model"
	"log"

	"github.com/gin-gonic/gin"
)

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
