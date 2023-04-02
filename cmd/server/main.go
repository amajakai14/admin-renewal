package main

import (
	"fmt"

	"github.com/amajakai14/admin-renewal/internal/db"
	appUser "github.com/amajakai14/admin-renewal/internal/user"
	httpTransport "github.com/amajakai14/admin-renewal/internal/transport/http"
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

	userService := appUser.NewService(db)
	services := httpTransport.Services{UserService: userService}

	httpHandler := httpTransport.NewHandler(services)
	if err := httpHandler.Serve(); err != nil {
		return err
	}
	
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
