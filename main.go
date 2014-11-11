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
	Help    bool   `long:"help" description:"Show this help message"`
	Version bool   `short:"v" long:"version" description:"Print version"`
	Host    string `short:"h" long:"host" description:"Server hostname or IP" default:"localhost"`
	Port    string `short:"P" long:"port" description:"Server port" default:"3306"`
	User    string `short:"u" long:"user" description:"Database user"`
	Pass    string `short:"p" long:"password" description:"Password for user"`
}

func exitWithMessage(message string) {
	fmt.Println("Error:", message)
	os.Exit(1)
}

func initOptions() {
	p := flags.NewParser(&options, flags.Default&^flags.HelpFlag)
	_, err := p.Parse()
	if err != nil {
		os.Exit(1)
	}

	if options.Help {
		p.WriteHelp(os.Stdout)
		os.Exit(0)
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
	fmt.Println("\nPrivileges for user", options.User)
	fmt.Println("--------------------------------------------------------------------- ")
	var row string
	for rows.Next() {
		err = rows.Scan(&row)
		if err != nil {
			exitWithMessage(err.Error())
		}
		fmt.Println(row)
	}
	fmt.Println("---------------------------------------------------------------------\n ")
}
