package main

import (
	"bank/handler"
	"bank/logs"
	"bank/repository"
	"bank/service"
	"fmt"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	initTimeZone()
	db := initDatabase()

	customerRepositoryDB := repository.NewCustomerRepositoryDB(db)
	accountRepositoryDB := repository.NewAccountRepository(db)
	// customerRepositoryMock := repository.NewCustomerRepositoryMock()
	// _ = customerRepositoryMock

	customerService := service.NewCustomerService(customerRepositoryDB)
	customerHandler := handler.NewCustomerHandler(customerService)

	accountService := service.NewAccountService(accountRepositoryDB)
	accountHandler := handler.NewAccountHandler(accountService)

	router := mux.NewRouter()

	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}/accounts", accountHandler.GetAccount).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}/accounts", accountHandler.NewAccount).Methods(http.MethodPost)

	// log.Printf("Banking service started at port %v", viper.GetString("app.port"))
	logs.Info("Banking service started at port " + viper.GetString("app.port"))
	http.ListenAndServe(fmt.Sprintf(":%v", viper.GetString("app.port")), router)

}

func initConfig() {
	viper.SetConfigName("")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv() // auto detect config on ENV
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}

func initDatabase() *sqlx.DB {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%v", viper.GetString("db.endpoint")))
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db

}
