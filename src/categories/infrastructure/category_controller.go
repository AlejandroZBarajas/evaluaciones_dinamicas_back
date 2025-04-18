package categoryInfrastructure

import (
	"encoding/json"
	"net/http"

	categoryApplication "evaluaciones/src/categories/application"
	categoryEntity "evaluaciones/src/categories/domain/entity"
)

type CategoryController struct {
	CreateUseCase       *categoryApplication.CreateCategory
	GetByIDUseCase      *categoryApplication.GetCategoryByID
	GetByTeacherUseCase *categoryApplication.GetAllCategoriesByTeacherID
	UpdateUseCase       *categoryApplication.UpdateCategory
	DeleteUseCase       *categoryApplication.DeleteCategory
}

func NewCategoryController(
	create *categoryApplication.CreateCategory,
	getByID *categoryApplication.GetCategoryByID,
	getByTeacher *categoryApplication.GetAllCategoriesByTeacherID,
	update *categoryApplication.UpdateCategory,
	delete *categoryApplication.DeleteCategory,
) *CategoryController {
	return &CategoryController{
		CreateUseCase:       create,
		GetByIDUseCase:      getByID,
		GetByTeacherUseCase: getByTeacher,
		UpdateUseCase:       update,
		DeleteUseCase:       delete,
	}
}

func (cc *CategoryController) HandleCreateCategory(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name      string `json:"name"`
		TeacherID int32  `json:"teacher_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := cc.CreateUseCase.Run(body.Name, body.TeacherID)
	if err != nil {
		http.Error(w, "Error creating category", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (cc *CategoryController) HandleGetCategoryByID(w http.ResponseWriter, r *http.Request) {
	var body struct {
		ID int32 `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	category, err := cc.GetByIDUseCase.Run(body.ID)
	if err != nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(category)
}

func (cc *CategoryController) HandleGetCategoriesByTeacherID(w http.ResponseWriter, r *http.Request) {
	var body struct {
		TeacherID int32 `json:"teacher_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	categories, err := cc.GetByTeacherUseCase.Run(body.TeacherID)
	if err != nil {
		http.Error(w, "Error retrieving categories", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(categories)
}

// PUT /categories/update
func (cc *CategoryController) HandleUpdateCategory(w http.ResponseWriter, r *http.Request) {
	var body struct {
		ID        int32  `json:"id"`
		Name      string `json:"name"`
		TeacherID int32  `json:"teacher_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedCategory := &categoryEntity.CategoryEntity{
		Name:      body.Name,
		TeacherID: body.TeacherID,
	}

	err := cc.UpdateUseCase.Run(body.ID, updatedCategory)
	if err != nil {
		http.Error(w, "Error updating category", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (cc *CategoryController) HandleDeleteCategory(w http.ResponseWriter, r *http.Request) {
	var body struct {
		ID int32 `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := cc.DeleteUseCase.Run(body.ID)
	if err != nil {
		http.Error(w, "Error deleting category", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
