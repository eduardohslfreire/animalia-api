package main

import (
	"log"

	"github.com/eduardohslfreire/animalia-api/api"
	_ "github.com/eduardohslfreire/animalia-api/docs"
)

// @title           Animalia API
// @version         1.0.0
// @description     This is a sample server API to control animalia kingdom
// @termsOfService  http://swagger.io/terms/

// @contact.name   Eduardo Freire
// @contact.url    https://www.linkedin.com/in/edudufreire/
// @contact.email  eduardofreire1995@outlook.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:5000
// @BasePath  /api/v1
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	api := new(api.Server)
	api.StartServer()
}
