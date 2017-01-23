# Echo Custom Middlewares

[![CircleCI](https://circleci.com/gh/kaikoo-tech/echo_middlewares/tree/master.svg?style=svg)](https://circleci.com/gh/kaikoo-tech/echo_middlewares/tree/master)

## Referer
Like a GoogleConsoleAPI Feferer+AuthKey Aunthenticate Middleware

```
config := echo_middlewares.RefererMiddlewareConfig{
	Prefix: "Bearer",
	Credientials: []echo_middlewares.RefererCrediential {
		{"http://example.com/*", "a582dd7904bb5d01ff3e63be3d07e873efac2207091b2135662515f6e5ab9075"}
	}
}
e.Use(echo_middlewares.RefererTokenMiddleware(config))
```

## License
[MIT](https://github.com/labstack/echo/blob/master/LICENSE)