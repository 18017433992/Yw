package controllers

import (
	"fmt"
	"net/http"

	"example.com/blog/global"
	"example.com/blog/middleware"
	"example.com/blog/models"
	"github.com/gin-gonic/gin"
)

// 创建文章
func CreatePost(ctx *gin.Context) {
	var post models.Post

	userID, ok := middleware.Somehandler(ctx.Copy())

	if !ok {
		ctx.JSON(http.StatusInternalServerError, "用户不存在")
		return
	}
	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validatePost(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post.UserID = userID
	fmt.Println("创建文章时的", userID)
	if err := global.Db.AutoMigrate(&post); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := global.Db.Create(&post).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errosr": err.Error()})
		return
	}
	response := models.PostResponse{
		ID:        post.ID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		Title:     post.Title,
		Content:   post.Content,
		UserID:    post.UserID,
	}
	ctx.JSON(http.StatusCreated, response)
}

// 查找所有文章
func GetPostList(ctx *gin.Context) {
	var posts []models.Post
	if err := global.Db.Find(&posts).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var responses []models.PostResponse //实例化一个PostResponse 切片
	for _, post := range posts {        //循环posts切片
		reponse := models.PostResponse{ID: post.ID, //把posts切片的数据装到reponse
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
			Title:     post.Title,
			Content:   post.Content,
			UserID:    post.UserID}
		responses = append(responses, reponse) //把reponse的依次添加到reponses切片中

	}
	ctx.JSON(http.StatusOK, responses) //返回reponses切片

}

// 根据文章id查询文章信息
func GetPostById(ctx *gin.Context) {
	var post models.Post
	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := global.Db.Where("id = ?", post.ID).First(&post).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "文章不存在"})
		return
	}
	response := models.PostResponse{
		ID:        post.ID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		Title:     post.Title,
		Content:   post.Content,
		UserID:    post.UserID,
	}
	ctx.JSON(http.StatusOK, response)

}

// 根据文章id更新文章标题和内容
func UpdatePost(ctx *gin.Context) {
	var post models.Post
	var user models.Post
	var postInput models.Post

	userID, ok := middleware.Somehandler(ctx.Copy())
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "用户信息获取失败"})
		return
	}
	user.ID = userID //把传进来的userID赋值给user实例

	if err := ctx.ShouldBindJSON(&postInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//验证文章标题和内容必填
	if err := validatePost(&postInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//传进来的post从数据查找后赋值给post对象
	if err := global.Db.Where("id = ?", postInput.ID).First(&post).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "文章不存在"})
		return
	}
	//对比修改人和文章创建人是否一致
	if post.UserID == user.ID {
		if err := global.Db.Where("id = ?", post.ID).Find(&post).Updates(map[string]interface{}{
			"title":   postInput.Title,
			"content": postInput.Content,
		}).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": "fail", "message": "无修改权限"})
		return

	}
	response := models.PostResponse{
		ID:        post.ID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		Title:     post.Title,
		Content:   post.Content,
		UserID:    post.UserID,
	}
	ctx.JSON(http.StatusOK, response)
}

// 删除文章
func DeletePost(ctx *gin.Context) {
	var post models.Post
	var postInput models.Post

	var id uint
	userID, ok := middleware.Somehandler(ctx.Copy())
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "用户不存在"})
		return
	}
	id = userID
	if err := ctx.ShouldBindJSON(&postInput); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := global.Db.Where("id = ?", postInput.ID).First(&post).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "查不到该文章"})
		return
	}
	if post.UserID != id {
		ctx.JSON(http.StatusOK, gin.H{"error": "无修改权限"})
		return
	} else {
		if err := global.Db.Where("id = ?", postInput.ID).Delete(&postInput).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

	}
	ctx.JSON(http.StatusOK, gin.H{"code": "success", "message": "删除成功"})
}

// 自定义验证函数 文章标题和内容必填
func validatePost(post *models.Post) error {
	if post.Title == "" {
		return fmt.Errorf("title is required")
	}
	if post.Content == "" {
		return fmt.Errorf("content is required")
	}
	return nil
}
