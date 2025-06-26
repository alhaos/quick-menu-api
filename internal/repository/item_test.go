package repository

import (
	"github.com/alhaos/quick-menu-api/internal/database"
	"github.com/alhaos/quick-menu-api/internal/model"
	"testing"
)

const client_id = "35aeac17-fbb1-4b3a-8f13-e4b48d14119e"

func TestRepository_CreateItem(t *testing.T) {

	dbConf := database.Config{
		Host:     "alhaos.online",
		Port:     5432,
		User:     "qm_owner",
		Pass:     "kuk2Zumba!",
		Database: "quick_menu",
	}

	db, err := database.New(dbConf)
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	r := New(db)

	item := model.MockItem()

	err = r.CreateItem(client_id, &item)
	if err != nil {
		t.Error(err)
	}
}
