package routers

import (
	"time"

	"github.com/Afomiat/Loan-Tracker-API/config"
	"github.com/Afomiat/Loan-Tracker-API/delivery/controllers"
	"github.com/Afomiat/Loan-Tracker-API/infrastructure"
	"github.com/Afomiat/Loan-Tracker-API/repositories"
	"github.com/Afomiat/Loan-Tracker-API/usecases"
	"github.com/gin-gonic/gin"
)

func NewRouter(env *config.Env) *gin.Engine {
    router := gin.Default()

    // Initialize services
    tokenService := infrastructure.NewTokenService(
        env.AccessTokenSecret,
        env.RefreshTokenSecret,
        time.Duration(env.AccessTokenExpiryHour)*time.Hour,
        time.Duration(env.RefreshTokenExpiryHour)*time.Hour,
    )

    // Initialize repositories
    userRepo := repositories.NewUserRepository(env.MongoURI, env.DBName)

    // Initialize use cases
    userUsecase := usecases.UserUsecases{
        UserRepo:     userRepo,
        EmailService: &infrastructure.EmailService{
            SMTPHost:     env.SMTPHost,
            SMTPPort:     env.SMTPPort,
            SMTPUsername: env.SMTPUsername,
            SMTPPassword: env.SMTPPassword,
        },
        TokenService: tokenService,
    }

    // adminUsecase := usecases.AdminUsecases{
    //     UserRepo: userRepo,
    // }

    // Initialize controllers
    userController := controllers.UserController{
        UserUsecases: userUsecase,
    }

    // adminController := controllers.AdminController{
    //     AdminUsecases: adminUsecase,
    // }

    // Routes
    userRoutes := router.Group("/users")
    {
        userRoutes.POST("/register", userController.Register)
        userRoutes.GET("/verify-email", userController.VerifyEmail)
        // userRoutes.POST("/login", userController.Login)
        // userRoutes.POST("/token/refresh", userController.RefreshToken)
        // userRoutes.GET("/profile", userController.GetProfile)
        // userRoutes.POST("/password-reset", userController.RequestPasswordReset)
        // userRoutes.POST("/password-update", userController.UpdatePasswordAfterReset)
    }

    // adminRoutes := router.Group("/admin")
    // {
    //     adminRoutes.GET("/users", adminController.GetAllUsers)
    //     adminRoutes.DELETE("/users/:id", adminController.DeleteUser)
    // }

    return router
}