package resultsInfrastructure

import (
	"encoding/json"
	resultsApplication "evaluaciones/src/results/application"
	resulstEntity "evaluaciones/src/results/domain/entity"
	"net/http"
	"strconv"
)

type ResultsController struct {
	createResult        *resultsApplication.CreateResult
	getAllResultsByExam *resultsApplication.GetAllResultsByExam
	deleteAllByExam     *resultsApplication.DeleteAllResultsByExam
}

func NewResultsController(
	create *resultsApplication.CreateResult,
	getAll *resultsApplication.GetAllResultsByExam,
	deleteAll *resultsApplication.DeleteAllResultsByExam,
) *ResultsController {
	return &ResultsController{
		createResult:        create,
		getAllResultsByExam: getAll,
		deleteAllByExam:     deleteAll,
	}
}

func (c *ResultsController) HandleCreateResult(w http.ResponseWriter, r *http.Request) {
	var result resulstEntity.ResultsEntity
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.createResult.Run(&result)
	if err != nil {
		http.Error(w, "Error al crear el resultado: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (c *ResultsController) HandleGetAllResultsByExam(w http.ResponseWriter, r *http.Request) {
	var body struct {
		ExamID int32 `json:"exam_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	results, err := c.getAllResultsByExam.Run(body.ExamID)
	if err != nil {
		http.Error(w, "Error al obtener los resultados: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func (c *ResultsController) HandleDeleteAllResultsByExam(w http.ResponseWriter, r *http.Request) {
	examIDStr := r.URL.Query().Get("exam_id")
	examID, err := strconv.Atoi(examIDStr)
	if err != nil {
		http.Error(w, "ID de examen inválido", http.StatusBadRequest)
		return
	}

	err = c.deleteAllByExam.Run(int32(examID))
	if err != nil {
		http.Error(w, "Error al eliminar los resultados: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
