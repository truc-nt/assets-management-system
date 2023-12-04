package handlers

import (
	"fmt"
	"server/internal/models"
	"server/internal/services"

	"github.com/gin-gonic/gin"

	"errors"
	"log"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

var (
	funcGetUserHandler = func() services.IUserService {
		return services.NewUserService()
	}
)

type IUserHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
	GetUsers(c *gin.Context)
	GetUserById(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type UserHandler struct {
	BaseHandler
	Service services.IUserService
}

func NewUserHandler() IUserHandler {
	return &UserHandler{
		Service: funcGetUserHandler(),
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var registerReqParam models.UserRegisterParam
	if err := c.BindJSON(&registerReqParam); err != nil {
		return
	}

	err := h.Service.FindUserByUsername(registerReqParam.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		id, _err := h.Service.CreateUser(registerReqParam)
		if _err != nil {
			h.handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, "User created with id: "+strconv.FormatUint(uint64(id), 10))
		return
	}

	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusBadRequest, "User already exist")
}

func (h *UserHandler) Login(c *gin.Context) {
	var loginReqParam models.UserAuthParam
	if err := c.BindJSON(&loginReqParam); err != nil {
		return
	}

	user, err := h.Service.AuthenticateUser(loginReqParam)
	if err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccessLogin(c, user)
}

func (h *UserHandler) Logout(c *gin.Context) {
	var id struct{ Id uint32 }
	if err := c.BindJSON(&id); err != nil {
		return
	}

	err := h.Service.SetLogoutUser(id.Id)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, id)
}

func (h *UserHandler) handleError(c *gin.Context, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "User not exist",
		})
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Server Error",
		})
	}
	log.Println(err.Error())
}

func (h *UserHandler) handleSuccessLogin(c *gin.Context, user models.User) {
	err := h.Service.SetLoginUser(user.ID)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, struct {
		Id    uint32
		Role  uint32
		DName string
	}{
		Id:    user.ID,
		Role:  user.Role,
		DName: user.DName,
	})
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	var param models.GetUsersParam
	c.ShouldBindQuery(&param)

	assets, err := h.Service.GetUsers(&param)
	if err != nil {
		h.handleError(c, err)
		return
	}
	h.handleSuccessGet(c, &assets)
}

func (h *UserHandler) GetUserById(c *gin.Context) {
	id := h.parseId(c, c.Param("user_id"))
	if id == 0 {
		return
	}
	user, err := h.Service.GetUserById(id)
	if err != nil {
		h.handleError(c, err)
		return
	}
	h.handleSuccessGet(c, &user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	user := models.User{}

	if err := h.validateInput(c, &user); err != nil {
		return
	}

	id := h.parseId(c, c.Param("user_id"))
	if id == 0 {
		return
	}

	fmt.Println("cuu")

	if err := h.Service.UpdateUser(id, &user); err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccessUpdate(c)
}
