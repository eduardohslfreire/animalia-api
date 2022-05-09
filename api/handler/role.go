package handler

import (
	"github.com/eduardohslfreire/animalia-api/api/dto"
	"github.com/eduardohslfreire/animalia-api/pkg/logger"
	"github.com/eduardohslfreire/animalia-api/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// RoleHandler ...
type RoleHandler struct {
	RoleUsecase usecase.IRoleUsecase
	Logger      logger.GenericLogger
}

// NewRoleHandler ...
func NewRoleHandler(roleUsecase usecase.IRoleUsecase) *RoleHandler {
	return &RoleHandler{
		RoleUsecase: roleUsecase,
		Logger:      logger.NewLogger(),
	}
}

// Find ...
func (r *RoleHandler) Find(ctx *gin.Context) {
	roleID := new(dto.RoleIDRequest)
	if err := ctx.ShouldBindUri(roleID); err != nil {
		ctx.Error(err)
		return
	}
	role, err := r.RoleUsecase.FindByID(roleID.ID)
	if err != nil {
		ctx.Error(err)
		return
	}

	roleResponse := new(dto.RoleBodyResponse)
	copier.Copy(roleResponse, role)

	ctx.JSON(200, roleResponse)
}

// FindAll ...
func (r *RoleHandler) FindAll(ctx *gin.Context) {
	roles, err := r.RoleUsecase.FindAll()
	if err != nil {
		ctx.Error(err)
		return
	}

	rolesResponse := new(dto.RolesBodyResponse)
	copier.Copy(rolesResponse, roles)

	ctx.JSON(200, rolesResponse)
}
