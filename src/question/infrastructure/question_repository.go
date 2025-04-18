package questionInfrastructure

import (
	"database/sql"
	"encoding/json"
	questionDomain "evaluaciones/src/question/domain"
	questionEntity "evaluaciones/src/question/domain/entity"
)

type QuestionRepository struct {
	db *sql.DB
}

func NewQuestionRepository(db *sql.DB) questionDomain.QuestionInterface {
	return &QuestionRepository{db: db}
}

func (qr *QuestionRepository) CreateQuestion(q *questionEntity.QuestionEntity) error {
	data, err := json.Marshal(q.QuestionData)
	if err != nil {
		return err
	}

	query := `INSERT INTO questions (question_data, category_id, exam_id) VALUES ($1, $2, $3)`
	_, err = qr.db.Exec(query, data, q.CategoryID, q.ExamID)
	return err
}

func (qr *QuestionRepository) GetQuestionByID(id int32) (*questionEntity.QuestionEntity, error) {
	query := `SELECT id, question_data, category_id, exam_id FROM questions WHERE id = $1`
	row := qr.db.QueryRow(query, id)

	var q questionEntity.QuestionEntity
	var data []byte

	err := row.Scan(&q.ID, &data, &q.CategoryID, &q.ExamID)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &q.QuestionData)
	if err != nil {
		return nil, err
	}

	return &q, nil
}

func (qr *QuestionRepository) UpdateQuestion(id int32, q *questionEntity.QuestionEntity) error {
	data, err := json.Marshal(q.QuestionData)
	if err != nil {
		return err
	}

	query := `UPDATE questions SET question_data = $1, category_id = $2, exam_id = $3 WHERE id = $4`
	_, err = qr.db.Exec(query, data, q.CategoryID, q.ExamID, id)
	return err
}

func (qr *QuestionRepository) DeleteQuestion(id int32) error {
	query := `DELETE FROM questions WHERE id = $1`
	_, err := qr.db.Exec(query, id)
	return err
}

func (qr *QuestionRepository) GetAllQuestionsByExam(examID int32) ([]*questionEntity.QuestionEntity, error) {
	query := `SELECT id, question_data, category_id, exam_id FROM questions WHERE exam_id = $1`
	rows, err := qr.db.Query(query, examID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []*questionEntity.QuestionEntity
	for rows.Next() {
		var q questionEntity.QuestionEntity
		var data []byte
		if err := rows.Scan(&q.ID, &data, &q.CategoryID, &q.ExamID); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(data, &q.QuestionData); err != nil {
			return nil, err
		}
		questions = append(questions, &q)
	}
	return questions, nil
}

func (qr *QuestionRepository) GetAllQuestionsByCategory(categoryID int32) ([]*questionEntity.QuestionEntity, error) {
	query := `SELECT id, question_data, category_id, exam_id FROM questions WHERE category_id = $1`
	rows, err := qr.db.Query(query, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []*questionEntity.QuestionEntity
	for rows.Next() {
		var q questionEntity.QuestionEntity
		var data []byte
		if err := rows.Scan(&q.ID, &data, &q.CategoryID, &q.ExamID); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(data, &q.QuestionData); err != nil {
			return nil, err
		}
		questions = append(questions, &q)
	}
	return questions, nil
}
