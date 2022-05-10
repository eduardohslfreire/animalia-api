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
func NewCitizenHandler(r *gin.RouterGroup, citizenUsecase usecase.ICitizenUsecase) {
	handler := &CitizenHandler{
		CitizenUsecase: citizenUsecase,
	}
	r.POST("/citizens", handler.Create)
	r.GET("/citizens", handler.FindAllByFilter)
	r.GET("/citizens/:citizen_id", handler.Find)
	r.PUT("/citizens/:citizen_id", handler.Update)
	r.DELETE("/citizens/:citizen_id", handler.Delete)
}

// Find ...
// @Summary      Find citizens
// @Description  Find citizens by ID
// @Tags         citizens
// @Produce      json
// @Param        citizen_id  path      int                           true  "Citizen ID"
// @Success      200      {object}  dto.CitizenBodyResponse
// @Failure      400      {object}  errors.ErrorResponse
// @Failure      404         {object}  errors.ErrorResponse
// @Failure      500      {object}  errors.ErrorResponse
// @Router       /citizens/{citizen_id} [get]
func (c *CitizenHandler) Find(ctx *gin.Context) {
	citizenID := new(dto.CitizenIDRequest)
	if err := ctx.ShouldBindUri(citizenID); err != nil {
		ctx.Error(errors.BadRequest("Invalid ID"))
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
// @Summary      Find all citizens
// @Description  Find all citizens by filters
// @Tags         citizens
// @Produce      json
// @Param        name           query     string  false  "Citizen name"        example(Eduardo)
// @Param        species        query     string  false  "Species"             example(cat)
// @Param        description    query     string  false  "Description"         example(Happy Feline)
// @Param        has_pet_human  query     bool    false  "Has pet human"       example(false)
// @Param        sort           query     string  false  "Sort by"             default(id)
// @Param        page           query     int     false  "Number page"         default(1)
// @Param        limit          query     int     false  "Limit rows by page"  default(10
// @Success      200            {object}  dto.PaginationCitizenBodyResponse
// @Failure      400            {object}  errors.ErrorResponse
// @Failure      404            {object}  errors.ErrorResponse
// @Failure      500            {object}  errors.ErrorResponse
// @Router       /citizens [get]
func (c *CitizenHandler) FindAllByFilter(ctx *gin.Context) {
	citizenParams := new(dto.FindAllCitizensQueryParamsRequest)
	if err := ctx.ShouldBindQuery(citizenParams); err != nil {
		ctx.Error(err)
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
// @Summary      Create citizens
// @Description  Create a new citizens
// @Tags         citizens
// @Accept       json
// @Produce      json
// @Param        citizen  body      dto.CreateCitizenBodyRequest  true  "Citizen data"
// @Success      200         {object}  dto.CitizenBodyResponse
// @Failure      400         {object}  errors.ErrorResponse
// @Failure      500         {object}  errors.ErrorResponse
// @Router       /citizens [post]
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
// @Summary      Update citizens
// @Description  Update citizen by ID
// @Tags         citizens
// @Accept       json
// @Produce      json
// @Param        citizen_id  path      int  true  "Citizen ID"
// @Param        citizen     body      dto.UpdateCitizenBodyRequest  true  "Citizen data"
// @Success      200         {object}  dto.CitizenBodyResponse
// @Failure      400         {object}  errors.ErrorResponse
// @Failure      404         {object}  errors.ErrorResponse
// @Failure      500         {object}  errors.ErrorResponse
// @Router       /citizens/{citizen_id} [put]
func (c *CitizenHandler) Update(ctx *gin.Context) {
	citizenID := new(dto.CitizenIDRequest)
	if err := ctx.ShouldBindUri(citizenID); err != nil {
		ctx.Error(errors.BadRequest("Invalid ID"))
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
// @Summary      Delete citizens
// @Description  Delete citizen by ID
// @Tags         citizens
// @Accept       json
// @Produce      json
// @Param        citizen_id  path      int  true  "Citizen ID"
// @Success      204         {object}  dto.CitizenBodyResponse
// @Failure      400         {object}  errors.ErrorResponse
// @Failure      404         {object}  errors.ErrorResponse
// @Failure      500         {object}  errors.ErrorResponse
// @Router       /citizens/{citizen_id} [delete]
func (c *CitizenHandler) Delete(ctx *gin.Context) {
	citizenID := new(dto.CitizenIDRequest)
	if err := ctx.ShouldBindUri(citizenID); err != nil {
		ctx.Error(errors.BadRequest("Invalid ID"))
		return
	}

	err := c.CitizenUsecase.Delete(citizenID.ID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
