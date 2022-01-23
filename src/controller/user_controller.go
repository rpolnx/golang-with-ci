package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rpolnx.com.br/golang-with-ci/src/model/entities"
	"rpolnx.com.br/golang-with-ci/src/ports/in"
	"rpolnx.com.br/golang-with-ci/src/util"
)

type UserController interface {
	GetOne(c *gin.Context)
	GetAll(c *gin.Context)
	Post(c *gin.Context)
	Put(c *gin.Context)
	Delete(c *gin.Context)
}

type userController struct {
	userService in.UserUsecase
}

func (ctrl *userController) GetAll(c *gin.Context) {
	users, err := ctrl.userService.GetAllUsers()

	if err != nil {
		util.HandleUnexpectedError(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (ctrl *userController) GetOne(c *gin.Context) {
	id := c.Param("id")

	user, err := ctrl.userService.GetOneUser(id)

	if err != nil {
		util.HandleUnexpectedError(c, err)
	}

	c.JSON(http.StatusOK, user)
}

func (ctrl *userController) Post(c *gin.Context) {
	var dto entities.User
	if err := c.ShouldBind(&dto); err != nil {
		c.JSON(http.StatusBadRequest, util.WrapHttpError(http.StatusBadRequest, "bad input", c.FullPath()))
		return
	}

	user, err := ctrl.userService.PostUser(dto)

	if err != nil {
		util.HandleUnexpectedError(c, err)
	}

	c.JSON(http.StatusCreated, user)
}

func (ctrl *userController) Put(c *gin.Context) {
	id := c.Param("id")

	var dto entities.User
	if err := c.ShouldBind(&dto); err != nil {
		c.JSON(http.StatusBadRequest, util.WrapHttpError(http.StatusBadRequest, "bad input", c.FullPath()))
		return
	}

	err := ctrl.userService.PutUser(id, dto)

	if err != nil {
		util.HandleUnexpectedError(c, err)
	}

	c.Status(http.StatusAccepted)
}

func (ctrl *userController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := ctrl.userService.DeleteUser(id)

	if err != nil {
		util.HandleUnexpectedError(c, err)
	}

	c.Status(http.StatusNoContent)
}

func InitializeUserController(userService in.UserUsecase) UserController {
	return &userController{
		userService,
	}
}
