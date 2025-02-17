package brands

import (
	"relynolli_oil_selection/core/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ListBrands(c *fiber.Ctx) error {
	bs := services.NewBrandService()
	brands, err := bs.ListBrands()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(brands)
}

func ListModels(c *fiber.Ctx) error {
	bs := services.NewBrandService()
	brandId := c.Params("id")

	if brandId == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Brand ID is required",
		})
	}

	exists, err := bs.CheckBrandIsExists(brandId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	if !exists {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Brand not found",
		})
	}

	models, err := bs.ListModels(brandId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(models)
}

func ListSpecifications(c *fiber.Ctx) error {
	bs := services.NewBrandService()
	modelId := c.Params("id")

	if modelId == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Model ID is required",
		})
	}

	exists, err := bs.CheckModelIsExists(modelId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	if !exists {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Model not found",
		})
	}

	specifications, err := bs.ListSpecifications(modelId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(specifications)
}

func ListOil(c *fiber.Ctx) error {
	bs := services.NewBrandService()
	specificationId := c.Params("id")

	if specificationId == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Specification ID is required",
		})
	}

	exists, err := bs.CheckSpecificationIsExists(specificationId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	if !exists {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Specification not found",
		})
	}

	oils, err := bs.ListOil(specificationId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	for i := range oils {
		parameters, err := bs.ListOilParameters(strconv.Itoa(oils[i].Id))
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
			})
		}
		oils[i].Parameters = parameters
	}

	return c.Status(200).JSON(oils)
}
