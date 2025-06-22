package repo

import (
	"github.com/alhaos/quick-menu-api/internal/database"
	"github.com/alhaos/quick-menu-api/internal/model"

	"testing"
)

func TestRepository_AddItem(t *testing.T) {

	data := []struct {
		clientID string
	}{{
		clientID: "e1dcc9d5-24ae-403b-aebb-813075793148",
	}}

	dc := database.Config{
		Host:     "alhaos.online",
		Port:     5432,
		User:     "quickmenu",
		Pass:     "kuk@Zumba!",
		Database: "quickmenu",
	}

	db, err := database.New(dc)
	if err != nil {
		t.Error(err)
	}

	r := New(db)

	for _, datum := range data {
		id, err := r.NewItem(datum.clientID)
		t.Log("id:", id)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestRepository_GetItemById(t *testing.T) {

	data := []struct {
		id string
	}{
		{id: "b789d532-7d49-4480-8765-5902eb860800"},
	}

	dc := database.Config{
		Host:     "alhaos.online",
		Port:     5432,
		User:     "quickmenu",
		Pass:     "kuk@Zumba!",
		Database: "quickmenu",
	}

	db, err := database.New(dc)
	if err != nil {
		t.Error(err)
	}

	r := New(db)

	for _, datum := range data {
		_, err = r.GetItemById(datum.id)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestRepository_UpdateItem(t *testing.T) {

	data := []struct {
		item *model.Item
	}{
		{&model.Item{
			ID:            "b789d532-7d49-4480-8765-5902eb860800",
			Name:          "Шашлык",
			Description:   "Очень вкусный шашлык",
			ImageFilename: "",
			IsActive:      true,
		}},
	}

	dc := database.Config{
		Host:     "alhaos.online",
		Port:     5432,
		User:     "quickmenu",
		Pass:     "kuk@Zumba!",
		Database: "quickmenu",
	}
	db, err := database.New(dc)
	if err != nil {
		t.Error(err)
	}
	r := New(db)
	for _, datum := range data {
		err = r.UpdateItem(datum.item)
		if err != nil {
			t.Error(err)
		}
	}
}
