package examInfrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"

	//stronv"

	examApplication "evaluaciones/src/exam/application"
	examEntity "evaluaciones/src/exam/domain/entity"
)

type ExamController struct {
	CreateUseCase           *examApplication.CreateExam
	GetAllByTeacher         *examApplication.GetAllExamsByTeacherID
	GetByID                 *examApplication.GetExamByID
	UpdateUseCase           *examApplication.UpdateExam
	DeleteUseCase           *examApplication.DeleteExam
	GetbyTeacherAndCategory *examApplication.GetbyTeacherAndCategory
}

func NewExamController(
	create *examApplication.CreateExam,
	getAllByTeacher *examApplication.GetAllExamsByTeacherID,
	getByID *examApplication.GetExamByID,
	update *examApplication.UpdateExam,
	delete *examApplication.DeleteExam,
	teacherAndCategory *examApplication.GetbyTeacherAndCategory,
) *ExamController {
	return &ExamController{
		CreateUseCase:           create,
		GetAllByTeacher:         getAllByTeacher,
		GetByID:                 getByID,
		UpdateUseCase:           update,
		DeleteUseCase:           delete,
		GetbyTeacherAndCategory: teacherAndCategory,
	}
}

func (c *ExamController) HandleCreateExam(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name           string `json:"name"`
		TotalQuestions int32  `json:"total_questions"`
		TeacherID      int32  `json:"teacher_id"`
		CategoryId     int32  `json:"category_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := c.CreateUseCase.Run(body.Name, body.TotalQuestions, body.TeacherID, body.CategoryId)
	if err != nil {
		http.Error(w, "Error creating exam", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *ExamController) HandleGetAllByTeacher(w http.ResponseWriter, r *http.Request) {
	var body struct {
		TeacherID int32 `json:"teacher_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	exams, err := c.GetAllByTeacher.Run(body.TeacherID)
	if err != nil {
		http.Error(w, "Error retrieving exams", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(exams)
}

func (c *ExamController) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	var body struct {
		ID int32 `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	exam, err := c.GetByID.Run(body.ID)
	if err != nil {
		http.Error(w, "Exam not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(exam)
}

func (c *ExamController) HandleUpdateExam(w http.ResponseWriter, r *http.Request) {
	var body examEntity.ExamEntity

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updated, err := c.UpdateUseCase.Run(body.ID, &body)
	if err != nil {
		http.Error(w, "Error updating exam", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updated)
}

func (c *ExamController) HandleDeleteExam(w http.ResponseWriter, r *http.Request) {
	var body struct {
		ID int32 `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := c.DeleteUseCase.Run(body.ID); err != nil {
		http.Error(w, "Error deleting exam", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *ExamController) HandleTeacherAndCategory(w http.ResponseWriter, r *http.Request) {
	var req struct {
		TeacherID  int32 `json:"teacher_id"`
		CategoryID int32 `json:"category_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "❌ Error al decodificar el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	exams, err := c.GetbyTeacherAndCategory.Run(req.TeacherID, req.CategoryID)
	if err != nil {
		http.Error(w, fmt.Sprintf("❌ Error al obtener exámenes: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(exams)
}
