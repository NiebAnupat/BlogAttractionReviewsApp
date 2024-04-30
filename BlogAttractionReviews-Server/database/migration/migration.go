package main

import (
	"log"

	"github.com/NiebAnupat/BlogAttractionReviewsApp/Server/config"
	database "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/database"
	"github.com/NiebAnupat/BlogAttractionReviewsApp/Server/entities"
	"gorm.io/gorm"
)

func main() {
	conf := config.ConfigGetting()
	db := database.NewPostgresDatabase(conf.Database)

	tx := db.Connect().Begin()

	dropAllTable(tx)

	userMigration(tx)
	blogMigration(tx)
	userLikeMigration(tx)
	userFavoriteMigration(tx)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}

	log.Println("Migration completed")
}

func userMigration(tx *gorm.DB) {
	_ = tx.Migrator().CreateTable(&entities.User{})
}

func blogMigration(tx *gorm.DB) {
	_ = tx.Migrator().CreateTable(&entities.BlogPost{})
	_ = tx.Migrator().CreateTable(&entities.BlogContent{})
}

func userLikeMigration(tx *gorm.DB) {
	_ = tx.Migrator().CreateTable(&entities.UserLike{})
}

func userFavoriteMigration(tx *gorm.DB) {
	_ = tx.Migrator().CreateTable(&entities.UserFavorite{})
}

func dropAllTable(tx *gorm.DB) {
	isTableExist := tx.Migrator().HasTable(&entities.User{})
	if isTableExist {
		_ = tx.Migrator().DropTable(&entities.User{})
		_ = tx.Migrator().DropTable(&entities.BlogPost{})
		_ = tx.Migrator().DropTable(&entities.BlogContent{})
		_ = tx.Migrator().DropTable(&entities.UserLike{})
		_ = tx.Migrator().DropTable(&entities.UserFavorite{})

		log.Println("All table has been dropped")
	}
}
