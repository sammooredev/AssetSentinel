package main

import (
    "fmt"
    "os"
)

// AssetSpy
//
//       manage, scan, diff
// 
// manage usage: $ asset-spy manage list
//   Update - used to create an asset spy database
//   List   - used to list databases
//
// scan usage: $ asset-spy scan <db-name>

func UpdateDB(prog_name string) {
    //check if db exists?
    fmt.Print("Program DB selcted: " + prog_name + "\n")
    if _, err := os.Stat("./dbs/" + prog_name); err == nil {
        fmt.Print("DB exists, need endpoints.txt to update db with\n")
    } else {
        fmt.Print("no db\n")
    }
}

func main() {
    if len(os.Args) == 1 { 
        fmt.Print("no args supplied")
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
    var arg5 string
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
            arg4 = os.Args[4]
            arg5 = os.Args[5]
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
            if(arg2 == "update"){
                fmt.Print("Mode selected: manage update\n")
                UpdateDB(arg3)
            }
            // handles "manage list"
            if(arg2 == "list"){
                fmt.Print("Mode selected: manage list\n")
            }

        case arg1 == "scan": //scan handler
            fmt.Print("\nMode selected: scan")
        case arg1 == "diff":
            fmt.Print("\nMode selected: diff")
    }
}
