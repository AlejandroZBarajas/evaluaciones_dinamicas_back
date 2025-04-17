package userInfrastructure

import (
	"database/sql"
	userEntity "evaluaciones/src/users/domain/entity"
	"fmt"
	"regexp"
	"strings"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) CreateUser(user *userEntity.UserEntity) error {
	parts := strings.Split(user.Email, "@")
	if len(parts) != 2 || parts[1] != "ids.upchiapas.edu.mx" {
		return fmt.Errorf("correo no válido: solo se aceptan cuentas institucionales")
	}

	if matched, _ := regexp.MatchString(`^\d{6}$`, parts[0]); matched {
		user.RoleID = 1
	} else {
		user.RoleID = 2
	}

	query := `INSERT INTO users (email, matricula, role_id) VALUES ($1, $2, $3) RETURNING id`
	err := repo.db.QueryRow(query, user.Email, user.Matricula, user.RoleID).Scan(&user.Id)
	if err != nil {
		return fmt.Errorf("error al insertar usuario: %w", err)
	}

	return nil
}

func (repo *UserRepository) GetById(id int32) (*userEntity.UserEntity, error) {
	query := "SELECT id, email, matricula, role_id FROM users WHERE id = $1"
	row := repo.db.QueryRow(query, id)
	var user userEntity.UserEntity
	err := row.Scan(&user.Id, &user.Email, &user.Matricula, &user.RoleID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) GetByEmail(email string) (*userEntity.UserEntity, error) {
	query := "SELECT id, email, matricula, role_id FROM users WHERE email = $1"
	row := repo.db.QueryRow(query, email)
	var user userEntity.UserEntity
	err := row.Scan(&user.Id, &user.Email, &user.Matricula, &user.RoleID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) GetByMatricula(matricula string) (*userEntity.UserEntity, error) {
	query := "SELECT id, email, matricula, role_id FROM users WHERE matricula = $1"
	row := repo.db.QueryRow(query, matricula)
	var user userEntity.UserEntity
	err := row.Scan(&user.Id, &user.Email, &user.Matricula, &user.RoleID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) GetAll() ([]*userEntity.UserEntity, error) {
	query := "SELECT id, email, matricula, role_id FROM users"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*userEntity.UserEntity
	for rows.Next() {
		var user userEntity.UserEntity
		err := rows.Scan(&user.Id, &user.Email, &user.Matricula, &user.RoleID)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (repo *UserRepository) GetByRole(role int32) ([]*userEntity.UserEntity, error) {
	query := "SELECT id, email, matricula, role_id FROM users WHERE role_id = $1"
	rows, err := repo.db.Query(query, role)
	fmt.Println("🔍 Buscando usuarios con role_id =", role)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*userEntity.UserEntity
	for rows.Next() {
		var user userEntity.UserEntity
		err := rows.Scan(&user.Id, &user.Email, &user.Matricula, &user.RoleID)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (repo *UserRepository) ExistsByEmail(email string) (bool, error) {
	query := "SELECT 1 FROM users WHERE email = $1"
	row := repo.db.QueryRow(query, email)
	var exists int
	err := row.Scan(&exists)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *UserRepository) ExistsByMatricula(matricula string) (bool, error) {
	query := "SELECT 1 FROM users WHERE matricula = $1"
	row := repo.db.QueryRow(query, matricula)
	var exists int
	err := row.Scan(&exists)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
