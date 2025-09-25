package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db5 *gorm.DB

func init() {
	// 连接数据库
	var err error
	db5, err = gorm.Open(sqlite.Open("test5.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database:" + err.Error())
	}
	db5.Exec("DROP TABLE IF EXISTS Users")
	db5.Exec("DROP TABLE IF EXISTS Posts")
	db5.Exec("DROP TABLE IF EXISTS Comments")
	err = db5.AutoMigrate(&User{}, &Post{}, &Comment{}) //自动迁移模型
	if err != nil {
		panic("failed to auto migrate database:" + err.Error())
	}

	// 初始化测试数据
	db5.Create(&User{Name: "张三", Email: "zhangsan@example.com"})
	db5.Create(&User{Name: "李四", Email: "lisi@example.com"})
	db5.Create(&User{Name: "王五", Email: "wangwu@example.com"})
	db5.Create(&Post{Title: "第一篇文章", Content: "这是第一篇文章的内容", UserID: 1})
	db5.Create(&Post{Title: "第二篇文章", Content: "这是第二篇文章的内容", UserID: 1})
	db5.Create(&Post{Title: "第三篇文章", Content: "这是第三篇文章的内容", UserID: 2})
	db5.Create(&Comment{Content: "第1条评论", PostID: 1})
	db5.Create(&Comment{Content: "第2条评论", PostID: 1})
	db5.Create(&Comment{Content: "第3条评论", PostID: 2})
	db5.Create(&Comment{Content: "第4条评论", PostID: 2})
	db5.Create(&Comment{Content: "第5条评论", PostID: 3})
	db5.Create(&Comment{Content: "第6条评论", PostID: 3})
	db5.Create(&Comment{Content: "第7条评论", PostID: 1})
}

type User struct {
	ID    uint   `gorm:"primarykey"`
	Name  string `gorm:"size:20;not null;"`
	Email string `gorm:"size:50;uniqueIndex;"`

	// 关联文章
	PostCount int    `gorm:"default:0"`
	Posts     []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
	ID            uint   `gorm:"primarykey"`
	Title         string `gorm:"size:100;not null;"`
	Content       string `gorm:"type:text;not null"`
	CommentStatus string `gorm:"default:'有评论'"`

	UserID   uint
	Comments []Comment `gorm:"foreignKey:PostID"`
}

type Comment struct {
	ID      uint   `gorm:"primarykey"`
	Content string `gorm:"type:text;not null"`
	PostID  uint   `gorm:"not null"`
}

// 查询某个用户发布的所有文章及其对应的评论信息
func queryUserPostsAndComments(userID uint) ([]Post, error) {
	var posts []Post
	err := db5.Preload("Comments").Where("user_id = ?", userID).Find(&posts).Error
	return posts, err
}

// 查询评论数量最多的文章信息
func queryMostCommentPost() (Post, int, error) {
	var post Post
	type Result struct {
		PostID uint
		Count  int
	}
	var result Result
	err := db5.Model(&Comment{}).
		Select("post_id, COUNT(*) AS count").
		Group("post_id").
		Order("count DESC").
		First(&result).Error
	if err != nil {
		return post, 0, err
	}
	err = db5.First(&post, result.PostID).Error
	return post, result.Count, err
}

// AfterCreate 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段
func (p *Post) AfterCreate(tx *gorm.DB) error {
	return tx.Model(&User{}).Where("id = ?", p.UserID).
		Update("post_count", gorm.Expr("post_count + ?", 1)).Error
}

// AfterDelete 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	var count int64
	err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		return tx.Model(&Post{}).Where("id = ?", c.PostID).
			Update("comment_status", "无评论").Error
	}
	return nil
}

/*
1:假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章），
Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表

2:基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。

3:继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/
func main() {
	posts, _ := queryUserPostsAndComments(2)
	for _, post := range posts {
		fmt.Printf("%+v\n", post)
	}
	post, count, _ := queryMostCommentPost()
	fmt.Printf("Post: %s, Comment Count: %d\n", post.Title, count)
}
