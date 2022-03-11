package tests

import (
	"log"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestGo(t *testing.T) {
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./secrets")
	viper.AddConfigPath("../secrets")
	viper.ReadInConfig()

	t.Log("Content: ", viper.Get("demo.key"))
	assert.NotNil(t, viper.Get("demo.key"))
	assert.NotEmpty(t, viper.Get("demo.key"))

	db := NewSqlXInstance(viper.GetString("demo.db"))
	err := db.DB.Ping()
	t.Log(err)
	assert.Nil(t, err)
}

func NewSqlXInstance(databaseURL string) *sqlx.DB {
	dbx, err := sqlx.Open("postgres", databaseURL)
	if err != nil {
		log.Panic(err)
	}

	if err := dbx.Ping(); err != nil {
		log.Fatal(err)
	}

	dbx.SetConnMaxLifetime(time.Minute * 5)
	dbx.SetMaxIdleConns(0)
	dbx.SetMaxOpenConns(8)
	return dbx
}
