package initialize

import (
	"gin-practice/config"
	"gin-practice/pkg/models"
	"log"
	"os"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetGorm() *gorm.DB {
	p := config.SELF_CONFIG.Postgres
	dsnSlice := []string{
		"user=",
		p.User,
		" password=",
		p.Password,
		" dbname=",
		p.Dbname,
		" host=",
		p.Host,
		" port=",
		p.Port,
		" sslmode=",
		p.SslMode,
		" search_path=",
		p.Schema,
		" TimeZone=",
		p.TimeZone,
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      false,
			Colorful:                  false,
		},
	)
	dsn := strings.Join(dsnSlice, "")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("数据库连接失败, error=" + err.Error())
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(p.MaxOpenConns)
	sqlDB.SetMaxIdleConns(p.MaxIdleConns)
	return db
}

func CreateTables(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comment{},
	)
}
