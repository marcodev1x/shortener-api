package main

import (
	"go-project/infra"

	"gorm.io/gorm"
)

var (
	DatabaseDomain *gorm.DB
	boot           infra.Bootstrap
)

func main() {
	env := boot.LoadEnv()

	db := boot.SetupDatabase(env)
	DatabaseDomain = db

	boot.RunServer()
}
