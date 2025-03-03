package upsertmodels

type FoodItemUpsert struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Uom      int     `json:"uoM"`
	Quantity float32 `json:"quantity"`
	Calories float32 `json:"calories"`
	Protein  float32 `json:"protein"`
	Carbs    float32 `json:"carbs"`
	Fat      float32 `json:"fat"`
}

type FoodItemDelete struct {
	ID uint
}
