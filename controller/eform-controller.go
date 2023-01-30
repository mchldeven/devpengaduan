package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/michaeldeven/microserviceuniversalbpr/dto"
	entity "github.com/michaeldeven/microserviceuniversalbpr/entity/models"
	"github.com/michaeldeven/microserviceuniversalbpr/helper"
	"github.com/michaeldeven/microserviceuniversalbpr/service"
	"net/http"
	"strconv"
)

type EformController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type eformController struct {
	eformService service.EformService
	jwtService   service.JWTService
}

func NewEformController(eformServ service.EformService, jwtServ service.JWTService) EformController {
	return &eformController{
		eformService: eformServ,
		jwtService:   jwtServ,
	}
}

func (c *eformController) All(context *gin.Context) {
	var eforms []entity.Eform = c.eformService.All()
	res := helper.BuildResponse(true, "OK", eforms)
	context.JSON(http.StatusOK, res)
}

func (c *eformController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var eform entity.Eform = c.eformService.FindByID(id)
	if (eform == entity.Eform{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", eform)
		context.JSON(http.StatusOK, res)
	}
}

func (c *eformController) Insert(context *gin.Context) {
	var eformCreateDTO dto.EformCreateDTO
	errDTO := context.ShouldBind(&eformCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Gagal memprosses Request! error!", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			eformCreateDTO.UserID = convertedUserID
		}
		result := c.eformService.Insert(eformCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *eformController) Update(context *gin.Context) {
	var eformUpdateDTO dto.EformUpdateDTO
	errDTO := context.ShouldBind(&eformUpdateDTO)
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
	if c.eformService.IsAllowedToEdit(userID, eformUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			eformUpdateDTO.UserID = id
		}
		result := c.eformService.Update(eformUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *eformController) Delete(context *gin.Context) {
	var eform entity.Eform
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	eform.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.eformService.IsAllowedToEdit(userID, eform.ID) {
		c.eformService.Delete(eform)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *eformController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
