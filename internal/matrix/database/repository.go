package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/matrix/domain"
	"github.com/sumelms/microservice-course/pkg/errors"
)

// Repository Matrix
type Repository struct {
	*sqlx.DB
}

// Matrix get the matrix by given id
func (r *Repository) Matrix(id uuid.UUID) (domain.Matrix, error) {
	var m domain.Matrix
	query := `SELECT * FROM matrices WHERE deleted_at IS NULL AND uuid = $1`
	if err := r.Get(&m, query, id); err != nil {
		return domain.Matrix{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting course")
	}
	return m, nil
}

// Matrices get the list of matrices
func (r *Repository) Matrices() ([]domain.Matrix, error) {
	var mm []domain.Matrix
	query := `SELECT * FROM matrices WHERE deleted_at IS NULL`
	if err := r.Select(&mm, query); err != nil {
		return []domain.Matrix{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting matrices")
	}
	return mm, nil
}

// CreateMatrix create a new matrix
func (r *Repository) CreateMatrix(m *domain.Matrix) error {
	query := `INSERT INTO matrices (title, description, course_id) VALUES ($1, $2, $3) RETURNING *`
	if err := r.Get(m, query, m.Title, m.Description, m.CourseID); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error creating matrix")
	}
	return nil
}

// UpdateMatrix updates the given matrix
func (r *Repository) UpdateMatrix(m *domain.Matrix) error {
	query := `UPDATE matrices SET title = $1, description = $2, course_id = $3 WHERE uuid = $4 RETURNING *`
	if err := r.Get(m, query, m.Title, m.Description, m.CourseID, m.UUID); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error updating matrix")
	}
	return nil
}

// DeleteMatrix delete the given matrix by uuid
func (r *Repository) DeleteMatrix(id uuid.UUID) error {
	query := `UPDATE matrices SET deleted_at = $1 WHERE uuid = $2`
	if _, err := r.Exec(query, time.Now(), id); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error deleting matrix")
	}
	return nil
}
