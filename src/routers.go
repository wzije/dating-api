package src

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Routers(router *fiber.App) {
	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Dating API!")
	})

	router.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	//inject
	db := DB()
	database := NewDatabase(db)
	handler := NewHandler(database)

	//router
	api := router.Group("/api/v1")
	api.Get("/register", handler.Register)
	api.Get("/login", handler.Login)

	//for authorize user
	api.Get("/home", AuthMiddleware, handler.Home)

}

// Authentication middleware
func AuthMiddleware(ctx *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		ContextKey:  "jwtKey",
		KeyFunc:     JwtKeyFunc(),
		TokenLookup: "header:Authorization",
		Claims:      new(jwt.MapClaims),
		AuthScheme:  "Bearer",
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return Json(ctx, Response{
				Code:    fiber.StatusUnauthorized,
				Message: "forbidden access",
			})
		},
		SuccessHandler: func(ctx *fiber.Ctx) error {
			//parsing claim to payload data, so this can be access anywhere
			ParsePayload(ctx)
			return ctx.Next()
		},
	})(ctx)
}
