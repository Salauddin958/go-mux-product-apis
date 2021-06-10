package post

import (
	"context"
	"database/sql"

	models "github.com/Salauddin958/go-mux-product-apis/models"
	repo "github.com/Salauddin958/go-mux-product-apis/repository"
)

// NewSQLProductRepo returns implement of post repository interface
func NewSQLProductRepo(Conn *sql.DB) repo.ProductRepo {
	return &mysqlProductRepo{
		Conn: Conn,
	}
}

type mysqlProductRepo struct {
	Conn *sql.DB
}

func (m *mysqlProductRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Product, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Product, 0)
	for rows.Next() {
		data := new(models.Product)

		err := rows.Scan(
			&data.ID,
			&data.Name,
			&data.Price,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *mysqlProductRepo) Fetch(ctx context.Context, num int64) ([]*models.Product, error) {
	query := "Select id, name, price From products limit ?"

	return m.fetch(ctx, query, num)
}

func (m *mysqlProductRepo) GetByID(ctx context.Context, id int64) (*models.Product, error) {
	query := "Select id, name, price From products where id=?"

	rows, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	payload := &models.Product{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}

func (m *mysqlProductRepo) Create(ctx context.Context, b *models.Product) (int64, error) {
	query := "INSERT INTO products(name, price) VALUES(?, ?)"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, b.Name, b.Price)
	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (m *mysqlProductRepo) Update(ctx context.Context, b *models.Product) (*models.Product, error) {
	query := "UPDATE products SET name=?, price=? WHERE id=?"
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		b.Name,
		b.Price,
		b.ID,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return b, nil
}

func (m *mysqlProductRepo) Delete(ctx context.Context, id int64) (bool, error) {
	query := "Delete From products Where id=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
