package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/michaeldeven/microserviceuniversalbpr/dto"
	entity "github.com/michaeldeven/microserviceuniversalbpr/entity/models"
	"github.com/michaeldeven/microserviceuniversalbpr/helper"
	"github.com/michaeldeven/microserviceuniversalbpr/service"
)

type PengaduanController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type pengaduanController struct {
	pengaduanService service.PengaduanService
	jwtService       service.JWTService
}

func NewPengaduanController(pengaduanServ service.PengaduanService, jwtServ service.JWTService) PengaduanController {
	return &pengaduanController{
		pengaduanService: pengaduanServ,
		jwtService:       jwtServ,
	}
}

func (c *pengaduanController) All(context *gin.Context) {
	var pengaduans []entity.Pengaduan = c.pengaduanService.All()
	res := helper.BuildResponse(true, "OK", pengaduans)
	context.JSON(http.StatusOK, res)
}

func (c *pengaduanController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var pengaduan entity.Pengaduan = c.pengaduanService.FindByID(id)
	if (pengaduan == entity.Pengaduan{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", pengaduan)
		context.JSON(http.StatusOK, res)
	}
}

func (c *pengaduanController) Insert(context *gin.Context) {
	var pengaduanCreateDTO dto.PengaduanCreateDTO
	errDTO := context.ShouldBind(&pengaduanCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Gagal memprosses Request! error!", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			pengaduanCreateDTO.UserID = convertedUserID
		}
		result := c.pengaduanService.Insert(pengaduanCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *pengaduanController) Update(context *gin.Context) {
	var pengaduanUpdateDTO dto.PengaduanUpdateDTO
	errDTO := context.ShouldBind(&pengaduanUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Gagal memprosses Request! error!", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.pengaduanService.IsAllowedToEdit(userID, pengaduanUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			pengaduanUpdateDTO.UserID = id
		}
		result := c.pengaduanService.Update(pengaduanUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *pengaduanController) Delete(context *gin.Context) {
	var pengaduan entity.Pengaduan
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	pengaduan.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.pengaduanService.IsAllowedToEdit(userID, pengaduan.ID) {
		c.pengaduanService.Delete(pengaduan)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *pengaduanController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
