package main

import (
	"github.com/gin-gonic/gin"
	"github.com/michaeldeven/microserviceuniversalbpr/config"
	"github.com/michaeldeven/microserviceuniversalbpr/controller"
	"github.com/michaeldeven/microserviceuniversalbpr/middleware"
	"github.com/michaeldeven/microserviceuniversalbpr/repository"
	"github.com/michaeldeven/microserviceuniversalbpr/service"
	"gorm.io/gorm"
)

var (
	db                  *gorm.DB                       = config.SetupDatabaseConnection()
	userRepository      repository.UserRepository      = repository.NewUserRepository(db)
	pengaduanRepository repository.PengaduanRepository = repository.NewPengaduanRepository(db)
	eformRepository     repository.EformRepository     = repository.NewEformRepository(db)
	jwtService          service.JWTService             = service.NewJWTService()
	userService         service.UserService            = service.NewUserService(userRepository)
	pengaduanService    service.PengaduanService       = service.NewPengaduanService(pengaduanRepository)
	eformService        service.EformService           = service.NewEformService(eformRepository)
	authService         service.AuthService            = service.NewAuthService(userRepository)
	authController      controller.AuthController      = controller.NewAuthController(authService, jwtService)
	userController      controller.UserController      = controller.NewUserController(userService, jwtService)
	pengaduanController controller.PengaduanController = controller.NewPengaduanController(pengaduanService, jwtService)
	eformController     controller.EformController     = controller.NewEformController(eformService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	PengaduanRoutes := r.Group("api/pengaduan", middleware.AuthorizeJWT(jwtService))
	{
		PengaduanRoutes.GET("/", pengaduanController.All)
		PengaduanRoutes.POST("/", pengaduanController.Insert)
		PengaduanRoutes.GET("/:id", pengaduanController.FindByID)
		PengaduanRoutes.PUT("/:id", pengaduanController.Update)
		PengaduanRoutes.DELETE("/:id", pengaduanController.Delete)
	}

	EformRoutes := r.Group("api/eform", middleware.AuthorizeJWT(jwtService))
	{
		EformRoutes.GET("/", eformController.All)
		EformRoutes.POST("/", eformController.Insert)
		EformRoutes.GET("/:id", eformController.FindByID)
		EformRoutes.PUT("/:id", eformController.Update)
		EformRoutes.DELETE("/:id", eformController.Delete)
	}

	r.Run()
}
