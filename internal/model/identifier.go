package model

type Identifier struct {
	Id               int    `json:"id" db:"id"`
	StringIdentifier string `json:"string_identifier" db:"string_identifier" binding:"required"`
	Description      string `json:"description" db:"description" binding:"required"`
}
