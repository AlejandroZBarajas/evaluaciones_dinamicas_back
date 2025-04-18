package categoryInfrastructure

import (
	"database/sql"

	categoryDomain "evaluaciones/src/categories/domain"
	categoryEntity "evaluaciones/src/categories/domain/entity"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) categoryDomain.CategoryInterface {
	return &CategoryRepository{db: db}
}

func (repo *CategoryRepository) CreateCategory(category *categoryEntity.CategoryEntity) error {
	query := "INSERT INTO categories (name, teacher_id) VALUES ($1, $2)"
	_, err := repo.db.Exec(query, category.Name, category.TeacherID)
	return err
}

func (repo *CategoryRepository) GetCategoryByID(id int32) (*categoryEntity.CategoryEntity, error) {
	query := "SELECT id, name, teacher_id FROM categories WHERE id = $1"
	row := repo.db.QueryRow(query, id)

	var category categoryEntity.CategoryEntity
	err := row.Scan(&category.ID, &category.Name, &category.TeacherID)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (repo *CategoryRepository) GetAllCategoriesByTeacherID(teacherID int32) ([]*categoryEntity.CategoryEntity, error) {
	query := "SELECT id, name, teacher_id FROM categories WHERE teacher_id = $1"
	rows, err := repo.db.Query(query, teacherID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*categoryEntity.CategoryEntity
	for rows.Next() {
		var category categoryEntity.CategoryEntity
		err := rows.Scan(&category.ID, &category.Name, &category.TeacherID)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}

	return categories, nil
}

func (repo *CategoryRepository) UpdateCategory(categoryID int32, category *categoryEntity.CategoryEntity) error {
	query := "UPDATE categories SET name = $1, teacher_id = $2 WHERE id = $3"
	_, err := repo.db.Exec(query, category.Name, category.TeacherID, categoryID)
	return err
}
func (repo *CategoryRepository) DeleteCategory(categoryID int32) error {
	query := "DELETE FROM categories WHERE id = $1"
	_, err := repo.db.Exec(query, categoryID)
	return err
}
