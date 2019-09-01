package controllers

import (
	"net/http"
	"strconv"

	"../constants"
	"../helpers"
	"../objects"
	"../services"
	"github.com/gin-gonic/gin"
)

type V1ItemController struct {
	itemService services.V1ItemService
	errorHelper helpers.ErrorHelper
}

func (handler *V1ItemController) GetByID(context *gin.Context) {

	id, err := strconv.ParseUint(context.Param("id"), 10, 32)

	if nil != err {
		handler.errorHelper.HTTPResponseError(context, err, constants.RequestParameterInvalid)
	}

	result, err := handler.itemService.GetByID(uint(id))
	if nil != err {
		handler.errorHelper.HTTPResponseError(context, err, constants.InternalServerError)
	}

	context.JSON(http.StatusOK, result)
}

func (handler *V1ItemController) Create(context *gin.Context) {

	requestObject := objects.V1ItemObjectRequest{}
	context.ShouldBind(&requestObject)

	result, err := handler.itemService.Create(requestObject)
	if nil != err {
		handler.errorHelper.HTTPResponseError(context, err, constants.InternalServerError)
	}

	context.JSON(http.StatusOK, result)
}

func (handler *V1ItemController) UpdateByID(context *gin.Context) {

	id, err := strconv.Atoi(context.Param("id"))
	if nil != err {
		handler.errorHelper.HTTPResponseError(context, err, constants.RequestParameterInvalid)
	}

	requestObject := objects.V1ItemObjectRequest{}
	context.ShouldBind(&requestObject)

	result, err := handler.itemService.UpdateByID(id, requestObject)
	if nil != err {
		handler.errorHelper.HTTPResponseError(context, err, constants.InternalServerError)
	}

	context.JSON(http.StatusOK, result)
}

func (handler *V1ItemController) GetByName(context *gin.Context) {

	name := context.Param("name")

	result, err := handler.itemService.GetByName(name)
	if nil != err {
		handler.errorHelper.HTTPResponseError(context, err, constants.InternalServerError)
	}

	context.JSON(http.StatusOK, result)
}
