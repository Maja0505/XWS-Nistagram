package model

type Admin struct {
	User `bson:",inline"`
}
