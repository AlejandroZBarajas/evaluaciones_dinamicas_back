package studentExamInfrastructure

import (
	"encoding/json"
	studentExamApplication "evaluaciones/src/student_exam/application"
	studentExamEntity "evaluaciones/src/student_exam/domain/entity"
	"net/http"
	"strconv"
	"strings"
)

type StudentExamController struct {
	createStudentExam  *studentExamApplication.CreateStudentExam
	getAllByExamID     *studentExamApplication.GetAllByExamID
	getStudentExamByID *studentExamApplication.GetStudentExamByID
	deleteStudentExam  *studentExamApplication.DeleteStudentExam
}

func NewStudentExamController(
	create *studentExamApplication.CreateStudentExam,
	getAll *studentExamApplication.GetAllByExamID,
	getByID *studentExamApplication.GetStudentExamByID,
	deleteExam *studentExamApplication.DeleteStudentExam,
) *StudentExamController {
	return &StudentExamController{
		createStudentExam:  create,
		getAllByExamID:     getAll,
		getStudentExamByID: getByID,
		deleteStudentExam:  deleteExam,
	}
}

func (c *StudentExamController) HandleCreateStudentExam(w http.ResponseWriter, r *http.Request) {
	var studentExam studentExamEntity.StudentExamEntity
	if err := json.NewDecoder(r.Body).Decode(&studentExam); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.createStudentExam.Run(&studentExam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(studentExam)
}

func (c *StudentExamController) HandleGetAllByExamID(w http.ResponseWriter, r *http.Request) {
	examIDStr := r.URL.Query().Get("exam_id")
	examID, err := strconv.Atoi(examIDStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	studentExams, err := c.getAllByExamID.Run(int32(examID))

	if err != nil {
		http.Error(w, "Error al obtener los exámenes del estudiante: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(studentExams)
}

func (c *StudentExamController) HandleGetStudentExamByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/student-exams/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	studentExam, err := c.getStudentExamByID.Run(int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(studentExam)
}

func (c *StudentExamController) HandleDeleteStudentExam(w http.ResponseWriter, r *http.Request) {

	idStr := strings.TrimPrefix(r.URL.Path, "/student-exams/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = c.deleteStudentExam.Run(int32(id))
	if err != nil {
		http.Error(w, "Error al eliminar el examen del estudiante: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Eliminado con éxito"})

	w.WriteHeader(http.StatusNoContent)
}
