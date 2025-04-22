package studentExamInfrastructure

import (
	"database/sql"
	"encoding/json"
	studentExamEntity "evaluaciones/src/student_exam/domain/entity"
)

type StudentExamRepository struct {
	db *sql.DB
}

func NewStudentExamRepository(db *sql.DB) *StudentExamRepository {
	return &StudentExamRepository{db}
}

func (r *StudentExamRepository) CreateStudentExam(studentExam *studentExamEntity.StudentExamEntity) error {
	jsonResponses, err := json.Marshal(studentExam.Responses)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO student_exams (exam_id, student_id, total_questions, correct_answers, responses, submitted_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, submitted_at
	`

	err = r.db.QueryRow(query,
		studentExam.ExamID,
		studentExam.StudentID,
		studentExam.TotalQuestions,
		studentExam.CorrectAnswers,
		jsonResponses,
		studentExam.SubmittedAt,
	).Scan(&studentExam.StudentExamID, &studentExam.SubmittedAt)

	if err != nil {
		return err
	}

	return nil
}

func (r *StudentExamRepository) GetAllByExamID(examID int32) ([]*studentExamEntity.StudentExamEntity, error) {
	query := `
		SELECT id, exam_id, student_id, total_questions, correct_answers, responses, submitted_at
		FROM student_exams
		WHERE exam_id = $1
	`

	rows, err := r.db.Query(query, examID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*studentExamEntity.StudentExamEntity
	for rows.Next() {
		var e studentExamEntity.StudentExamEntity
		var jsonStr []byte

		if err := rows.Scan(&e.StudentExamID, &e.ExamID, &e.StudentID, &e.TotalQuestions, &e.CorrectAnswers, &jsonStr, &e.SubmittedAt); err != nil {
			return nil, err
		}

		json.Unmarshal(jsonStr, &e.Responses)
		result = append(result, &e)
	}

	return result, nil
}

func (r *StudentExamRepository) GetStudentExamByID(studentExamID int32) (*studentExamEntity.StudentExamEntity, error) {
	query := `
		SELECT id, exam_id, student_id, total_questions, correct_answers, responses, submitted_at
		FROM student_exams
		WHERE id = $1
	`

	var e studentExamEntity.StudentExamEntity
	var jsonStr []byte

	err := r.db.QueryRow(query, studentExamID).Scan(
		&e.StudentExamID, &e.ExamID, &e.StudentID, &e.TotalQuestions, &e.CorrectAnswers, &jsonStr, &e.SubmittedAt,
	)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(jsonStr, &e.Responses)
	return &e, nil
}

func (r *StudentExamRepository) DeleteStudentExam(studentExamID int32) error {
	query := `DELETE FROM student_exams WHERE id = $1`
	_, err := r.db.Exec(query, studentExamID)
	return err
}
