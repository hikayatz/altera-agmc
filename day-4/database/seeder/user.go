package seeder

import (
	"log"
	"time"

	"github.com/hikayat13/alterra-agcm/day-2/submission/models"
	"gorm.io/gorm"
)

func userSeeder(db *gorm.DB) {
	now := time.Now()
	var users = []models.User{
		{
			Common:    models.Common{ID: 1, CreatedAt: now, UpdatedAt: now},
			Name:      "devoncthomas",
			Email:     "devoncthomas@superrito.com",
			Password:  "123456",
			CreatedBy: 2,
		},
		{
			Common:    models.Common{ID: 2, CreatedAt: now, UpdatedAt: now},
			Name:      "hikayat",
			Email:     "hikayat@gmail.com",
			Password:  "123456",
			CreatedBy: 1,
		},
		{
			Common:    models.Common{ID: 3, CreatedAt: now, UpdatedAt: now},
			Name:      "Reza",
			Email:     "Reza@gmail.com",
			Password:  "123456",
			CreatedBy: 1,
		},
		{
			Common:    models.Common{ID: 4, CreatedAt: now, UpdatedAt: now},
			Name:      "Fahrul",
			Email:     "Fahrul@gmail.com",
			Password:  "123456",
			CreatedBy: 1,
		},
		{
			Common:   models.Common{ID: 5, CreatedAt: now, UpdatedAt: now},
			Name:     "Dito",
			Email:    "Dito@gmail.com",
			Password: "123456",
		},
	}
	if err := db.Create(&users).Error; err != nil {
		log.Printf("cannot seed data users, with error %v\n", err)
	}
	log.Println("success seed data users")
}