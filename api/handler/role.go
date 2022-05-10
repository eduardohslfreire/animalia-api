package handler

import (
	"github.com/eduardohslfreire/animalia-api/api/dto"
	"github.com/eduardohslfreire/animalia-api/api/errors"
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
func NewRoleHandler(r *gin.RouterGroup, roleUsecase usecase.IRoleUsecase) {
	handler := &RoleHandler{
		RoleUsecase: roleUsecase,
		Logger:      logger.NewLogger(),
	}
	r.GET("/roles", handler.FindAll)
	r.GET("/roles/:role_id", handler.Find)
}

// Find ...
func (r *RoleHandler) Find(ctx *gin.Context) {
	roleID := new(dto.RoleIDRequest)
	if err := ctx.ShouldBindUri(roleID); err != nil {
		ctx.Error(errors.BadRequest("Invalid ID"))
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
