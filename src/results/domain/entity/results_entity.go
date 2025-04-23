package resulstEntity

import "time"

type ResultsEntity struct {
	ID            int32     `jaon:"id"`
	StundetExamID int32     `json:"student_exam_id"`
	Result        string    `json:"result"`
	CreatedAt     time.Time `json:"created_at"`
}

func CreateResult(id int32, studentExamID int32, result string, cratedat time.Time) *ResultsEntity {
	return &ResultsEntity{ID: id, StundetExamID: studentExamID, Result: result, CreatedAt: cratedat}
}

/*
[
	{
		"id": 3,
		"exam_id": 2,
		"student_id": 1,
		"total_questions": 10,
		"correct_answers": 8,
		"responses": {
			"1": "A",
			"2": "B",
			"3": "C",
			"4": "D"
		},
		"submitted_at": "0001-01-01T00:00:00Z"
	},
	{
		"id": 4,
		"exam_id": 2,
		"student_id": 1,
		"total_questions": 10,
		"correct_answers": 8,
		"responses": {
			"1": "A",
			"2": "B",
			"3": "C",
			"4": "D"
		},
		"submitted_at": "0001-01-01T00:00:00Z"
	},
	{
		"id": 5,
		"exam_id": 2,
		"student_id": 1,
		"total_questions": 10,
		"correct_answers": 8,
		"responses": {
			"1": "A",
			"2": "B",
			"3": "C",
			"4": "D"
		},
		"submitted_at": "0001-01-01T00:00:00Z"
	},
	{
		"id": 6,
		"exam_id": 2,
		"student_id": 1,
		"total_questions": 10,
		"correct_answers": 8,
		"responses": {
			"1": "A",
			"2": "B",
			"3": "C",
			"4": "D"
		},
		"submitted_at": "2025-04-22T18:30:00Z"
	}
]

*/
