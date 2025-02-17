package brands

import (
	"github.com/gofiber/fiber/v2"
)


func SetupRoutes(api fiber.Router) {
	api.Get("/brands", ListBrands)
	api.Get("/brands/:id/models", ListModels)
	api.Get("/brands/:brandId/models/:id/specifications", ListSpecifications)
	api.Get("/brands/:brandId/models/:modelId/specifications/:id/oils", ListOil)
}

