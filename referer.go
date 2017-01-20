package echo_middlewares

import (
	"fmt"
	_ "net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	glob "github.com/ryanuber/go-glob"
	
)

type RefererCrediential struct {
	Referer string `json:"Referer"`
	Token string `json:"token"`
}

type RefererMiddlewareConfig struct {
	Skipper middleware.Skipper
	Header string `json:"header"`
	Prefix string `json:"prefix"`
	Credientials []RefererCrediential `json:"credientials"`
}

func RefererTokenMiddleware(config RefererMiddlewareConfig) echo.MiddlewareFunc {
	// Defaults
	if config.Skipper == nil {
		config.Skipper = func(c echo.Context) bool {
			return false
		}
	}
	if config.Header == "" {
		config.Header = "Authorization"
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}
			if config.Credientials == nil {
				return next(c)
			}
			req := c.Request()
			for _, cred := range config.Credientials {
				token := func(conf RefererMiddlewareConfig, cred RefererCrediential) string {
					return fmt.Sprintf("%s %s", config.Prefix, cred.Token)
				}
				if token(config, cred) == req.Header.Get(config.Header) {
					if cred.Referer == "" {
						return next(c)
					} else {
						if req.Referer() == "" {
							return echo.ErrUnauthorized
						}
						if glob.Glob(cred.Referer, req.Referer()) {
							return next(c)
						}
					}
				}
			}
			return echo.ErrUnauthorized
		}
	}
}
