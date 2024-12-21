package controller

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tamaco489/k6_load_test/api/jester/internal/gen"
)

func NewHJesterAPIServer() (*http.Server, error) {

	r := gin.New()

	// setup logging
	r.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		if http.StatusInternalServerError <= params.StatusCode {
			slog.ErrorContext(params.Request.Context(), params.ErrorMessage,
				"status", params.StatusCode,
				"method", params.Method,
				"path", params.Path,
				"ip", params.ClientIP,
				"latency_ms", params.Latency.Milliseconds(),
				"user_agent", params.Request.UserAgent(),
				"host", params.Request.Host,
			)
			return ""
		}
		if params.StatusCode >= http.StatusBadRequest && params.StatusCode <= http.StatusInternalServerError {
			slog.WarnContext(params.Request.Context(), params.ErrorMessage,
				"status", params.StatusCode,
				"method", params.Method,
				"path", params.Path,
				"ip", params.ClientIP,
				"latency_ms", params.Latency.Milliseconds(),
				"user_agent", params.Request.UserAgent(),
				"host", params.Request.Host,
			)
			return ""
		}
		slog.InfoContext(params.Request.Context(), "access",
			"status", params.StatusCode,
			"method", params.Method,
			"path", params.Path,
			"ip", params.ClientIP,
			"latency_ms", params.Latency.Milliseconds(),
			"user_agent", params.Request.UserAgent(),
			"host", params.Request.Host,
		)
		return ""
	}))

	cnf := cors.DefaultConfig()
	cnf.AllowOrigins = []string{"*"}
	cnf.AllowHeaders = append(cnf.AllowHeaders, "Authorization", "Access-Control-Allow-Origin")

	r.Use(cors.New(cnf))

	r.Use(gin.Recovery())

	// NOTE: envは一旦固定値で渡す
	env := "dev"
	apiController := NewControllers(env)
	strictServer := gen.NewStrictHandler(apiController, nil)

	gen.RegisterHandlersWithOptions(
		r,
		strictServer,
		gen.GinServerOptions{
			BaseURL: "/jester/",
			ErrorHandler: func(ctx *gin.Context, err error, i int) {
				_ = ctx.Error(err)
				ctx.JSON(i, gin.H{"msg": err.Error()})
			},
		},
	)

	// NOTE: portは一旦固定値で渡す
	port := "8080"
	server := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf(":%s", port),
	}

	return server, nil
}
