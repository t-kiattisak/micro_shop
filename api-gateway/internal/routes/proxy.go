package routes

import (
	"api-gateway/internal/middlewares"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func SetupProxyRoutes(app *fiber.App) {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		panic("JWT_SECRET environment variable is not set")
	}

	publicAPIRoutes := app.Group("/api")
	api := publicAPIRoutes.Group("", middlewares.JWTMiddleware(secretKey))

	publicAPIRoutes.Get("/help", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "This is the API Gateway. Available routes: /orders/*, /inventory/*, /payment/*, /shipping/*",
		})
	})
	api.Get("/help", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "This is the API Gateway. Available routes: /orders/*, /inventory/*, /payment/*, /shipping/*",
		})
	})
	api.All("/orders/*", proxyTo("http://localhost:8081"))
	api.All("/inventory/*", proxyTo("http://localhost:8082"))
	api.All("/payment/*", proxyTo("http://localhost:8083"))
	api.All("/shipping/*", proxyTo("http://localhost:8084"))
}

func proxyTo(target string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := fasthttp.AcquireRequest()
		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseRequest(req)
		defer fasthttp.ReleaseResponse(resp)

		c.Request().CopyTo(req)

		original := c.OriginalURL()
		relativePath := original[len("/api"):]
		targetURL := target + relativePath
		req.SetRequestURI(targetURL)

		fmt.Printf("🔁 proxyTo → %s\n", targetURL)

		if auth := c.Get("Authorization"); auth != "" {
			req.Header.Set("Authorization", auth)
		}

		err := fasthttp.Do(req, resp)
		if err != nil {
			return c.Status(500).SendString("Proxy error: " + err.Error())
		}

		return c.Status(resp.StatusCode()).Send(resp.Body())
	}
}
