package db

import (
	"log"

	"golang-apiserver/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *db

// db DB params
type db struct {
	Con      *gorm.DB
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	TimeZone string
}

// NewDB create new db
func NewDB() db {
	return db{
		Host:     "localhost",
		User:     "test",
		Password: "test",
		DBName:   "test",
		Port:     "5432",
		TimeZone: "Asia/Tehran",
	}
}

// Connect db connect with params
func (db *db) Connect() *db {
	//TODO must be seperate for each database driver
	//TODO the config must be in default and set it from outside
	//TODO this section can be better programming method
	dsn := "host=" + db.Host
	dsn = dsn + " user=" + db.User
	dsn = dsn + " password=" + db.Password
	dsn = dsn + " dbname=" + db.DBName
	dsn = dsn + " port=" + db.Port
	dsn = dsn + " TimeZone=" + db.TimeZone
	newDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB Connected")
	db.Con = newDB
	DB = db
	return db
}

// AutoMigrate auto migrate of db
func AutoMigrate(dst interface{}) {
	err := DB.Con.AutoMigrate(dst)
	if err != nil {
		log.Println(err)
	}
}

// MigrateModels
func MigrateModels() {
	//TODO this function is not good place here. must be change into another location and structure
	AutoMigrate(&model.User{})
	AutoMigrate(&model.Product{})
}
