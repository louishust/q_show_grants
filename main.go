package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jessevdk/go-flags"
	"os"
)

const VERSION = "1.0.0"

var con *sql.DB = nil

var options struct {
	Version bool   `short:"v" long:"version" description:"Print version"`
	Host    string `long:"host" description:"Server hostname or IP"`
	Port    string `long:"port" description:"Server port" default:"3306"`
	User    string `long:"user" description:"Database user"`
	Pass    string `long:"pass" description:"Password for user"`
}

func exitWithMessage(message string) {
	fmt.Println("Error:", message)
	os.Exit(1)
}

func initOptions() {
	_, err := flags.ParseArgs(&options, os.Args)

	if err != nil {
		os.Exit(1)
	}

	if options.Version {
		fmt.Printf("q_show_grants v%s\n", VERSION)
		os.Exit(0)
	}
}

func initConnection() {
	conn, err := sql.Open("mysql", options.User+":"+options.Pass+"@tcp("+options.Host+":"+options.Port+")/")
	if err != nil {
		exitWithMessage(err.Error())
	}
	con = conn
}

func main() {
	initOptions()
	initConnection()
	if con != nil {
		defer con.Close()
	}

	rows, err := con.Query("SHOW GRANTS")
	if err != nil {
		exitWithMessage(err.Error())
	}
	var row string
	for rows.Next() {
		err = rows.Scan(&row)
		if err != nil {
			exitWithMessage(err.Error())
		}
		fmt.Println(row)
	}
}
