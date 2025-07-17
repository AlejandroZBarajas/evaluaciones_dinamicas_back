package resultsInfrastructure

import (
	"database/sql"
	"encoding/json"
	resultsEntity "evaluaciones/src/results/domain/entity"
)

type ResultsRepository struct {
	db *sql.DB
}

func NewResultsRepository(db *sql.DB) *ResultsRepository {
	return &ResultsRepository{db: db}
}

func (r *ResultsRepository) CreateResult(result *resultsEntity.ResultsEntity) error {
	query := `INSERT INTO results (student_exam_id, result, created_at) VALUES ($1, $2, $3) RETURNING id`

	resultJSON, err := json.Marshal(result.Result)
	if err != nil {
		return err
	}

	return r.db.QueryRow(query, result.StundetExamID, resultJSON, result.CreatedAt).Scan(&result.ID)
}

func (r *ResultsRepository) GetAllResultsByExam(examID int32) ([]*resultsEntity.ResultsEntity, error) {
	query := `
		SELECT r.id, r.student_exam_id, r.result, r.created_at
		FROM results r
		JOIN student_exams se ON r.student_exam_id = se.id
		WHERE se.exam_id = $1
	`

	rows, err := r.db.Query(query, examID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*resultsEntity.ResultsEntity

	for rows.Next() {
		var res resultsEntity.ResultsEntity
		var resultJSON []byte
		if err := rows.Scan(&res.ID, &res.StundetExamID, &resultJSON, &res.CreatedAt); err != nil {
			return nil, err
		}

		json.Unmarshal(resultJSON, &res.Result)
		results = append(results, &res)
	}

	return results, nil
}

func (r *ResultsRepository) DeleteAllResultsByExam(examID int32) error {
	query := `
		DELETE FROM results
		WHERE student_exam_id IN (
			SELECT id FROM student_exams WHERE exam_id = $1
		)
	`

	_, err := r.db.Exec(query, examID)
	return err
}
