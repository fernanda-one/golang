package main

import (
	"github.com/fernanda-one/golang_api/config"
	"github.com/fernanda-one/golang_api/controllers"
	"github.com/fernanda-one/golang_api/middleware"
	"github.com/fernanda-one/golang_api/repository"
	"github.com/fernanda-one/golang_api/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                   = config.ConnectDatabase()
	userRepository repository.UserRepository  = repository.NewUserRepository(db)
	bookRepository repository.BookRepository  = repository.NewBookRepository(db)
	jwtService     service.JWTService         = service.NewJWTService()
	userService    service.UserService        = service.NewUserService(userRepository)
	authService    service.AuthService        = service.NewAuthService(userRepository)
	bookService    service.BookService        = service.NewBookService(bookRepository)
	authController controllers.AuthController = controllers.NewAuthController(authService, jwtService)
	userController controllers.UserController = controllers.NewuserController(userService, jwtService)
	bookController controllers.BookController = controllers.NewBookController(bookService, jwtService)
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
	bookRoutes := r.Group("api/books", middleware.AuthorizeJWT(jwtService))
	{
		bookRoutes.GET("/", bookController.All)
		bookRoutes.POST("/", bookController.Insert)
		bookRoutes.GET("/:id", bookController.FindById)
		bookRoutes.PUT("/:id", bookController.Update)
		bookRoutes.DELETE("/:id", bookController.Delete)
	}
	r.Run("localhost:8088")

}
