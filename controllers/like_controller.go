package controllers

import (
	"exchange_backend/global"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
)

// 点赞
func LikeArticle(ctx *gin.Context) {
	articleID := ctx.Param("id")

	//Redis中键的命名规则：单词之间用冒号间隔
	likeKey := "article:" + articleID + ":likes"

	//Incr让键自增1
	if err := global.RedisDB.Incr(likeKey).Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successful"})
}

// 获取点赞数
func GetArticleLikes(ctx *gin.Context) {
	articleID := ctx.Param("id")
	likeKey := "article:" + articleID + ":likes"

	likes, err := global.RedisDB.Get(likeKey).Result()
	if err == redis.Nil {
		likes = "0"
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"likes": likes})
}
