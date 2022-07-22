package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mailru/easyjson"
	"net/http"
	"social_club/app/models"
	"social_club/app/usecases"
	"social_club/pkg/errors"
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
	var isCreated models.Created
	isCreated.CreatedInfo = true
	isCreatedJson, err := isCreated.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusCreated, "application/json; charset=utf-8", isCreatedJson)
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
