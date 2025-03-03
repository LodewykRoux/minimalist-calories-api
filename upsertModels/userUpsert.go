package upsertmodels

type UserUpsert struct {
	Name     string
	Email    string
	Password string
}

type UserLogin struct {
	Email    string
	Password string
}
