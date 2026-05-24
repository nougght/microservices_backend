package auth

import (
	"store-server/config"
	"store-server/internal/auth/handlers"
	"store-server/internal/auth/repositories"
	"store-server/internal/auth/services"
	"store-server/internal/auth/tools"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type AuthModule struct {
	AuthService     *services.AuthService
	FavItemsService *services.FavouriteItemsService

	authRepo     *repositories.AuthRepository
	favItemsRepo *repositories.FavouriteItemsRepository

	authHandler     *handlers.AuthHandler
	favItemsHandler *handlers.FavouriteItemsHandler
}

func NewAuthModule(cfg *config.Config, db *sqlx.DB, jwtTools *tools.JwtTools) *AuthModule {
	authRepo := repositories.NewAuthRepository(db)
	favItemsRepo := repositories.NewFavouriteItemsRepository(db)
	authService := services.NewAuthService(cfg, authRepo, jwtTools)
	favItemsService := services.NewFavouriteItemsService(favItemsRepo)

	return &AuthModule{
		authRepo:        authRepo,
		favItemsRepo:    favItemsRepo,
		AuthService:     authService,
		FavItemsService: favItemsService,
		authHandler:     handlers.NewAuthHandler(authService),
		favItemsHandler: handlers.NewFavouriteItemsHandler(favItemsService),
	}
}

func (m *AuthModule) RegisterPublicRoutes(r *gin.RouterGroup) {
	r.POST("/auth/login", m.authHandler.LogIn)
	r.POST("/auth/register", m.authHandler.Register)
	r.POST("/auth/refresh", m.authHandler.Refresh)
	r.POST("/auth/code/send", m.authHandler.SendCode)
	r.POST("/auth/code/verify", m.authHandler.VerifyCode)
	r.POST("/user/check/:email_or_phone", m.authHandler.CheckUserExists)
}
func (m *AuthModule) RegisterPrivateRoutes(r *gin.RouterGroup) {
	r.GET("/user/:user_id/favourites", m.favItemsHandler.GetFavouritesByUserID)
	r.POST("/user/:user_id/favourites", m.favItemsHandler.AddToFavourites)
	r.DELETE("/user/:user_id/favourites/:product_id", m.favItemsHandler.DeleteFromFavourites)

	r.DELETE("/user/:user_id", m.authHandler.DeleteUserByID)
	r.GET("/user/:user_id", m.authHandler.GetUser)
	r.GET("/user/:user_id/session", m.authHandler.GetUserSession)
	r.POST("/user/logout/:user_id", m.authHandler.LogOut)

	r.GET("/yandex-map-key", m.authHandler.GetYandexAPIKey)
}
