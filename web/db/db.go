package db

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var db *gorm.DB

type DataSource struct {
	DriverName string
	Host       string
	Port       string
	Database   string
	Username   string
	Password   string
}

func GetDb() *gorm.DB {
	var err error
	if db != nil {
		return db
	}
	var dataSource DataSource
	dataSource.DriverName = viper.GetString("datasource.driverName")
	dataSource.Host = viper.GetString("datasource.host")
	dataSource.Port = viper.GetString("datasource.port")
	dataSource.Database = viper.GetString("datasource.database")
	dataSource.Username = viper.GetString("datasource.username")
	dataSource.Password = viper.GetString("datasource.password")
	log.Print("a")
	log.Println(dataSource.DriverName)
	//dataSource.DriverName = "mysql"
	//dataSource.Host = "localhost"
	//dataSource.Port = "3306"
	//dataSource.Database = "testoceango"
	//dataSource.Username = "root"
	//dataSource.Password = "root"

	if dataSource.DriverName == "" {
		return nil
	}

	// TODO Need Refactoring
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Asia/Shanghai", dataSource.Username, dataSource.Password, dataSource.Host, dataSource.Port, dataSource.Database)
	//args := "root:root@tcp(localhost:3306)/testoceango?charset=utf8"
	db, err = gorm.Open(dataSource.DriverName, args)
	if err != nil {
		log.Print(err)
		panic(err.Error())
	}

	// TODO is there need close db
	return db
}
