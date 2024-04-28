package main

import (
	"flag"
	"fmt"
	goHttp "net/http"

	"backend/data/migrate"
	"backend/firebase"

	"backend/config"
	"backend/data"
	"backend/http"
)

var (
	RunPopulateFlag    bool
	SkipMigrationsFlag bool
)

// CheckCommandLineArguments checks if any command-line arguments are passed and runs the corresponding functions
// Returns a boolean indicating whether the program should continue running or not
func CheckCommandLineArguments() {

	// Parse command-line arguments and flags
	flag.BoolVar(&RunPopulateFlag, "run-populate", false, "Populate the database with dummy data")
	flag.BoolVar(&SkipMigrationsFlag, "skip-migrations", false, "Skip running database migrations")
	flag.Parse()

}

// main is the entry point of the application
func main() {

	// load the config from the .env file
	loadedConfig, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// check command line arguments
	CheckCommandLineArguments()

	// connect to the database
	db, err := data.ConnectToDB()
	data.DB = db

	// initialize firebase
	err = firebase.InitFirebase()
	if err != nil {
		panic(err)
	}

	// run all the database migrations
	if !SkipMigrationsFlag {
		if err = migrate.RunMigrations(); err != nil {
			panic(err)
		}
	}

	// run all the database population
	if RunPopulateFlag {
		if err = migrate.PopulateDatabase(); err != nil {
			panic(err)
		}
		return
	}

	// initialize all the routes
	r := http.InitRoutes()

	address := fmt.Sprintf(
		":%s",
		loadedConfig.ServerPort,
	)

	// start the server
	if err = goHttp.ListenAndServe(address, r); err != nil {
		return
	}
}
