package environment

import (
	"fmt"
	"github.com/DiegoJCordeiro/golang-study/chapter6/configs"
	"github.com/DiegoJCordeiro/golang-study/chapter6/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var (
	postgresInstance *gorm.DB
	sqliteInstance   *gorm.DB
	jwtSecret        string
	jwtExpiresIn     int32
)

func init() {

	env, err := configs.LoadConfigurationEnvironment("./cmd/server-api")

	if err != nil {
		log.Fatal(err)
	}

	configPostgres(env)
	configSQLite(env)
	configJwt(env)
	autoMigrateEntities()
}

func GetJwtSecret() string {
	return jwtSecret
}

func GetJwtExpiresIn() int32 {
	return jwtExpiresIn
}

func GetDBPostgres() *gorm.DB {

	return postgresInstance
}

func GetDBSQLite() *gorm.DB {

	return sqliteInstance
}

func configJwt(env *configs.ConfigurationEnvironment) {
	jwtSecret = env.JWTSecret
	jwtExpiresIn = env.JWTExpiresIn
}

func configPostgres(env *configs.ConfigurationEnvironment) {

	datasource := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", env.DBName, env.DBPass, env.DBHost, env.DBPort, env.DBName)

	db, errOpen := gorm.Open(postgres.Open(datasource), &gorm.Config{})

	if errOpen != nil {
		panic(errOpen)
	}

	postgresInstance = db
}

func configSQLite(env *configs.ConfigurationEnvironment) {

	datasource := fmt.Sprintf("%s", env.DBCacheName)

	db, errOpen := gorm.Open(sqlite.Open(datasource), &gorm.Config{})

	if errOpen != nil {
		panic(errOpen)
	}

	sqliteInstance = db
}

func autoMigrateEntities() {

	errPostgresAutoMigrate := postgresInstance.AutoMigrate(&entity.Role{}, &entity.User{})
	if errPostgresAutoMigrate != nil {
		panic(errPostgresAutoMigrate)
	}

	errMemAutoMigrate := postgresInstance.AutoMigrate(&entity.Role{}, &entity.User{})
	if errMemAutoMigrate != nil {
		panic(errMemAutoMigrate)
	}
}
