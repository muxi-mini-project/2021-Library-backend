package model

//映射于数据库
type Summary struct {
	Id int64 `json:"id"`
	User_id int64 `json:"user_id"`
	Book_id int64 `json:"book_id"`
	Class_id int64 `json:"class_id"`
	Title string `json:"title"`
	Chapter string `json:"chapter"`
	Summary_information string `json:"summary_information"`
	Thought string `json:"thought"`
	Date string `gorm:"type:timestamp;" json:"date"`
	Public bool `json:"public"`
}

type SummaryClass struct {
	Class_id int64 `json:"class_id"`
	User_id int64 `json:'user_id'`
	Class_name string `json:"class_name"`
}

type Book struct {
	Book_id int64 `json:"book_id"`
	Book_name string `json:"book_name"`
	Book_auther string `json:"book_auther"`
	Book_information string `json:"book_information"`
	Book_picture string `json:"book_picture"`
	Class_id int64 `json:"class_id"`
	Click_sum int `json:"click_sum"`
}

type BookClass struct {
	Class_id int64 `json:"class_id"`
	Class_name string `json:"class_name"`
	Class_picture string `json:"class_picture"`
	Book_sum int64 `json:"book_sum"`
}

//显示于页面
type SummaryRow struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	Date string `json:"date"`
	Public bool `json:"public" `
}

type SummaryInfo struct {
	Title string `json:"title"`
	Chapter string `json:"chapter"`
	Summary_information string `json:"summary_information"`
	Thought string `json:"thought"`
	Public bool `json:"public" `
}

type SummaryClassName struct {
	Name string `json:"name"`
}

type Search struct {
	Word string `json:"word"`
}

