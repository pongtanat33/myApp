package db

import (
  "gorm.io/gorm"
  "gorm.io/driver/sqlserver"
  "fmt"
  "os"

  "github.com/joho/godotenv"
)

var isMigrate = false

func DB() *gorm.DB {
	fmt.Fprintf(os.Stderr,"h")
	dbhost:= os.Getenv("DB_HOST")
	dbpass:= os.Getenv("MSSQL_PASSWORD")
	dbport:= os.Getenv("DB_PORT")
	dbname:= os.Getenv("DB_NAME")
	fmt.Println("DB",dbhost)
	if(dbhost == ""){
		LoadEnv()
		get := GetEnvWithKey
		dbhost= get("DB_HOST_DEV")
		dbpass= get("MSSQL_PASSWORD")
		dbport= get("DB_PORT")
		dbname= get("DB_NAME")
	}
	//dsn := "sqlserver://sa:P@ssw0rd@db:1433?database=demo"
	
	dsn := fmt.Sprintf("sqlserver://sa:%v@%v:%v?database=%v",dbpass,dbhost,dbport,dbname)
	fmt.Println(dsn)
    


	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect db")
	}

	// if !isMigrate{
	// 	db.AutoMigrate(&UserDB{})
	// 	isMigrate = true
	// }
	return db
}

func Migrate(){
	db := DB()
	db.AutoMigrate(&UserDB{})
	db.AutoMigrate(&TodoDB{})
}

func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

func LoadEnv(){
	err := godotenv.Load(".env")
	if err != nil{
		fmt.Println("cannot load file")
	}
	fmt.Println("load file success")
}