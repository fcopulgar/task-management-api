package persistence

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Models = []interface{}{
	&UserModel{},
	&SessionModel{},
	&TaskModel{},
	&CommentModel{},
}

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func (c Config) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode,
	)
}

func NewConnection(cfg Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error al conectar a postgresql: %w", err)
	}
	return db, nil
}

func RunMigrations(db *gorm.DB) error {
	for _, model := range Models {
		if err := db.AutoMigrate(model); err != nil {
			return fmt.Errorf("error en AutoMigrate de %T: %w", model, err)
		}
	}
	log.Println("AutoMigrate completado correctamente")
	return nil
}
