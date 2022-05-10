package handler

import (
	"github.com/eduardohslfreire/animalia-api/api/dto"
	"github.com/eduardohslfreire/animalia-api/api/errors"
	"github.com/eduardohslfreire/animalia-api/pkg/logger"
	"github.com/eduardohslfreire/animalia-api/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// CitizenRoleHandler ...
type CitizenRoleHandler struct {
	CitizenUsecase usecase.ICitizenUsecase
	Logger         logger.GenericLogger
}

// NewCitizenRoleHandler ...
func NewCitizenRoleHandler(r *gin.RouterGroup, citizenUsecase usecase.ICitizenUsecase) {
	handler := &CitizenRoleHandler{
		CitizenUsecase: citizenUsecase,
		Logger:         logger.NewLogger(),
	}
	r.GET("/citizens/:citizen_id/roles", handler.Find)
	r.PUT("/citizens/:citizen_id/roles/:role_id", handler.Associate)
	r.DELETE("/citizens/:citizen_id/roles/:role_id", handler.Disassociate)
}

// Find ...
// @Summary      Find roles
// @Description  Find all roles by citizen ID
// @Tags         citizens
// @Produce      json
// @Param        citizen_id  path      int  true  "Citizens ID"
// @Success      200         {object}  dto.RolesBodyResponse
// @Failure      400         {object}  errors.ErrorResponse
// @Failure      404         {object}  errors.ErrorResponse
// @Failure      500         {object}  errors.ErrorResponse
// @Router       /citizens/{citizen_id}/roles [get]
func (cr *CitizenRoleHandler) Find(ctx *gin.Context) {
	citizenID := new(dto.CitizenIDRequest)
	if err := ctx.ShouldBindUri(citizenID); err != nil {
		ctx.Error(errors.BadRequest("Invalid ID"))
		return
	}

	roles, err := cr.CitizenUsecase.FindAllRolesByID(citizenID.ID)
	if err != nil {
		ctx.Error(err)
		return
	}

	rolesResponse := make(dto.RolesBodyResponse, 0)
	copier.Copy(&rolesResponse, roles)

	ctx.JSON(200, rolesResponse)
}

// Associate ...
// @Summary      Associate
// @Description  Associate a citizen with a role
// @Tags         citizens
// @Accept       json
// @Produce      json
// @Param        citizen_id  path      int     true  "Citizen ID"
// @Param        role_id     path      int     true  "Role ID"
// @Success      204         {string}  string  "No content"
// @Failure      400         {object}  errors.ErrorResponse
// @Failure      404         {object}  errors.ErrorResponse
// @Failure      422         {object}  errors.ErrorResponse
// @Failure      500         {object}  errors.ErrorResponse
// @Router       /citizens/{citizen_id}/roles/{role_id} [put]
func (cr *CitizenRoleHandler) Associate(ctx *gin.Context) {
	citizenRoleID := new(dto.CitizenRoleIDRequest)
	if err := ctx.ShouldBindUri(citizenRoleID); err != nil {
		ctx.Error(errors.BadRequest("Invalid ID"))
		return
	}

	err := cr.CitizenUsecase.AssociateRole(citizenRoleID.CitizenID, citizenRoleID.RoleID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.Status(204)
}

// Disassociate ...
// @Summary      Disassociate
// @Description  Disassociates a citizen from a role
// @Tags         citizens
// @Accept       json
// @Produce      json
// @Param        citizen_id  path      int     true  "Citizen ID"
// @Param        role_id     path      int     true  "Role ID"
// @Success      204         {string}  string  "No content"
// @Failure      400         {object}  errors.ErrorResponse
// @Failure      404         {object}  errors.ErrorResponse
// @Failure      422         {object}  errors.ErrorResponse
// @Failure      500         {object}  errors.ErrorResponse
// @Router       /citizens/{citizen_id}/roles/{role_id} [delete]
func (cr *CitizenRoleHandler) Disassociate(ctx *gin.Context) {
	citizenRoleID := new(dto.CitizenRoleIDRequest)
	if err := ctx.ShouldBindUri(citizenRoleID); err != nil {
		ctx.Error(errors.BadRequest("Invalid ID"))
		return
	}

	err := cr.CitizenUsecase.DisassociateRole(citizenRoleID.CitizenID, citizenRoleID.RoleID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.Status(204)
}
