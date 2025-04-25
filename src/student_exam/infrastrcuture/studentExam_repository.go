package studentExamInfrastructure

import (
	"database/sql"
	"encoding/json"
	studentExamEntity "evaluaciones/src/student_exam/domain/entity"

	"fmt"
	"time"
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

func (r *StudentExamRepository) EvaluateStudentExam(submission *studentExamEntity.ExamSubmissionInput) (*studentExamEntity.StudentExamResult, error) {

	rows, err := r.db.Query(`
		SELECT id, question_data
		FROM questions
		WHERE exam_id = ?`, submission.Responses[0].QuestionID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener las preguntas: %w", err)
	}
	defer rows.Close()

	var questions []struct {
		ID            int32
		CorrectAnswer string
	}

	for rows.Next() {
		var q struct {
			ID            int32
			CorrectAnswer string
		}
		var questionDataJSON []byte
		if err := rows.Scan(&q.ID, &questionDataJSON); err != nil {
			return nil, fmt.Errorf("error al escanear la pregunta: %w", err)
		}

		var qData map[string]interface{}
		if err := json.Unmarshal(questionDataJSON, &qData); err != nil {
			return nil, fmt.Errorf("error al deserializar la pregunta: %w", err)
		}

		if ans, ok := qData["answer"].(string); ok {
			q.CorrectAnswer = ans
		}
		questions = append(questions, q)
	}

	total := int32(len(questions))
	var correct int32 = 0
	for _, q := range questions {
		for _, response := range submission.Responses {
			if response.QuestionID == q.ID && fmt.Sprint(response.Answer) == q.CorrectAnswer {
				correct++
			}
		}
	}

	score := fmt.Sprintf("%02d/%02d", correct, total)
	submittedAt := time.Now()

	_, err = r.db.Exec(`
		INSERT INTO results (student_exam_id, result, submitted_at)
		VALUES (?, ?, ?)`,
		submission.StudentExamID, score, submittedAt)
	if err != nil {
		return nil, fmt.Errorf("error al insertar el resultado: %w", err)
	}

	entity := studentExamEntity.StudentExamResult{
		CorrectAnswers: correct,
		TotalQuestions: total,
		Score:          score,
	}

	return &entity, nil
}
