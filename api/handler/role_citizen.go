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
func NewRoleCitizenHandler(r *gin.RouterGroup, roleUsecase usecase.IRoleUsecase) {
	handler := &RoleCitizenHandler{
		RoleUsecase: roleUsecase,
		Logger:      logger.NewLogger(),
	}
	r.GET("/roles/:role_id/citizens", handler.Find)
}

// Find ...
// @Summary      Find all citizens
// @Description  Find all citizens by role ID
// @Tags         roles
// @Produce      json
// @Param        roles_id  path      int  true  "Role ID"
// @Success      200       {object}  dto.CitizensBodyResponse
// @Failure      400       {object}  errors.ErrorResponse
// @Failure      404       {object}  errors.ErrorResponse
// @Failure      500       {object}  errors.ErrorResponse
// @Router       /roles/{roles_id}/citizens [get]
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
