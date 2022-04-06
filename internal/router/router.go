package router

import (
	appConfig "github.com/babulal107/go-kubernetes-poc/internal/config"
	"github.com/babulal107/go-kubernetes-poc/pkg/api/model"
	"github.com/babulal107/go-kubernetes-poc/pkg/middleware"
	"github.com/babulal107/go-kubernetes-poc/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/martian/log"
	"net/http"
)

// Init services, repositories, gin middleware and router
func Init(configs *appConfig.Application) *gin.Engine {

	r := gin.New()
	r.HandleMethodNotAllowed = utils.HandleMethodNotAllowed
	r.UseRawPath = utils.UseRawPath
	r.Use(gin.Recovery(), gin.Logger()) // default gin recovery and logger middleware used
	r.Use(func(context *gin.Context) {
		log.Debugf("%v", context.Request.Header)
		log.Debugf("Service Name %s", context.GetHeader(utils.HttpHeaderServiceName))
		context.Next()
	})
	r.Use(middleware.WithCacheHeaderControl("300")) // 5 mines

	// server status check route
	r.GET("/health", func(c *gin.Context) {

		serviceInfo := make(map[string]string)
		serviceInfo["app-name"] = configs.APP.GetName()
		serviceInfo["app-version"] = configs.APP.GetVersion()
		serviceInfo["app-log-level"] = configs.APP.GetLogLevel()

		c.JSON(http.StatusOK, model.Response{
			Code:    http.StatusOK,
			Message: "SERVICE UP",
			Data:    serviceInfo,
		})
	})

	return r
}
