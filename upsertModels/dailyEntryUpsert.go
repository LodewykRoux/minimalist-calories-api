package upsertmodels

import "time"

type DailyEntryUpsert struct {
	ID          uint
	Date        time.Time
	Quantity    float32
	FoodItemsId uint
}

type DailyEntryDelete struct {
	ID uint
}
