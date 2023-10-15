package entity


import (

"gorm.io/gorm"

)


type User struct {

gorm.Model
UserName string
FirstName string
LastName string
Email string
Address string
Phone string
Password string

}