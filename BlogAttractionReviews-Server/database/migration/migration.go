package main

import (
	"log"

	"github.com/NiebAnupat/BlogAttractionReviewsApp/Server/config"
	database "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/database"
	"github.com/NiebAnupat/BlogAttractionReviewsApp/Server/entities"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"gorm.io/gorm"
)

func main() {
	conf := config.ConfigGetting()
	db := database.NewPostgresDatabase(conf.Database)

	tx := db.Connect().Begin()

	dropAllTableIfExits(tx)
	clearS3Buckets(conf)

	// userMigration(tx)
	// blogMigration(tx)
	// userLikeMigration(tx)
	// userFavoriteMigration(tx)

	// auto migration
	if err := tx.AutoMigrate(&entities.User{}, &entities.BlogPost{}, &entities.BlogContent{}, &entities.UserLike{}, &entities.UserFavorite{}); err != nil {
		tx.Rollback()
		panic(err)
	}

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

func dropAllTableIfExits(tx *gorm.DB) {
	tables, err := tx.Migrator().GetTables()
	if err != nil {
		panic(err)
	}
	if len(tables) > 0 {
		for _, table := range tables {
			_ = tx.Migrator().DropTable(table)
			log.Printf("----| Table %s has been dropped", table)
		}
		log.Println("All table has been dropped")
	}
}

func clearS3Buckets(conf *config.Config) {
	s3Client := s3.New(session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(conf.AWS.Region),
		Credentials: credentials.NewStaticCredentials(conf.AWS.S3.AccessKeyID, conf.AWS.S3.SecretAccessKey, ""),
	})))

	iter := s3manager.NewDeleteListIterator(s3Client, &s3.ListObjectsInput{
		Bucket: aws.String(conf.AWS.S3.Bucket),
	})

	if err := s3manager.NewBatchDeleteWithClient(s3Client).Delete(aws.BackgroundContext(), iter); err != nil {
		panic(err)
	}

	log.Println("All files in S3 bucket has been deleted")
}
