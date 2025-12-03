package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	// == "ariga.io/atlas/sql/postgres"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"

)

type Server struct {
	DB *gorm.DB
	Router *mux.Router
}

type AppConfig struct{
	AppName string
	AppEnv string
	AppPort string
}

type DBConfig struct{
	DBHost string
	DBUser string
	DBPassword string
	DBName string
	DBPort string
}

func (server *Server) Initialize (appConfig AppConfig, dbConfig DBConfig){
	fmt.Println( "welcome to "+ appConfig.AppName )

	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbConfig.DBHost ,dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBPort)
	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// server.DB, err = gorm_postgres.Open(dsn), &gorm.Config{} //  Kode baru

	if err != nil {
		panic("Failed on conecting to the datahase")
	}

	server.Router = mux.NewRouter()
	server.InitializeRoutes()
}

func (server *Server ) Run (addr string){
	fmt.Println("Listening to port ", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func getEnv(key, fullback string)string{
	if value, ok:= os.LookupEnv(key); ok{
		return value
	}
	return fullback
}

func Run(){
	var server = Server{}
	var appConfig = AppConfig{}
	var dbConfig = DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error on loading. env file")
	}

	appConfig.AppName =getEnv("APP_NAME", "goTOKO")
	appConfig.AppEnv =getEnv("APP_ENV","dev" )
	appConfig.AppPort = getEnv("APP_PORT","9000")

	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
dbConfig.DBUser = getEnv("DB_USER", "user")
dbConfig.DBPassword = getEnv("DB_PASSWORD", "12345")
dbConfig.DBName = getEnv("DB_NAME", "dbname")
dbConfig.DBPort = getEnv("DB_PORT", "5432")

		
	server.Initialize(appConfig , dbConfig)
	server.Run(":" +appConfig.AppPort)
}

