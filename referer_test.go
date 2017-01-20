package echo_middlewares

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestRefererTokenMiddleware(t *testing.T) {
	e := echo.New()
	
	req, _ := http.NewRequest(echo.GET, "/", nil)
	res := httptest.NewRecorder()
	var h echo.HandlerFunc
	
	/**
	 * Valid Credientials
	 */
	req.Header.Set("Authorization", "Bearer a582dd7904bb5d01ff3e63be3d07e873efac2207091b2135662515f6e5ab9")
	req.Header.Set("Referer", "http://example.com")
	c := e.NewContext(req, res)
	h = RefererTokenMiddleware(RefererMiddlewareConfig{
		Prefix: "Bearer",
		Credientials: []RefererCrediential{
			{"http://example.com", "a582dd7904bb5d01ff3e63be3d07e873efac2207091b2135662515f6e5ab9"},
		},
	})(func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	})
	assert.NoError(t, h(c))
	
	
	/**
	 * InValid Authorization
	 */
	req.Header.Set("Authorization", "")
	req.Header.Set("Referer", "http://example.com")
	c = e.NewContext(req, res)
	h = RefererTokenMiddleware(RefererMiddlewareConfig{
		Prefix: "Bearer",
		Credientials: []RefererCrediential{
			{"http://example.com", "a582dd7904bb5d01ff3e63be3d07e873efac2207091b2135662515f6e5ab9"},
		},
	})(func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	})
	assert.Error(t, h(c))
	
	/**
	 * InValid Referer
	 */
	req.Header.Set("Authorization", "Bearer a582dd7904bb5d01ff3e63be3d07e873efac2207091b2135662515f6e5ab9")
	req.Header.Set("Referer", "")
	c = e.NewContext(req, res)
	h = RefererTokenMiddleware(RefererMiddlewareConfig{
		Prefix: "Bearer",
		Credientials: []RefererCrediential{
			{"http://example.com", "a582dd7904bb5d01ff3e63be3d07e873efac2207091b2135662515f6e5ab9"},
		},
	})(func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	})
	assert.Error(t, h(c))
	
	/**
	 * Valid User WildCard Reffer
	 */
	req.Header.Set("Authorization", "Bearer a582dd7904bb5d01ff3e63be3d07e873efac2207091b2135662515f6e5ab9")
	req.Header.Set("Referer", "http://example.com/test")
	c = e.NewContext(req, res)
	h = RefererTokenMiddleware(RefererMiddlewareConfig{
		Prefix: "Bearer",
		Credientials: []RefererCrediential{
			{"http://example.com/*", "a582dd7904bb5d01ff3e63be3d07e873efac2207091b2135662515f6e5ab9"},
		},
	})(func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	})
	assert.NoError(t, h(c))
	
	/**
	 * Valid User WildCard Reffer
	 */
	req.Header.Set("Authorization", "Bearer a582dd7904bb5d01ff3e63be3d07e873efac2207091b2135662515f6e5ab9")
	req.Header.Set("Referer", "http://example.com/test")
	c = e.NewContext(req, res)
	h = RefererTokenMiddleware(RefererMiddlewareConfig{
		Prefix: "Bearer",
		Credientials: []RefererCrediential{
			{"", "a582dd7904bb5d01ff3e63be3d07e873efac2207091b2135662515f6e5ab9"},
		},
	})(func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	})
	assert.NoError(t, h(c))
	
	/**
	 * InValid User WildCard Reffer
	 */
	req.Header.Set("Authorization", "Bearer a582dd7904bb5d01ff3e63be3d07e873efac2207091b2135662515f6e5ab9")
	req.Header.Set("Referer", "http://example.com:8080/")
	c = e.NewContext(req, res)
	h = RefererTokenMiddleware(RefererMiddlewareConfig{
		Prefix: "Bearer",
		Credientials: []RefererCrediential{
			{"http://example.com/*", "a582dd7904bb5d01ff3e63be3d07e873efac2207091b2135662515f6e5ab9"},
		},
	})(func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	})
	assert.Error(t, h(c))
	
}