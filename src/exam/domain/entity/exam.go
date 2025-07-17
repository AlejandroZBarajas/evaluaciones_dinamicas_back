package examEntity

type ExamEntity struct {
	ID             int32  `json:"id"`
	Name           string `json:"name"`
	TotalQuestions int32  `json:"total_questions"`
	TeacherID      int32  `json:"teacher_id"`
	CategoryId     int32  `json:"category_id"`
}

func CreateExam(name string, totalQ int32, teacherId int32, catID int32) *ExamEntity {
	return &ExamEntity{Name: name, TotalQuestions: totalQ, TeacherID: teacherId, CategoryId: catID}
}
