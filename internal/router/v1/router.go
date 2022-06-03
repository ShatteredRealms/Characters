package v1

import (
    "github.com/ShatteredRealms/Accounts/pkg/middlewares"
    "github.com/ShatteredRealms/Accounts/pkg/model"
    "github.com/ShatteredRealms/Characters/internal/log"
    "github.com/ShatteredRealms/Characters/internal/option"
    "github.com/gin-gonic/gin"
    ginSwagger "github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"
    "gorm.io/gorm"
    "net/http"
)

func InitRouter(db *gorm.DB, config option.Config, logger log.LoggerService) (*gin.Engine, error) {

    if config.IsRelease() {
        gin.SetMode(gin.ReleaseMode)
    }

    router := gin.Default()
    router.Use(middlewares.ContentTypeMiddleWare())
    router.Use(middlewares.CORSMiddleWare())
    router.NoRoute(noRouteHandler())

    apiV1 := router.Group("/v1")
    SetHealthRoutes(apiV1)
    setupDocRouters(apiV1)

    logger.Log(log.Info, "Router fully initialized")

    return router, nil
}

func noRouteHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.JSON(http.StatusNotFound, model.NewGenericNotFoundResponse(c))
    }
}

func setupDocRouters(rg *gin.RouterGroup) {
    rg.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
