package repository

import (
	"context"
	"fmt"
	"github.com/alhaos/quick-menu-api/internal/model"
	"github.com/alhaos/quick-menu-api/internal/utils"
	"time"
)

func (r *Repository) CreateCategory(clientID string, category *model.Category) error {

	timeout := time.Millisecond * utils.Timeout()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	const query = `INSERT INTO client_data.categories (client_id, name, description) values ($1, $2, $3) RETURNING id;`

	err := r.db.QueryRowEx(
		ctx,
		query,
		nil,
		clientID,
		category.Name,
		category.Description,
	).Scan(&category.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetCategoryByID(id string, clientID string) (*model.Category, error) {

	ctx, cancel := context.WithTimeout(context.Background(), utils.Timeout())
	defer cancel()

	var category model.Category

	query := `SELECT id, name, description FROM client_data.categories WHERE id = $1 and client_Id=$2;`

	row := r.db.QueryRowEx(ctx, query, nil, id, clientID)

	err := row.Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *Repository) UpdateCategory(clientID string, category *model.Category) error {

	ctx, cancel := context.WithTimeout(context.Background(), utils.Timeout())
	defer cancel()

	query := "UPDATE client_data.categories SET name = $1, description = $2 WHERE id = $3 and client_id = $4 RETURNING id, client_id, name, description;"

	err := r.db.QueryRowEx(
		ctx,
		query,
		nil,
		category.Name,
		category.Description,
		category.ID,
		clientID,
	).Scan(
		&category.ID,
		&category.Name,
		&category.Description,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteCategoryByID(clientID string, id string) error {

	ctx, cancel := context.WithTimeout(context.Background(), utils.Timeout())
	defer cancel()
	query := "DELETE FROM client_data.categories WHERE id = $1 and client_id = $2;"
	ex, err := r.db.ExecEx(ctx, query, nil, id, clientID)
	if err != nil {
		return err
	}

	if ex.RowsAffected() == 0 {
		return fmt.Errorf("category with id: [%s] does not exist", id)
	}

	return nil
}

func (r *Repository) ListAllCategories(clientID string) ([]model.Category, error) {

	ctx, cancel := context.WithTimeout(context.Background(), utils.Timeout())
	defer cancel()

	var categories []model.Category

	query := `SELECT id, client_id, name, description FROM client_data.categories WHERE client_Id=$1;`

	rows, err := r.db.QueryEx(ctx, query, nil, clientID)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var category model.Category
		err = rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}
