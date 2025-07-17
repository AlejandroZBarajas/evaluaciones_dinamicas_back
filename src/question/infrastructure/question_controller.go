package questionInfrastructure

import (
	"encoding/json"
	questionApplication "evaluaciones/src/question/application"
	questionEntity "evaluaciones/src/question/domain/entity"
	"net/http"
)

type QuestionController struct {
	CreateUseCase           *questionApplication.CreateQuestion
	GetByIDUseCase          *questionApplication.GetQuestionByID
	UpdateUseCase           *questionApplication.UpdateQuestion
	DeleteUseCase           *questionApplication.DeleteQuestion
	GetAllByExamUseCase     *questionApplication.GetAllQuestionsByExam
	GetAllByCategoryUseCase *questionApplication.GetAllQuestionsByCategory
}

func NewQuestionController(
	create *questionApplication.CreateQuestion,
	getByID *questionApplication.GetQuestionByID,
	update *questionApplication.UpdateQuestion,
	deleteQ *questionApplication.DeleteQuestion,
	getByExam *questionApplication.GetAllQuestionsByExam,
	getByCategory *questionApplication.GetAllQuestionsByCategory,
) *QuestionController {
	return &QuestionController{
		CreateUseCase:           create,
		GetByIDUseCase:          getByID,
		UpdateUseCase:           update,
		DeleteUseCase:           deleteQ,
		GetAllByExamUseCase:     getByExam,
		GetAllByCategoryUseCase: getByCategory,
	}
}

func (qc *QuestionController) HandleCreateQuestion(w http.ResponseWriter, r *http.Request) {
	var q questionEntity.QuestionEntity
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}
	err := qc.CreateUseCase.Run(q.QuestionData, q.CategoryID, q.ExamID)
	if err != nil {
		http.Error(w, "Error creating question", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (qc *QuestionController) HandleGetQuestionByID(w http.ResponseWriter, r *http.Request) {
	var body struct {
		ID int32 `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}
	question, err := qc.GetByIDUseCase.Run(body.ID)
	if err != nil {
		http.Error(w, "Question not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(question)
}

func (qc *QuestionController) HandleUpdateQuestion(w http.ResponseWriter, r *http.Request) {
	var q questionEntity.QuestionEntity
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}
	err := qc.UpdateUseCase.Run(q.ID, q.QuestionData, q.CategoryID, q.ExamID)
	if err != nil {
		http.Error(w, "Error updating question", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (qc *QuestionController) HandleDeleteQuestion(w http.ResponseWriter, r *http.Request) {
	var body struct {
		ID int32 `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}
	err := qc.DeleteUseCase.Run(body.ID)
	if err != nil {
		http.Error(w, "Error deleting question", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (qc *QuestionController) HandleGetAllByExam(w http.ResponseWriter, r *http.Request) {
	var body struct {
		ExamID int32 `json:"exam_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}
	questions, err := qc.GetAllByExamUseCase.Run(body.ExamID)
	if err != nil {
		http.Error(w, "Error getting questions", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(questions)
}

func (qc *QuestionController) HandleGetAllByCategory(w http.ResponseWriter, r *http.Request) {
	var body struct {
		CategoryID int32 `json:"category_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}
	questions, err := qc.GetAllByCategoryUseCase.Run(body.CategoryID)
	if err != nil {
		http.Error(w, "Error getting questions", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(questions)
}
