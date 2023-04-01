package main

import (
	"fmt"

	"github.com/amajakai14/admin-renewal/internal/db"
)

func Run() error {
	fmt.Println("starting an application")
	db, err := db.NewDatabase()
	if err != nil {
		return err
	}

	if err := db.MigrateDB(); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
