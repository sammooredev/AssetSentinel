package main

import (
	"database/sql"
    "fmt"
    "bufio"
    "os"
    "log"
    "regexp"
    "github.com/gookit/color"
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
func DBInsertData(prog_name string, endpoints_path string){
    //if progname is table + endpoints exists: do insert data into table
    //else error
    //check if specified endpoints files exists
    if _, err := os.Stat("./" + endpoints_path); err == nil {
        fmt.Print("Endpoints selected: " + endpoints_path + "\n")
        os.Exit(1)
    } else {
        fmt.Print("You need to specify a program name and path to your endpoints file\n" +
                "Usage:\n" +
                "\t$ assetspy manage foobar endpoints.txt\n")
    }
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

func CreateDB(prog_name string, endpoints_path string) {
    //check if db directory has been deleted for some reason
    if _, err := os.Stat("./db"); os.IsNotExist(err){
        fmt.Print("Seems your database directory (./db) was deleted... Creating a new one.\n")
        newdb := os.MkdirAll("./db", 0755)
        if newdb != nil {
            log.Fatal(err)
        }
        fmt.Print("Database directory created!.\n")
    }
    //check if db exists?
    if _, err := os.Stat("./db/assetspy-database.db"); err == nil {
		return
    } else {
        // this code runs if the db doesnt exists. Creates a db for program
        fmt.Print("Creating a new database for assetspy...\n")
        database, err := os.Create("./db/assetspy-database.db")
        if err != nil {
            log.Fatal(err)
        }
        database.Close()
        fmt.Print("Database created, you can now add programs and endpoints!\n" + 
                "Usage:\n" +
                "*note: program name must be one word, and only letters A-Z, a-z*\n" +
                "\n\t$ assetspy manage update ExampleProgramName /path/to/endpoints.txt\n")
	}
}
func CreateTable(prog_name string) {
    //if db doesnt exist tell the user how to create it
    if _, err := os.Stat("./db/assetspy-database.db"); os.IsNotExist(err) {
        fmt.Print("Error: you havent created a database for assetspy yet!\n" +
        "To create a database:\n" +
        "\t$ assetspy manage update\n" +
        "\nAfter running the above, re-run this command to add a program to the database.\n" +
        "\t$ assetspy manage update example\n")
    } else {
    //if db does exist, then take the progname and create a table with it
        database, err := sql.Open("sqlite3", "./db/assetspy-database.db")
        if err != nil {
            log.Fatal(err)
        }
        defer database.Close()
        //regex check on program name, dont anyone breaking their databases 
        program_name_regex := regexp.MustCompile("^[a-zA-Z0-9]*$")
        if program_name_regex.FindStringIndex(prog_name) != nil  {
            fmt.Print("valid name\n")
        } else {
            fmt.Print(program_name_regex.FindString(prog_name))
            fmt.Print("Program name must be a single word with no spaces, or symbols.\nExample: Program1\n\nExiting...\n")
            os.Exit(1)
        }
        //check if table already exists (queries tables and does a comparison)
        doesTableExist, err := database.Query("SELECT name FROM sqlite_master WHERE type='table' AND name NOT LIKE 'sqlite_%'")
        if err != nil {
            log.Fatal(err)
        }
        defer doesTableExist.Close()
        var table string
        for doesTableExist.Next() {
            err := doesTableExist.Scan(&table)
            if err != nil {
                log.Fatal(err)
            }
            if table == prog_name {
                fmt.Print("Program already exists in database!\nExiting!\n")
                os.Exit(1)
            }
        }
        statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS " + prog_name +" (id INTEGER PRIMARY KEY, endpoint TEXT, request TEXT, response TEXT)")
        if err != nil {
            fmt.Print("Adding a program to the database failed.\n" +
                "\t1. The program could already be in the database\n" +
                "\t2. You are trying a name like \"foo bar\". It must be one word like foobar\n\n")
            log.Fatal(err)
        }
        statement.Exec()
        fmt.Print(prog_name + " added to database!\n")
    }
}
// function that lists tables (programs) in the database
func ListDatabases(prog_name string, request string){
    switch {
        //case to handle $ assetspy list
        case len(prog_name) > 0:
            fmt.Print("\nTargets you're managing:\n\n")
            database, err := sql.Open("sqlite3", "./db/assetspy-database.db")
            if err != nil {
                log.Fatal(err)
            }
            defer database.Close()
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
                fmt.Println("\t" + table)
            }
            err = res.Err()
            if err != nil {
                log.Fatal(err)
            }
    }
}


func main() {
    ToolNameColor := color.New(color.Bold, color.Yellow)
    StepsColor := color.New(color.Bold, color.Red)
    if len(os.Args) == 1 {
        ToolNameColor.Print("\n-- AssetSpy --\n")
        StepsColor.Print("\nFirst time usage steps:\n")
        StepsColor.Print("\nStep 1: Create a database:\n")
        fmt.Print("\t$ assetspy manage update\n")
        StepsColor.Print("\nStep 2: Add Programs to your database:    *note: program names must be one word!*\n")
        fmt.Print("\t$ assetspy manage update ExampleProgramName\n")
        StepsColor.Print("\nStep 3: Add endpoints for monitoring to your program\n")
        fmt.Print("\t$ assetspy manage update ExampleProgramName /path/to/endpoints.txt\n")
        StepsColor.Print("\nStep 4: Run a scan against your program. This has many options, see more scan options below\n")
        fmt.Print("\t$ assetspy scan ExampleProgramName all\n")
        StepsColor.Print("\nStep 5: Re-Scan at a later date:\n")
        fmt.Print("\t$ assetspy scan ExampleProgramName all\n")
        StepsColor.Print("\nStep 6: Run diff to generate a report highlighting changes in the responses\n")
        fmt.Print("\t$ assetspy diff ExampleProgrmName <date> <date>\n\n")
        os.Exit(1)
    }
    if len(os.Args[1:]) == 1 {
        fmt.Print("\n-- AssetSpy --\n" +
        "\nFirst time usage steps:\n" +
        "\nStep 1: Create a database:\n" +
        "\t$ assetspy manage update\n" +
        "\nStep 2: Add Programs to your database:    *note: program names must be one word!*\n" +
        "\t$ assetspy manage update ExampleProgramName\n" +
        "\nStep 3: Add endpoints for monitoring to your program\n" +
        "\t$ assetspy manage update ExampleProgramName /path/to/endpoints.txt\n" +
        "\nStep 4: Run a scan against your program. This has many options, see more scan options below\n" +
        "\t$ assetspy scan ExampleProgramName\n")
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
                CreateDB("","none")
            }
            // handles "manage update <program>"
            if(arg2 == "update" && len(arg3) > 0 && len(arg4) == 0){
                CreateTable(arg3)
            }
            // handles "manage list"
            if(arg2 == "list"){
                fmt.Print("Mode selected: manage list\n")
                // runs ListDBs
                ListDatabases("f","")
            }

        case arg1 == "scan": //scan handler
            fmt.Print("\nMode selected: scan")
        case arg1 == "diff":
            fmt.Print("\nMode selected: diff")
    }
}
