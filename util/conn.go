package util

import (
	"fmt"

	"github.com/spf13/viper"
)

func DBConnString() string {
	viper.SetEnvPrefix("DB")
	viper.AutomaticEnv()

	viper.SetDefault("HOST", "localhost")
	viper.SetDefault("PORT", "30000")
	viper.SetDefault("USER", "dbadmin")
	viper.SetDefault("PASSWORD", "admin123")
	viper.SetDefault("NAME", "sqlboiler")

	host := viper.GetString("HOST")
	port := viper.GetString("PORT")
	user := viper.GetString("USER")
	password := viper.GetString("PASSWORD")
	dbName := viper.GetString("NAME")

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName,
	)
}
