package seeder

import (
	"github.com/hikayat13/alterra-agcm/day-2/submission/config"
	"gorm.io/gorm"
)

type seed struct {
	DB *gorm.DB
}

func NewSeeder() *seed {
	return &seed{config.DB}
}

func (s *seed) SeedAll() {
	userSeeder(s.DB)
}

func (s *seed) DeleteAll() {
	s.DB.Exec("DELETE FROM users")
}
