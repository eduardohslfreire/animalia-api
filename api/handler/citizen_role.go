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
