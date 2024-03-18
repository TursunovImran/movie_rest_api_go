package main

import (
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/spf13/viper"
	"github.com/joho/godotenv"

	"github.com/TursunovImran/movie_rest_api_go"
	"github.com/TursunovImran/movie_rest_api_go/pkg/handler"
	"github.com/TursunovImran/movie_rest_api_go/pkg/repository"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error occurred while initializing the config %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	_ = db

	r := handler.InitRouter(db)
	srv := new(movierestapigo.Server)
	
	if err := srv.Run(viper.GetString("port"), r); err != nil {
		log.Fatalf("error ocured while running http server %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}