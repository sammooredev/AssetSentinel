package main

import (
	"database/sql"
    "fmt"
    "bufio"
    "os"
    "log"
	_ "github.com/mattn/go-sqlite3"
)

// AssetSpy
//
//       manage, scan, diff
// 
// manage usage: $ asset-spy manage list
//   Update - used to create an asset spy database
//   List   - used to list databases
//
// scan usage: $ assetspy scan <db-name>
func ManageUpdateInsertData(prog_name string, endpoints_path string){
    fmt.Print("Insertin'\n")
    //open endpoints.txt for reading
    endpoints_file, err := os.Open(endpoints_path)
    if err != nil {
        log.Fatal(err)
    }
    defer endpoints_file.Close()

    scanner := bufio.NewScanner(endpoints_file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    //statement, _ = database.Prepare("INSERT INTO ReqAndResp (Request, Response) VALUES (?, ?)")
	//statement.Exec("foo1", "foo2")

	//rows, _ := database.Query("SELECT id, endpoint, request, response FROM ReqAndResp")
	//var id int
	//var Request string
	//var Response string
	//for rows.Next() {
	//	rows.Scan(&id, &Request, &Response)
	//	fmt.Println("id: " + strconv.Itoa(id) + "\n" + "endpoint: " + endpoint + "request: " + request + "\nresponse: " + response)
    //}
}

func ManageUpdate(prog_name string, endpoints_path string) {
    //check if db directory has been deleted for some reason
    if _, err := os.Stat("./dbs"); os.IsNotExist(err){
        fmt.Print("Seems your database directory (./dbs) was deleted... Creating a new one.\n")
        newdb := os.MkdirAll("./dbs", 0755)
        if newdb != nil {
            log.Fatal(err)
        }
        fmt.Print("Database directory created!.\n")
    }
    //check if db exists?
    if _, err := os.Stat("./dbs/assetspy-database.db"); err == nil {
		//check if specified endpoints files exists
        if _, err := os.Stat("./" + endpoints_path); err == nil {
            fmt.Print("Endpoints selected: " + endpoints_path + "\n")
            ManageUpdateInsertData(prog_name, endpoints_path)
			os.Exit(1)
		} else {
			fmt.Print("DB exists, but assetspy needs a list of endpoints to update the db with\n" +
					"Usage:\n" +
					"\t$ assetspy manage foobar endpoints.txt\n")
        }
	} else {
        // this code runs if the db doesnt exists. Creates a db for program
        fmt.Print("Creating a new database for assetspy...")
        database, err := os.Create("./dbs/assetspy-database.db")
        if err != nil {
            log.Fatal(err)
        }
        database.Close()
        fmt.Print("Database created, you can now add programs and endpoints!")
		//create tables
		//statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS foobar (id INTEGER PRIMARY KEY, endpoint TEXT, request TEXT, response TEXT)")
		//statement.Exec()
        //fmt.Print("created table...")
	}
}
// function that lists databases created by the user
func ListDatabases(prog_name string, request string){
    switch {
        //case to handle $ assetspy list
        case len(prog_name) > 0:
            fmt.Print("Targets you're managing:\n")
            database, _ := sql.Open("sqlite3", "./dbs/assetspy-database.db")
            //query all tables in DB
            res, err := database.Query("SELECT name FROM sqlite_master WHERE type='table' AND name NOT LIKE 'sqlite_%'")
            if err != nil {
                log.Fatal(err)
            }
            defer res.Close()
            var table string
            for res.Next() {
                err := res.Scan(&table)
                if err != nil {
                    log.Fatal(err)
                }
                log.Println("\t" + table)
            }
            err = res.Err()
            if err != nil {
                log.Fatal(err)
            }
    }
}


func main() {
    if len(os.Args) == 1 {
        fmt.Print("to get started run:\n" +
        "\t$ assetspy manage update\n" +
        "\t$ assetspy manage list\n" +
        "\t$ assetspy scan \n" +
        "\t$ assetspy diff\n")
        os.Exit(1)
    }
    if len(os.Args[1:]) == 1 {
        fmt.Print("try one of the following:\n" +
        "\t$ assetspy manage update\n" +
        "\t$ assetspy manage list\n" +
        "\t$ assetspy scan \n" +
        "\t$ assetspy diff\n")
        os.Exit(1)
    }

    //args handling
    var arg1 string
    var arg2 string
    var arg3 string
    var arg4 string
    //var arg5 string
    //switch case avoids any index out of range errors
    switch {
        case len(os.Args[1:]) == 1:
            arg1 = os.Args[1]
        case len(os.Args[1:]) == 2:
            arg1 = os.Args[1]
            arg2 = os.Args[2]
        case len(os.Args[1:]) == 3:
            arg1 = os.Args[1]
            arg2 = os.Args[2]
            arg3 = os.Args[3]
       case len(os.Args[1:]) == 4:
            arg1 = os.Args[1]
            arg2 = os.Args[2]
            arg3 = os.Args[3]
            arg4 = os.Args[4]
       case len(os.Args[1:]) == 5:
            arg1 = os.Args[1]
            arg2 = os.Args[2]
            arg3 = os.Args[3]
           // arg4 = os.Args[4]
           // arg5 = os.Args[5]
    }
    //arg1 = manage, scan or diffs
    //arg2 = (manage: update, list) (scan: db-name) (diff: db-name)
    //arg3 := os.Args[3] //arg3 = (manage: db-name) (scan: all, js, html) (diff: date1)
    //arg4 := os.Args[4] //arg4 = (manage: endpoints.txt) (scan: ) (diff: date2)

    //switch statment to handle args
    switch {
        // handles "manage"
        case arg1 == "manage":
            // handles "manage update"
            if(arg2 == "update" && len(arg3) == 0){
                ManageUpdate("","none")
            }
            // handles "manage update <program>"
            if(arg2 == "update" && len(arg3) > 0){
                fmt.Print("Mode selected: manage update\n")
                if len(arg4) > 0 {
					ManageUpdate(arg3, arg4)
				}
            }
            // handles "manage list"
            if(arg2 == "list"){
                fmt.Print("Mode selected: manage list\n")
                // runs ListDBs
                ListDatabases("","")
            }

        case arg1 == "scan": //scan handler
            fmt.Print("\nMode selected: scan")
        case arg1 == "diff":
            fmt.Print("\nMode selected: diff")
    }
}
