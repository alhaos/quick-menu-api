package repo

import (
	"context"
	"github.com/alhaos/quick-menu-api/internal/model"
	"github.com/alhaos/quick-menu-api/internal/utils"
	"time"
)

// NewItem ...
func (r *Repository) NewItem(clientID string) (*model.Item, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*utils.Timeout)
	defer cancel()

	const query = "INSERT INTO client_data.items (client_id) VALUES ($1) RETURNING id"

	var id string

	err := r.db.QueryRowEx(ctx, query, nil, clientID).Scan(&id)
	if err != nil {
		return nil, err
	}

	item := model.Item{
		ID: id,
	}

	return &item, nil
}

// GetItemById ...
func (r *Repository) GetItemById(id string) (model.Item, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*utils.Timeout)
	defer cancel()

	const query = `
	SELECT id,
	       coalesce(name, ''),
	       coalesce(description, ''),
	       coalesce(image_filename, ''),
           is_active
	  FROM client_data.items
	 WHERE id = $1`

	var item model.Item
	if err := r.db.QueryRowEx(ctx, query, nil, id).Scan(&item.ID, &item.Name, &item.Description, &item.ImageFilename, &item.IsActive); err != nil {
		return model.Item{}, err
	}
	return item, nil
}

// RemoveItemById ...
func (r *Repository) RemoveItemById(id string) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*utils.Timeout)
	defer cancel()

	const query = `DELETE FROM client_data.items WHERE id = $1`

	_, err := r.db.ExecEx(ctx, query, nil, id)
	if err != nil {
		return err
	}

	return nil

}

// UpdateItem ...
func (r *Repository) UpdateItem(item *model.Item) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*utils.Timeout)
	defer cancel()

	const query = `
UPDATE client_data.items
   SET name = $1,
       description = $2,
       image_filename = $3,
       is_active = $4 
 WHERE id = $5`

	_, err := r.db.ExecEx(ctx, query, nil, item.Name, item.Description, item.ImageFilename, item.IsActive, item.ID)
	if err != nil {
		return err
	}
	return nil
}
