package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/alhaos/quick-menu-api/internal/model"
	"github.com/alhaos/quick-menu-api/internal/utils"
)

func (r *Repository) CreateCategory(clientID string, category *model.Category) error {

	ctx, cancel := context.WithTimeout(context.Background(), utils.Timeout())
	defer cancel()

	tx, err := r.db.BeginEx(ctx, nil)
	if err != nil {
		return err
	}

	err = setSessionClientId(clientID, tx, ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	const query = `INSERT INTO categories (name, description, image) values ($1, $2, $3) RETURNING id;`

	err = r.db.QueryRowEx(
		ctx,
		query,
		nil,
		category.Name,
		category.Description,
		category.Image,
	).Scan(&category.ID)

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.CommitEx(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *Repository) GetCategoryByID(clientID string, id string) (*model.Category, error) {

	ctx, cancel := context.WithTimeout(context.Background(), utils.Timeout())
	defer cancel()

	tx, err := r.db.BeginEx(ctx, nil)
	if err != nil {
		return nil, err
	}

	err = setSessionClientId(clientID, tx, ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	var category model.Category

	query := `SELECT id, name, description FROM categories WHERE id = $1;`

	row := tx.QueryRowEx(ctx, query, nil, id)

	err = row.Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.CommitEx(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &category, nil
}

func (r *Repository) UpdateCategory(clientID string, category *model.Category) error {

	ctx, cancel := context.WithTimeout(context.Background(), utils.Timeout())
	defer cancel()

	tx, err := r.db.BeginEx(ctx, nil)
	if err != nil {
		return err
	}

	err = setSessionClientId(clientID, tx, ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	query := "UPDATE categories SET name = $1, description = $2, image = $3 WHERE id = $4"

	res, err := r.db.ExecEx(
		ctx,
		query,
		nil,
		category.Name,
		category.Description,
		category.Image,
		category.ID,
	)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return errors.New("no such category found")
	}

	err = tx.CommitEx(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *Repository) DeleteCategoryByID(clientID string, id string) error {

	ctx, cancel := context.WithTimeout(context.Background(), utils.Timeout())
	defer cancel()

	tx, err := r.db.BeginEx(ctx, nil)
	if err != nil {
		return err
	}

	err = setSessionClientId(clientID, tx, ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	query := "DELETE FROM categories WHERE id = $1;"
	ex, err := tx.ExecEx(ctx, query, nil, id)
	if err != nil {
		return err
	}

	if ex.RowsAffected() == 0 {
		return fmt.Errorf("category with id: [%s] does not exist", id)
	}

	err = tx.CommitEx(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *Repository) ListAllCategories(clientID string) ([]model.Category, error) {

	ctx, cancel := context.WithTimeout(context.Background(), utils.Timeout())
	defer cancel()

	tx, err := r.db.BeginEx(ctx, nil)
	if err != nil {
		return nil, err
	}

	err = setSessionClientId(clientID, tx, ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	var categories []model.Category

	query := `SELECT id, name, description, image FROM categories WHERE user_id=$1;`

	rows, err := r.db.QueryEx(ctx, query, nil, clientID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var category model.Category
		err = rows.Scan(&category.ID, &category.Name, &category.Description, &category.Image)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}
