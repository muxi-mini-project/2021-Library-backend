package model

//映射于数据库
type Summary struct {
	Id                  int64  `json:"id"`
	User_id             int64  `json:"user_id"`
	Book_id             int64  `json:"book_id"`
	Class_id            int64  `json:"class_id"`
	Title               string `json:"title"`
	Chapter             string `json:"chapter"`
	Summary_information string `json:"summary_information"`
	Thought             string `json:"thought"`
	Date                string `gorm:"type:timestamp;" json:"date"`
	Public              bool   `json:"public"`
}

type SummaryClass struct {
	Class_id   int64  `json:"class_id"`
	User_id    int64  `json:"user_id"`
	Class_name string `json:"class_name"`
}

type Book struct {
	Book_id          int64  `json:"book_id"`
	Book_name        string `json:"book_name"`
	Book_auther      string `json:"book_auther"`
	Book_information string `json:"book_information"`
	Book_picture     string `json:"book_picture"`
	Class_id         int64  `json:"class_id"`
	Click_sum        int    `json:"click_sum"`
}

type BookClass struct {
	Class_id      int64  `json:"class_id"`
	Class_name    string `json:"class_name"`
	Class_picture string `json:"class_picture"`
	Book_sum      int64  `json:"book_sum"`
}

//显示于页面
type SummaryRow struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Date   string `json:"date"`
	Public bool   `json:"public" `
}

type SummaryInfo struct {
	Title               string `json:"title"`
	Chapter             string `json:"chapter"`
	Summary_information string `json:"summary_information"`
	Thought             string `json:"thought"`
	Public              bool   `json:"public" `
}

type SummaryClassName struct {
	Name string `json:"name"`
}

type Search struct {
	Content string `json:"book_name" gorm:"book_name"`
}

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

type DigestAndClass struct {
	ClassId  string `json:"class_id" gorm:"class_id"`
	DigestId string `json:"digest_id" gorm:"digest_id"`
}

type UserAndBook struct {
	Id     string `json:"id" gorm:"id"`
	BookId string `json:"book_id" gorm:"book_id"`
	UserId string `json:"user_id" gorm:"user_id"`
}
type Token struct {
	Token string `json:"token"`
}
