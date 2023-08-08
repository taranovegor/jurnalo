package cmd

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/taranovegor/com.jurnalo/internal/config"
	"github.com/taranovegor/com.jurnalo/internal/container"
)

func Init(scope string) container.ServiceContainer {
	fmt.Println(fmt.Sprintf("[cmd/%s] %s! Version: %s", scope, config.AppName, config.Version))

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	sc, err := container.Init()
	if err != nil {
		panic(err)
	}

	return sc
}
