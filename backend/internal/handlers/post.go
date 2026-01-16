package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"social-media-scheduler/internal/dtos"
	"social-media-scheduler/internal/models"
	"social-media-scheduler/internal/queue"
	"social-media-scheduler/internal/services"
	"time"
)

type PostHandler struct {
	service   *services.PostService
	scheduler *queue.Scheduler
}

func NewPostHandler(svc *services.PostService, sch *queue.Scheduler) *PostHandler {
	return &PostHandler{service: svc, scheduler: sch}
}

func (h *PostHandler) Create(c *gin.Context) {
	userIdStr, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	var request dtos.CreatePostRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := uuid.Parse(userIdStr.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := &models.Post{
		ID:          uuid.New(),
		UserId:      userId,
		Title:       request.Title,
		Content:     request.Content,
		Channel:     request.Channel,
		ScheduledAt: *request.ScheduledAt,
		Status:      request.Status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = h.service.CreatePost(c.Request.Context(), post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.scheduler.AddPostQueue(c.Request.Context(), post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": map[string]string{
		"id": post.ID.String(),
	}})
	return
}

func (h *PostHandler) GetAll(c *gin.Context) {
	userIdStr, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	var query models.PostQuery
	err := c.ShouldBindQuery(&query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query.UserId = userIdStr.(string)

	users, err := h.service.GetAllPosts(c.Request.Context(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (h *PostHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	user, err := h.service.GetPostById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *PostHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var request dtos.UpdatePostRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isScheduleUpdated := false

	updates := map[string]interface{}{}

	if request.Title != nil {
		updates["title"] = *request.Title
	}

	if request.Content != nil {
		updates["content"] = *request.Content
	}

	if request.Status != nil {
		updates["status"] = *request.Status
	}

	if request.ScheduledAt != nil {
		isScheduleUpdated = true
		updates["scheduled_at"] = *request.ScheduledAt
	}

	err = h.service.UpdatePost(c.Request.Context(), id, updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if isScheduleUpdated {
		postId, err := uuid.Parse(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = h.scheduler.AddPostQueue(c.Request.Context(), &models.Post{
			ID:          postId,
			ScheduledAt: *request.ScheduledAt,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": map[string]string{
		"id": id,
	}})
}

func (h *PostHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeletePost(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.scheduler.RemovePostQueue(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully!"})
}
