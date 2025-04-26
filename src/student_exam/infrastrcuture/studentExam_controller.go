package studentExamInfrastructure

import (
	"encoding/json"
	studentExamApplication "evaluaciones/src/student_exam/application"
	studentExamEntity "evaluaciones/src/student_exam/domain/entity"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type StudentExamController struct {
	createStudentExam   *studentExamApplication.CreateStudentExam
	getAllByExamID      *studentExamApplication.GetAllByExamID
	getStudentExamByID  *studentExamApplication.GetStudentExamByID
	deleteStudentExam   *studentExamApplication.DeleteStudentExam
	evaluateStudentExam *studentExamApplication.EvaluateStudentExam
	generateRandomExam  *studentExamApplication.GenerateRandomExam
}

func NewStudentExamController(
	create *studentExamApplication.CreateStudentExam,
	getAll *studentExamApplication.GetAllByExamID,
	getByID *studentExamApplication.GetStudentExamByID,
	deleteExam *studentExamApplication.DeleteStudentExam,
	evaluateExam *studentExamApplication.EvaluateStudentExam,
	generate *studentExamApplication.GenerateRandomExam,
) *StudentExamController {
	return &StudentExamController{
		createStudentExam:   create,
		getAllByExamID:      getAll,
		getStudentExamByID:  getByID,
		deleteStudentExam:   deleteExam,
		evaluateStudentExam: evaluateExam,
		generateRandomExam:  generate,
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
func (c *StudentExamController) HandleExamEvaluation(w http.ResponseWriter, r *http.Request) {
	var submission studentExamEntity.ExamSubmissionInput

	// Decodificar el cuerpo de la solicitud JSON
	if err := json.NewDecoder(r.Body).Decode(&submission); err != nil {
		http.Error(w, fmt.Sprintf("Error al leer el cuerpo de la solicitud: %v", err), http.StatusBadRequest)
		return
	}

	// Llamar al caso de uso para evaluar las respuestas del examen
	result, err := c.evaluateStudentExam.Run(&submission)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al evaluar el examen: %v", err), http.StatusInternalServerError)
		return
	}

	// Crear una respuesta JSON con el resultado
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Retornar la calificación obtenida en formato JSON
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, fmt.Sprintf("Error al escribir la respuesta: %v", err), http.StatusInternalServerError)
	}
}
func (c *StudentExamController) HandleGenerateRandomExam(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var input studentExamEntity.RandomExamInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Error al decodificar el cuerpo de la solicitud: "+err.Error(), http.StatusBadRequest)
		return
	}

	questions, err := c.generateRandomExam.Run(&input)
	if err != nil {
		http.Error(w, "Error al generar preguntas aleatorias: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}
