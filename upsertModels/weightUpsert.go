package upsertmodels

import "time"

type WeightUpsert struct {
	ID     uint
	Date   time.Time
	Weight float32
}

type WeightDelete struct {
	ID uint
}
