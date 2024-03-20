package model

type User struct{
	Account //the struct User inherits Account
	Name string `json:"name" gorm:"not null;type:string"`
	Surname string `json:"surname" gorm:"not null;type:string"`
	Picture string  `json:"picture" gorm:"type:string"` //probably an url to picture
	Biography string `json:"biography" gorm:"type:string"` //long text 
	Moto string  `json:"moto" gorm:"type:string"` //citation
}