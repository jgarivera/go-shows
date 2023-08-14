package tickets

import (
	"database/sql"
)

type Repository interface {
	GetTicketById(id uint) (*Ticket, error)
	GetTickets() ([]*Ticket, error)
	CreateTicket(t *Ticket) (int64, error)
	UpdateTicket(t *Ticket) (int64, error)
	DeleteTicket(id uint) (int64, error)
	DoesTicketExist(id uint) (bool, error)
}

type SqlRepository struct {
	Database *sql.DB
}

func NewSqlRepository(Database *sql.DB) *SqlRepository {
	return &SqlRepository{
		Database,
	}
}

func (r *SqlRepository) GetTicketById(id uint) (*Ticket, error) {
	t := new(Ticket)

	row := r.Database.QueryRow("SELECT * FROM tickets WHERE id = ?", id)

	err := row.Scan(&t.ID, &t.Name, &t.Description, &t.CreatedAt, &t.UpdatedAt)

	return t, err
}

func (r *SqlRepository) GetTickets() ([]*Ticket, error) {
	rows, err := r.Database.Query("SELECT * FROM tickets")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tickets := []*Ticket{}

	for rows.Next() {
		t := new(Ticket)

		if err := rows.Scan(&t.ID, &t.Name, &t.Price, &t.Description, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}

		tickets = append(tickets, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tickets, nil
}

func (r *SqlRepository) CreateTicket(t *Ticket) (int64, error) {
	res, err := r.Database.Exec("INSERT INTO tickets(name, price, description) VALUES (?, ?, ?)", t.Name, t.Price, t.Description)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (r *SqlRepository) UpdateTicket(t *Ticket) (int64, error) {
	res, err := r.Database.Exec("UPDATE tickets SET name = ?, price = ?, description = ?, WHERE id = ?", t.Name, t.Price, t.Description, t.ID)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (r *SqlRepository) DeleteTicket(id uint) (int64, error) {
	res, err := r.Database.Exec("DELETE FROM tickets WHERE id = ?", id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (r *SqlRepository) DoesTicketExist(id uint) (bool, error) {
	row := r.Database.QueryRow("SELECT * FROM tickets WHERE id = ?", id)

	t := new(Ticket)

	if err := row.Scan(&t.ID); err != nil {
		return false, err
	} else {
		return true, nil
	}
}
