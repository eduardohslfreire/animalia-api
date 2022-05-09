package handler

import (
	"net/http"

	"github.com/eduardohslfreire/animalia-api/api/dto"
	"github.com/eduardohslfreire/animalia-api/pkg/logger"
	"github.com/eduardohslfreire/animalia-api/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// RoleCitizenHandler ...
type RoleCitizenHandler struct {
	RoleUsecase usecase.IRoleUsecase
	Logger      logger.GenericLogger
}

// NewRoleCitizenHandler ...
func NewRoleCitizenHandler(roleUsecase usecase.IRoleUsecase) *RoleCitizenHandler {
	return &RoleCitizenHandler{
		RoleUsecase: roleUsecase,
		Logger:      logger.NewLogger(),
	}
}

// Find ...
func (r *RoleCitizenHandler) Find(ctx *gin.Context) {
	roleID := new(dto.RoleIDRequest)
	if err := ctx.ShouldBindUri(roleID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	citizens, err := r.RoleUsecase.FindAllCitizensByID(roleID.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	citizensResponse := new(dto.CitizensBodyResponse)
	copier.Copy(citizensResponse, citizens)

	ctx.JSON(200, citizensResponse)
}
