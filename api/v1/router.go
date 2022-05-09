package v1

import (
	"github.com/eduardohslfreire/animalia-api/api/handler"
	"github.com/eduardohslfreire/animalia-api/usecase"
	"github.com/gin-gonic/gin"
)

// InitRouting ...
func InitRouting(v1 *gin.RouterGroup, citizenUsecase usecase.ICitizenUsecase, roleUsecase usecase.IRoleUsecase) {
	ch := handler.NewCitizenHandler(citizenUsecase)

	v1.POST("/citizens", ch.Create)
	v1.GET("/citizens", ch.FindAllByFilter)
	v1.GET("/citizens/:citizen_id", ch.Find)
	v1.PUT("/citizens/:citizen_id", ch.Update)
	v1.DELETE("/citizens/:citizen_id", ch.Delete)

	crh := handler.NewCitizenRoleHandler(citizenUsecase)

	v1.GET("/citizens/:citizen_id/roles", crh.Find)
	v1.PUT("/citizens/:citizen_id/roles/:role_id", crh.Associate)
	v1.DELETE("/citizens/:citizen_id/roles/:role_id", crh.Disassociate)

	rh := handler.NewRoleHandler(roleUsecase)

	v1.GET("/roles", rh.FindAll)
	v1.GET("/roles/:role_id", rh.Find)

	rch := handler.NewRoleCitizenHandler(roleUsecase)

	v1.GET("/roles/:role_id/citizens", rch.Find)
}
