package main

import (
	"flag"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
	"social_club/app/handlers"
	"social_club/app/middleware"
	"social_club/app/repositories"
	"social_club/app/repositories/impl"
	"social_club/app/usecases/impl"
	"strings"
)

type Flags struct {
	PostgresDatabase bool
}

func ParseFlag() (flags Flags) {
	flag.BoolVar(&flags.PostgresDatabase, "postgres", false, "use postgresql database")
	flag.Parse()
	return
}

type Config struct {
	postgresHost     string
	postgresUser     string
	postgresPassword string
	postgresDbName   string
	postgresPort     string

	logFile string
}

func ParseConfig() (conf Config) {
	viper.AddConfigPath("./cmd/api/")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	conf.postgresHost = viper.GetString("postgresHost")
	conf.postgresUser = viper.GetString("postgresUser")
	conf.postgresPassword = viper.GetString("postgresPassword")
	conf.postgresDbName = viper.GetString("postgresDbName")
	conf.postgresPort = viper.GetString("postgresPort")

	conf.logFile = viper.GetString("logFile")
	return
}

func main() {
	flags := ParseFlag()

	conf := ParseConfig()
	f, err := os.Create(conf.logFile)
	if err != nil {
		log.Fatal(err)
		return
	}

	gin.DefaultWriter = io.MultiWriter(f)
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST"}
	config.AllowCredentials = true

	var repo repositories.Repository

	if flags.PostgresDatabase {
		conn, err := pgx.ParseConnectionString(strings.Join([]string{"host=", conf.postgresHost, " user=", conf.postgresUser, " password=", conf.postgresPassword, " dbname=", conf.postgresDbName, " port=", conf.postgresPort}, ""))
		if err != nil {
			log.Fatal(err)
			return
		}

		db, err := pgx.NewConnPool(pgx.ConnPoolConfig{
			ConnConfig:     conn,
			MaxConnections: 100,
			AfterConnect:   nil,
			AcquireTimeout: 0,
		})
		if err != nil {
			log.Fatal(err)
			return
		}

		defer db.Close()
		repo = repositories_impl.MakePostgresRepository(db)
	} else {
		repo = repositories_impl.MakeInMemoryRepository()
	}

	useCase := usecases_impl.MakeUseCase(repo)
	handler := handlers.MakeHandler(useCase)

	router.Use(cors.New(config))
	router.Use(middleware.CheckError())

	routes := router.Group("/")
	{
		routes.POST("/msg", handler.CreateMessage)
		routes.GET("/info", handler.GetInformation)
	}

	err = router.Run(":5000")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
