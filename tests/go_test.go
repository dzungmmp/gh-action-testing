package tests

import (
	"testing"

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
}
