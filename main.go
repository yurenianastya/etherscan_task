package main

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"testtask/methods"
)

func loadConfig(path string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := loadConfig(".")
	if err != nil {
		return
	}
	e := echo.New()
	// route
	e.GET("/api/block/:block_number/total", methods.GetTransactionsAmount)
	e.Logger.Fatal(e.Start(viper.GetString("PORT")))
}
