package model

type Userinfo struct {
	UserId       string `json:"user_id" gorm:"AUTO_INCREMENT"`
	UserName     string `json:"user_name" gorm:"user_name"`
	UserPassword string `json:"user_password" gorm:"user_password"`
	UserPicture  string `json:"user_picture" gorm:"user_picture"`
	Motto        string
}

type BooksInfo struct {
	BookId      string `json:"book_id" gorm:"book_id"`
	BookName    string `json:"book_name" gorm:"book_name"`
	BookPicture string `json:"book_picture" gorm:"book_picture"`
	BookAuthor  string `json:"book_author" gorm:"book_author"`
	BookContent string `json:"book_information" gorm:"book_information"`
	BookClick   int    `json:"click_sum" gorm:"click_sum"`
	BookClass   string `json:"class_id" gorm:"class_id"`
}

type DigestInfo struct {
	DigestId    string `json:"id" gorm:"id"`
	UserId      string `json:"user_id" gorm:"user_id"`
	BookId      string `json:"book_id" gorm:"book_id"`
	ClassId     string `json:"class_id" gorm:"class_id"`
	DigestTitle string `json:"title" gorm:"title"`
	//摘录的原文
	DigestContent string `json:"summary_information" gorm:"summary"`
	DigestIdea    string `json:"thought" gorm:"thought"`
	DigestDate    string `json:"date" gorm:"date"`
	Public        bool   `json:"public" gorm:"public"`
}
