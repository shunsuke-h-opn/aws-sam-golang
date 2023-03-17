package eventbus

import (
	"fmt"
	// "gorm.io/gorm"
	// "gorm.io/gorm/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// "gorm.io/driver/postgres"

	"os"
)

var db *gorm.DB
var err error

func init() {
	fmt.Println("----------------================================------------------")
	fmt.Println(os.Getenv("DATABASE_PORT"))
	fmt.Println(os.Getenv("DATABASE_PASSWORD"))
	db, err = gorm.Open(
		"postgres",
		"host="+os.Getenv("DATABASE_HOST")+
			" user="+os.Getenv("DATABASE_USER")+
			" dbname="+os.Getenv("DATABASE_NAME")+
			" password="+os.Getenv("DATABASE_PASSWORD")+
			" port="+os.Getenv("DATABASE_PORT")+
			" sslmode=disable"+
			" connect_timeout=5")

	// dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=disable connect_timeout=5",
	// 	os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_HOST"))
	// _, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
	// 	Logger: logger.Default.LogMode(logger.Info),
	// })

	// fmt.Println(dsn)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	db.LogMode(true)

	// db.AutoMigrate(&Event{})
	// db.AutoMigrate(&Subscription{})
}
