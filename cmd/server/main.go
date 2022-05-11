package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/greendrop/todo-graphql-go-sample/domain/entity"
	"github.com/greendrop/todo-graphql-go-sample/infrastructure/persistence"
	"github.com/greendrop/todo-graphql-go-sample/interface/graph"
	graphgenerated "github.com/greendrop/todo-graphql-go-sample/interface/graph/generated"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const defaultPort = "8080"

func main() {
	logger, _ := prepareLogger()
	defer closeLogger(logger)

	zap.L().Info("Start server")

	appConfig, err := loadAppConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	persistence.AppConfig = appConfig

	gormDB, err := openDatabase()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer closeDatabase(gormDB)
	persistence.GormDB = gormDB

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graphgenerated.NewExecutableSchema(graphgenerated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func prepareLogger() (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	zap.ReplaceGlobals(logger)
	return logger, err
}

func closeLogger(logger *zap.Logger) {
	logger.Sync()
}

func loadAppConfig() (*entity.AppConfig, error) {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
	}

	var appConfig entity.AppConfig
	// configPath := "./configs/app_config." + appEnv + ".toml"
	// _, err := toml.DecodeFile(configPath, &appConfig)
	// if err != nil {
	// 	return nil, err
	// }

	appConfig.AppEnv = appEnv
	appConfig.Database.Url = os.Getenv("DATABASE_URL")

	return &appConfig, nil
}

func openDatabase() (*gorm.DB, error) {
	dsn := strings.Replace(persistence.AppConfig.Database.Url, "mysql://", "", 1)
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	zap.L().Info("Opened database")

	return gormDB, nil
}

func closeDatabase(gormDB *gorm.DB) {
	sqlDB, _ := gormDB.DB()

	if sqlDB != nil {
		sqlDB.Close()
		zap.L().Info("Closed database")
	}
}
