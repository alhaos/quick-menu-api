package repository

import (
	"context"
	"errors"
	"github.com/alhaos/quick-menu-api/internal/model"
	"github.com/alhaos/quick-menu-api/internal/utils"
	"github.com/jackc/pgx"
)

// CreateItem ...
func (r *Repository) CreateItem(clientID string, item *model.Item) error {

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

	query := `
INSERT INTO public.items(
	name, description, image, is_active)
	VALUES ($1, $2, $3, $4)
	RETURNING id`

	err = r.db.QueryRowEx(
		ctx,
		query,
		nil,
		item.Name,
		item.Description,
		item.Image,
		item.IsActive,
	).Scan(&item.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// GetItemById ...
func (r *Repository) GetItemById(clientID string, id string) (*model.Item, error) {

	ctx, cancel := context.WithTimeout(context.Background(), utils.Timeout())
	defer cancel()

	tx, err := r.db.BeginEx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = setSessionClientId(clientID, tx, ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	const query = `
	SELECT id,
	       coalesce(name, ''),
	       coalesce(description, ''),
	       coalesce(image, ''),
	       coalesce(price, ''),
           is_active
	  FROM items
	 WHERE id = $1`

	var item model.Item
	err = tx.QueryRowEx(ctx, query, nil, id).Scan(&item.ID, &item.Name, &item.Description, &item.Image, &item.Price, &item.IsActive)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.New("no item found")
	}
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &item, nil
}

// DeleteItemById ...
func (r *Repository) DeleteItemById(clientID string, id string) error {

	ctx, cancel := context.WithTimeout(context.Background(), utils.Timeout())
	defer cancel()

	tx, err := r.db.BeginEx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = setSessionClientId(clientID, tx, ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	const query = `DELETE FROM items WHERE id = $1`

	res, err := tx.ExecEx(ctx, query, nil, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	if res.RowsAffected() == 0 {
		return errors.New("no item found")
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// UpdateItem ...
func (r *Repository) UpdateItem(clientID string, item *model.Item) error {

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

	const query = `
UPDATE items
   SET name = $1,
       description = $2,
       image = $3,
       price = $4,
       is_active = $5 
 WHERE id = $6`

	res, err := tx.ExecEx(ctx, query, nil, item.Name, item.Description, item.Image, item.Price, item.IsActive, item.ID)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		tx.Rollback()
		return errors.New("no item found")
	}

	err = tx.Commit()

	return nil
}

// ListItems ...
func (r *Repository) ListItems(clientID string) ([]model.Item, error) {

	ctx, cancel := context.WithTimeout(context.Background(), utils.Timeout())
	defer cancel()

	tx, err := r.db.BeginEx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = setSessionClientId(clientID, tx, ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	const query = `
	SELECT id,
	       coalesce(name, ''),
	       coalesce(description, ''),
	       coalesce(image, ''),
	       coalesce(price, ''),
           is_active
	  FROM items`

	var items []model.Item
	var item model.Item

	rows, err := tx.QueryEx(ctx, query, nil)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&item.ID, &item.Name, &item.Description, &item.Image, &item.Price, &item.IsActive)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return items, nil
}

// setSessionClientId
// Set up app.current_user_id variable in postgresql session
func setSessionClientId(clientID string, tx *pgx.Tx, ctx context.Context) error {

	setClientIDCommand := "set app.current_user_id to '" + clientID + "'"

	_, err := tx.ExecEx(ctx, setClientIDCommand, nil)
	if err != nil {
		return err
	}
	return err
}
