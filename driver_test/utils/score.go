//.package utils

//考试成绩
type ExamScore struct {
	Id    int    `db:"id"`
	Name  string `db:"name"`
	Score int    `db:"score"`
}