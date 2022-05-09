package handler

import (
	"net/http"

	"github.com/eduardohslfreire/animalia-api/api/dto"
	"github.com/eduardohslfreire/animalia-api/api/errors"
	"github.com/eduardohslfreire/animalia-api/api/validation"
	"github.com/eduardohslfreire/animalia-api/entity"
	dto_model "github.com/eduardohslfreire/animalia-api/entity/dto"
	"github.com/eduardohslfreire/animalia-api/pkg/logger"
	"github.com/eduardohslfreire/animalia-api/usecase"
	"github.com/eduardohslfreire/animalia-api/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// CitizenHandler ...
type CitizenHandler struct {
	CitizenUsecase usecase.ICitizenUsecase
	Logger         logger.GenericLogger
}

// NewCitizenHandler ...
func NewCitizenHandler(citizenUsecase usecase.ICitizenUsecase) *CitizenHandler {
	return &CitizenHandler{
		CitizenUsecase: citizenUsecase,
		Logger:         logger.NewLogger(),
	}
}

// Find ...
func (c *CitizenHandler) Find(ctx *gin.Context) {
	citizenID := new(dto.CitizenIDRequest)
	if err := ctx.ShouldBindUri(citizenID); err != nil {
		ctx.Error(err)
		return
	}
	citizen, err := c.CitizenUsecase.Find(citizenID.ID)
	if err != nil {
		ctx.Error(err)
		return
	}

	citizenResponse := new(dto.CitizenBodyResponse)
	copier.Copy(citizenResponse, citizen)

	ctx.JSON(http.StatusOK, citizenResponse)
}

// FindAllByFilter ...
func (c *CitizenHandler) FindAllByFilter(ctx *gin.Context) {
	citizenParams := new(dto.FindAllCitizensQueryParamsRequest)
	if err := ctx.ShouldBindQuery(citizenParams); err != nil {
		ctx.Error(errors.BadRequest("Invalid ID"))
		return
	}
	pagination := new(dto_model.Pagination)
	copier.Copy(pagination, citizenParams.PaginationRequest)

	filters := util.ExtractValidQueryParams(ctx.Request.URL, validation.FindAllCitizensValidParams)

	paginationCitizen, err := c.CitizenUsecase.FindAllByFilter(filters, *pagination)
	if err != nil {
		ctx.Error(err)
		return
	}

	paginationCitizensResponse := new(dto.PaginationCitizenBodyResponse)
	copier.Copy(paginationCitizensResponse, paginationCitizen)

	ctx.JSON(http.StatusOK, paginationCitizensResponse)
}

// Create ...
func (c *CitizenHandler) Create(ctx *gin.Context) {
	citizenDTO := new(dto.CreateCitizenBodyRequest)
	if err := ctx.ShouldBindJSON(citizenDTO); err != nil {
		ctx.Error(err)
		return
	}

	citizen := new(entity.Citizen)
	copier.Copy(citizen, citizenDTO)

	citizen, err := c.CitizenUsecase.Create(citizen)
	if err != nil {
		ctx.Error(err)
		return
	}

	citizenResponse := new(dto.CitizenBodyResponse)
	copier.Copy(citizenResponse, citizen)

	ctx.JSON(http.StatusCreated, citizenResponse)
}

// Update ...
func (c *CitizenHandler) Update(ctx *gin.Context) {
	citizenID := new(dto.CitizenIDRequest)
	if err := ctx.ShouldBindUri(citizenID); err != nil {
		ctx.Error(err)
		return
	}

	citizenDTO := new(dto.UpdateCitizenBodyRequest)
	if err := ctx.ShouldBindJSON(citizenDTO); err != nil {
		ctx.Error(err)
		return
	}

	citizen := new(entity.Citizen)
	copier.Copy(citizen, citizenDTO)

	citizen, err := c.CitizenUsecase.Update(citizenID.ID, citizen)
	if err != nil {
		ctx.Error(err)
		return
	}

	citizenResponse := new(dto.CitizenBodyResponse)
	copier.Copy(citizenResponse, citizen)

	ctx.JSON(http.StatusOK, citizenResponse)
}

// Delete ...
func (c *CitizenHandler) Delete(ctx *gin.Context) {
	citizenID := new(dto.CitizenIDRequest)
	if err := ctx.ShouldBindUri(citizenID); err != nil {
		ctx.Error(err)
		return
	}

	err := c.CitizenUsecase.Delete(citizenID.ID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
