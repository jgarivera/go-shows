package shows

import (
	"database/sql"
)

type Repository interface {
	GetShowById(id uint) (*Show, error)
	GetShows() ([]*Show, error)
	CreateShow(s *Show) (int64, error)
	UpdateShow(s *Show) (int64, error)
	DeleteShow(id uint) (int64, error)
	DoesShowExist(id uint) (bool, error)
}

type SqlRepository struct {
	Database *sql.DB
}

func NewSqlRepository(Database *sql.DB) *SqlRepository {
	return &SqlRepository{
		Database,
	}
}

func (r *SqlRepository) GetShowById(id uint) (*Show, error) {
	s := new(Show)

	row := r.Database.QueryRow("SELECT * FROM shows WHERE id = ?", id)

	err := row.Scan(&s.ID, &s.Name, &s.Description, &s.CreatedAt, &s.UpdatedAt)

	return s, err
}

func (r *SqlRepository) GetShows() ([]*Show, error) {
	rows, err := r.Database.Query("SELECT * FROM shows")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	shows := []*Show{}

	for rows.Next() {
		s := new(Show)

		if err := rows.Scan(&s.ID, &s.Name, &s.Price, &s.Description, &s.CreatedAt, &s.UpdatedAt); err != nil {
			return nil, err
		}

		shows = append(shows, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return shows, nil
}

func (r *SqlRepository) CreateShow(s *Show) (int64, error) {
	res, err := r.Database.Exec("INSERT INTO shows(name, price, description) VALUES (?, ?, ?)", s.Name, s.Price, s.Description)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (r *SqlRepository) UpdateShow(s *Show) (int64, error) {
	res, err := r.Database.Exec("UPDATE shows SET name = ?, price = ?, description = ?, WHERE id = ?", s.Name, s.Price, s.Description, s.ID)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (r *SqlRepository) DeleteShow(id uint) (int64, error) {
	res, err := r.Database.Exec("DELETE FROM shows WHERE id = ?", id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (r *SqlRepository) DoesShowExist(id uint) (bool, error) {
	row := r.Database.QueryRow("SELECT * FROM shows WHERE id = ?", id)

	s := new(Show)

	if err := row.Scan(&s.ID); err != nil {
		return false, err
	} else {
		return true, nil
	}
}
