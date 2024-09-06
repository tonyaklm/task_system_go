package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"task_system_go/database"
	"task_system_go/models"
)

func CreatePost(c *gin.Context) {
	var post models.Post
	var payload CreatePostPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if userId, ok := c.Get("user_id"); !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization failed"})
		c.Abort()
		return
	} else {
		post.UserID = userId.(uint)
		post.Title = payload.Title
		post.Content = payload.Content
	}

	if err := post.CreatePost(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, gin.H{"post": post})
}

func UpdatePost(c *gin.Context) {
	var payload UpdatePostPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	var post models.Post
	if err := post.FindById(payload.PostId); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post does not exist"})
		c.Abort()
		return
	}

	post.Content = payload.Content
	post.Title = payload.Title
	database.Database.Save(&post)

	c.JSON(http.StatusOK, gin.H{"message": "Post was updated"})
}

func DeletePost(c *gin.Context) {
	userId, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization failed"})
		c.Abort()
		return
	}
	stringPostId := c.Param("post_id")
	postId, err := strconv.ParseUint(stringPostId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error  parsing post id"})
		c.Abort()
		return
	}

	var post models.Post
	if err := post.FindById(uint(postId)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post doesn't exist"})
		c.Abort()
		return
	}
	if post.UserID != userId {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not enough rights"})
		c.Abort()
		return
	}
	database.Database.Delete(&post)

	c.JSON(http.StatusOK, gin.H{"message": "Post was deleted"})
}

func GetPostById(c *gin.Context) {
	stringPostId := c.Param("post_id")
	postId, err := strconv.ParseUint(stringPostId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error  parsing post id"})
		c.Abort()
		return
	}

	var post models.Post
	if err := post.FindById(uint(postId)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post doesn't exist"})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, post)
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	res := database.Database.Find(&posts)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, posts)
}
