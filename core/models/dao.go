package models


type Brand struct {
	Id int `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Model struct {
	Id int `json:"id"`
	Name string `json:"name"`
	BrandId int `json:"brandId"`
}

type Specifiaction struct {
	Id int `json:"id"`
	Name string `json:"name"`
	ModelId int `json:"modelId" db:"model_id"`
	Year *string `json:"year" db:"year"`
	FuelType *string `json:"fuelType" db:"fuel_type"`
	EngineCapacity *int `json:"engineCapacity" db:"engine_capacity"`
	WheelDrive *string `json:"wheelDrive" db:"wheel_drive"`
	Power *int `json:"power" db:"power"`
}

type Oil struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Parameters interface{} `json:"parameters"`
}


