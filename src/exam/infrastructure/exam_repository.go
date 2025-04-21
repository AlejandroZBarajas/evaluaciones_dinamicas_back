package examInfrastructure

import (
	"database/sql"
	examEntity "evaluaciones/src/exam/domain/entity"
)

type ExamRepository struct {
	db *sql.DB
}

func NewExamRepository(db *sql.DB) *ExamRepository {
	return &ExamRepository{db: db}
}

func (repo *ExamRepository) CreateExam(exam *examEntity.ExamEntity) error {
	query := "INSERT INTO exams (name, total_questions, teacher_id, category_id) VALUES ($1, $2, $3, $4) RETURNING id"
	return repo.db.QueryRow(query, exam.Name, exam.TotalQuestions, exam.TeacherID, exam.CategoryId).Scan(&exam.ID)
}

func (repo *ExamRepository) GetAllByTeacherID(teacherID int32) (*[]examEntity.ExamEntity, error) {
	query := "SELECT id, name, total_questions, teacher_id, category_id FROM exams WHERE teacher_id = $1"
	rows, err := repo.db.Query(query, teacherID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exams []examEntity.ExamEntity
	for rows.Next() {
		var exam examEntity.ExamEntity
		if err := rows.Scan(&exam.ID, &exam.Name, &exam.TotalQuestions, &exam.TeacherID, &exam.CategoryId); err != nil {
			return nil, err
		}
		exams = append(exams, exam)
	}

	return &exams, nil
}

func (repo *ExamRepository) GetExamByID(examID int32) (examEntity.ExamEntity, error) {
	query := "SELECT id, name, total_questions, teacher_id, category_id FROM exams WHERE id = $1"
	var exam examEntity.ExamEntity
	err := repo.db.QueryRow(query, examID).Scan(&exam.ID, &exam.Name, &exam.TotalQuestions, &exam.TeacherID, &exam.CategoryId)
	return exam, err
}

func (repo *ExamRepository) UpdateExam(examID int32, updated *examEntity.ExamEntity) (*examEntity.ExamEntity, error) {
	query := "UPDATE exams SET name = $1, total_questions = $2, teacher_id = $3, category_id = $4 WHERE id = $5"
	_, err := repo.db.Exec(query, updated.Name, updated.TotalQuestions, updated.TeacherID, updated.CategoryId, examID)
	if err != nil {
		return nil, err
	}
	updated.ID = examID
	return updated, nil
}

func (repo *ExamRepository) DeleteExam(examID int32) error {
	query := "DELETE FROM exams WHERE id = $1"
	_, err := repo.db.Exec(query, examID)
	return err
}
