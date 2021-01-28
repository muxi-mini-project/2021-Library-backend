package model

type Userinfo struct {
	UserId       string `json:"user_id" gorm:"AUTO_INCREMENT"`
	UserName     string `json:"user_name" gorm:"user_name"`
	UserPassword string `json:"user_password" gorm:"user_password"`
	UserPicture  string `json:"user_picture" gorm:"user_picture"`
	Motto        string
}
