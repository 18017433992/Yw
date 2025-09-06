package controllers

import (
	"fmt"
	"net/http"

	"example.com/blog/global"
	"example.com/blog/middleware"
	"example.com/blog/models"
	"github.com/gin-gonic/gin"
)

// 创建评论
func CreateComment(ctx *gin.Context) {
	var comment models.Comment
	var post models.Post

	user_id, ok := middleware.Somehandler(ctx.Copy())
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "用户不存在"})
		return
	}
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "入参不正确"})
		return
	}
	comment.UserID = user_id
	if err := validateComment(&comment); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	if err := global.Db.AutoMigrate(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := global.Db.Where("id = ?", comment.PostID).First(&post).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "文章不存在"})
		return
	}
	if err := global.Db.Create(&comment).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := models.CommentResponse{
		Content: comment.Content,
		UserID:  comment.UserID,
		PostID:  comment.PostID,
	}

	ctx.JSON(http.StatusOK, response)
}

// 根据文章ID获取该文章的所有评论
func GetCommentsByPostID(ctx *gin.Context) {
	var post models.Post
	var comments []models.Comment

	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if post.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请选择文章！"})
		return
	}
	if err := global.Db.Where("id = ?", post.ID).First(&post).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "文章不存在"})
		return
	}

	if err := global.Db.Where("post_id= ?", post.ID).Find(&comments).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "该文章无评论"})
		return
	}
	if len(comments) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "该文章无评论"})
		return
	}

	var commentResponses []models.CommentResponse
	for _, comment := range comments {
		reponse := models.CommentResponse{
			ID:      comment.ID,
			Content: comment.Content,
			UserID:  comment.UserID,
			PostID:  comment.PostID}
		commentResponses = append(commentResponses, reponse)
	}
	ctx.JSON(http.StatusOK, commentResponses)
}

// 自定义验证函数 文章标题和内容必填
func validateComment(comment *models.Comment) error {
	if comment.Content == "" {
		return fmt.Errorf("Content is required")
	}
	if comment.PostID == 0 {
		return fmt.Errorf("PostID is required")
	}
	if comment.UserID == 0 {
		return fmt.Errorf("UserID is required")
	}
	return nil
}
