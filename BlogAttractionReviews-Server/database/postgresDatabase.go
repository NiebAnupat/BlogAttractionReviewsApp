package databases

import (
	"fmt"
	"log"
	"sync"

	"github.com/NiebAnupat/BlogAttractionReviewsApp/Server/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	*gorm.DB
}

var (
	postgresDatabaseInstance *postgresDatabase
	once                     sync.Once
)

func NewPostgresDatabase(conf *config.Database) Database {
	once.Do(func() {
		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s search_path=%s",
			conf.Host, conf.Port, conf.User, conf.Password, conf.DBName, conf.SSLMode, conf.Schema)
		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		log.Printf("Connected to database : %s\n", conf.DBName)
		postgresDatabaseInstance = &postgresDatabase{conn}
	})
	return postgresDatabaseInstance
}

func (p *postgresDatabase) Connect() *gorm.DB {
	return p.DB
}
