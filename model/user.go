package model

type User struct {
	Id       int    `gorm:"primarykey" json:"id"`
	Username string `gorm:"type:varchar(100);not null;" json:"user_name"`
	Password string `gorm:"type:varchar(100);not null;" json:"user_pwd"`
	RoleId   int    `gorm:"type:varchar(100)" json:"role_Id"`
}

type Userlogin struct {
	Username string `json:"dlusername"`
	Password string `json:"dlpassword"`
}

type Registuser struct {
	Username  string `json:"zcusername"`
	Password  string `json:"zcpassword"`
	Password2 string `json:"zcpassword2"`
}

type Flag struct {
	Flag int `json:"flag"`
}
