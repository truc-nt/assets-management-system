package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BaseHandler struct{}

func (h *BaseHandler) validateInput(c *gin.Context, input interface{}) error {
	if err := c.ShouldBindJSON(input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		log.Println(err.Error())
		return err
	}
	return nil
}

func (h *BaseHandler) handleError(c *gin.Context, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Not found",
		})
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Server Error",
		})
	}
	log.Println(err.Error())
}

func (h *BaseHandler) parseId(c *gin.Context, id string) uint32 {
	ID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return 0
	}
	return uint32(ID)
}

func (h *BaseHandler) handleSuccessGet(c *gin.Context, data interface{}) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, data)
}

func (h *BaseHandler) handleSuccessCreate(c *gin.Context) {
	c.Status(http.StatusCreated)
}

func (h *BaseHandler) handleSuccessUpdate(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (h *BaseHandler) handleSuccessDelete(c *gin.Context) {
	c.Status(http.StatusOK)
}
