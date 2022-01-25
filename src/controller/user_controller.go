package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rpolnx.com.br/golang-with-ci/src/model/dto"
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

	list := dto.UserDtoListFromEntity(users)

	c.JSON(http.StatusOK, list)
}

func (ctrl *userController) GetOne(c *gin.Context) {
	id := c.Param("id")

	user, err := ctrl.userService.GetOneUser(id)

	if err != nil {
		util.HandleUnexpectedError(c, err)
	}

	userDto := dto.UserDtoFromEntity(*user)

	c.JSON(http.StatusOK, userDto)
}

func (ctrl *userController) Post(c *gin.Context) {
	userDto := new(dto.UserDTO)
	if err := c.ShouldBind(userDto); err != nil {
		c.JSON(http.StatusBadRequest, util.WrapHttpError(http.StatusBadRequest, "bad input", c.FullPath()))
		return
	}

	userEntity := userDto.ToEntity()

	userID, err := ctrl.userService.PostUser(userEntity)

	if err != nil {
		util.HandleUnexpectedError(c, err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": userID,
	})
}

func (ctrl *userController) Put(c *gin.Context) {
	id := c.Param("id")

	userDto := new(dto.UserDTO)
	if err := c.ShouldBind(userDto); err != nil {
		c.JSON(http.StatusBadRequest, util.WrapHttpError(http.StatusBadRequest, "bad input", c.FullPath()))
		return
	}

	userEntity := userDto.ToEntity()

	err := ctrl.userService.PutUser(id, userEntity)

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
