package postgres

import (
	"fmt"
	"github.com/ivchip/go-meli-filter-ip/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once
)

func New() {
	once.Do(func() {
		var err error
		driverStr := os.Getenv("DRIVER_NAME")
		connStr := os.Getenv("CONNECT_DB")
		db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}
		fmt.Println("Connect to", driverStr)
	})
}

func DAOIpBlocking() domain.IpBlockingUseCases {
	return newPsqlIpBlocking(db)
}
