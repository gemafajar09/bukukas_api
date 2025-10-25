package http

import (
	"go-project/internal/config"
	"go-project/internal/delivery/http/handler"
	"go-project/internal/delivery/middleware"
	"go-project/internal/repository"
	"go-project/internal/usecase"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type Dependencies struct {
	Logger      *zap.Logger
	UserUsecase usecase.UserUsecase
	CashUsecase usecase.CashUsecase
}

func initDeps(db *gorm.DB) (*Dependencies, error) {
	cfgzap := zap.NewProductionConfig()
	cfgzap.OutputPaths = []string{"app.log", "stdout"}
	logger, err := cfgzap.Build()
	if err != nil {
		return nil, err
	}

	userUsecase := usecase.NewUserUsecase(repository.NewUserRepository(db))
	cashUsecase := usecase.NewCashUsecase(repository.NewCashRepository(db))

	return &Dependencies{
		Logger:      logger,
		UserUsecase: userUsecase,
		CashUsecase: cashUsecase,
	}, nil
}

func NewRouter(cfg config.Config, db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	deps, err := initDeps(db)
	if err != nil {
		panic(err)
	}
	defer deps.Logger.Sync()

	r.Use(middleware.LoggerMiddleware(deps.Logger))

	authHandler := handler.NewAuthHandler(deps.UserUsecase)
	userHandler := handler.NewUserHandler(deps.UserUsecase)
	cashHandler := handler.NewCashHandler(deps.CashUsecase)

	// Auth router group
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}

	// API router group
	apiGroup := r.Group("/api")
	apiGroup.Use(middleware.JWTAuthMiddleware(cfg.JWT))
	{
		// user router
		apiGroup.GET("/profile", userHandler.GetProfile)
		apiGroup.GET("/get-users", userHandler.GetUsers)

		// cash router
		apiGroup.POST("/cash/transactions", cashHandler.CreateTransaction)
		apiGroup.GET("/cash/transactions", cashHandler.GetTransactions)
		apiGroup.GET("/cash/balance", cashHandler.GetBalance)
		apiGroup.GET("/cash/categories", cashHandler.GetCategories)
	}

	return r
}
