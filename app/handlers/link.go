package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mailru/easyjson"
	"net/http"
	"ozon_test/app/models"
	"ozon_test/app/usecases"
	"ozon_test/pkg/errors"
)

type Handler struct {
	useCase usecases.UseCase
}

func MakeHandler(useCase_ usecases.UseCase) *Handler {
	return &Handler{useCase: useCase_}
}

func (handler *Handler) CreateMessage(c *gin.Context) {
	var msg models.Message
	err := easyjson.UnmarshalFromReader(c.Request.Body, &msg)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	err = handler.useCase.CreateMessage(&models.Node{Id: msg.FirstUser}, &models.Node{Id: msg.SecondUser})
	if err != nil {
		_ = c.Error(err)
		return
	}
	msgJson, err := msg.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusCreated, "application/json; charset=utf-8", msgJson)
}

func (handler *Handler) GetInformation(c *gin.Context) {

	information, err := handler.useCase.GetInformation()
	if err != nil {
		_ = c.Error(err)
		return
	}

	msgJson, err := information.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", msgJson)
}
