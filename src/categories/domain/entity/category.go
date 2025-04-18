package categoryEntity

type CategoryEntity struct {
	ID        int32  `json:"id"`
	Name      string `json:"name`
	TeacherID int32  `json:"teacher_id"`
}

func CreateCategory(name string, teacher_id int32) *CategoryEntity {
	return &CategoryEntity{Name: name, TeacherID: teacher_id}
}
