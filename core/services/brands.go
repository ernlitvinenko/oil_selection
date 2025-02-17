package services

import (
	"relynolli_oil_selection/core/internal_services/database"
	"relynolli_oil_selection/core/models"

	"github.com/gofiber/fiber/v2"
)

type brandService struct { }

type BrandsService interface {
	ListBrands() ([]models.Brand, error)
	ListModels(brandId string) ([]models.Model, error)
	CheckBrandIsExists(brandId string) (bool, error)
	CheckModelIsExists(modelId string) (bool, error)
	ListSpecifications(modelId string) ([]models.Specifiaction, error)
	CheckSpecificationIsExists(specificationId string) (bool, error)
	ListOil(specificationId string) ([]models.Oil, error)
	ListOilParameters(oilId string) (*fiber.Map, error)
}

func NewBrandService() BrandsService {
	return &brandService{}
}

func (s *brandService) ListBrands() ([]models.Brand, error) { 
	brands := []models.Brand{}

	err := database.DB.Select(&brands, "SELECT id, name FROM lst_value where lst_id = 1 order by name")
	if err != nil {
		return nil, err
	}
	return brands, nil
}

func (s *brandService) ListModels(brandId string) ([]models.Model, error) {
	models := []models.Model{}

	err := database.DB.Select(&models, "SELECT lv.id as id, lv.name as name, bm.brand_id as brandId FROM lst_value lv join brand_model bm on bm.model_id = lv.id where lv.lst_id = 2 and bm.brand_id = $1 order by name", brandId)
	if err != nil {
		return nil, err
	}
	return models, nil
 }

 func (s *brandService) CheckBrandIsExists(brandId string) (bool, error) {
	brand := models.Brand{}

	err := database.DB.Get(&brand, "SELECT id FROM lst_value where lst_id = 1 and id = $1", brandId)
	if err != nil {
		return false, err
	}
	return true, nil
 }

 func (s *brandService) CheckModelIsExists(modelId string) (bool, error) {
	model := models.Model{}

	err := database.DB.Get(&model, "SELECT id FROM lst_value where lst_id = 2 and id = $1", modelId)
	if err != nil {
		return false, err
	}
	return true, nil
 }

 func (s *brandService) ListSpecifications(modelId string) ([]models.Specifiaction, error) {
	specifications := []models.Specifiaction{}

	err := database.DB.Select(&specifications, "SELECT * FROM auto_specifications where model_id = $1", modelId)
	if err != nil {
		return nil, err
	}
	return specifications, nil
 }

 func (s *brandService) CheckSpecificationIsExists(specificationId string) (bool, error) {
	specification := models.Specifiaction{}

	err := database.DB.Get(&specification, "SELECT id FROM auto_specifications where id = $1", specificationId)
	if err != nil {
		return false, err
	}
	return true, nil
 }

 func (s *brandService) ListOil(specificationId string) ([]models.Oil, error) {
	oil := []models.Oil{}

	err := database.DB.Select(&oil, "SELECT oil.id, oil.name FROM lst_value oil join specification_motor_oil smo on smo.oil_id = oil.id where smo.spec_id = $1 and oil.lst_id = 8", specificationId)
	
	if err != nil {
		return nil, err
	}
	return oil, nil
 }

 func (s *brandService) ListOilParameters(oilId string) (*fiber.Map, error) {
	parameters := fiber.Map{}

	rows, err := database.DB.Query("SELECT name, val FROM lst_params where value_id = $1", oilId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var value string
		err = rows.Scan(&name, &value)
		if err != nil {
			return nil, err
		}
		parameters[name] = value
	 }
	return &parameters, nil
 }