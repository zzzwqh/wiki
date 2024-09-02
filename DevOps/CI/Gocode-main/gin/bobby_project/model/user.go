package model

type Account struct {
	Name   string `form:"name" json:"name" binding:"required"`
	Passwd string `form:"passwd" json:"passwd" binding:"required,min=3"`
}
