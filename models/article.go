package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title   string `binding:"required"` //标题
	Content string `binding:"required"` //文章内容
	Preview string `binding:"required"` //预览
	Likes   int    `gorm:"default:0"`   //点赞数
}

//这里由于前端部分传来的JSON都是开头大写的，所以不需要加JSON标签
