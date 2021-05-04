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

func UpdateDB(arg1 string) {
    //check if db exists?
    fmt.Print(arg1 + "\n")
    if _, err := os.Stat("./dbs/" + arg1); err == nil {
        fmt.Print("\nits there")
    } else {
        fmt.Print("\nno db")
    }
}

func main() {
    //args handling 
    arg1 := os.Args[1] //arg1 = manage, scan or diff
    arg2 := os.Args[2] //arg2 = (manage: update, list) (scan: db-name) (diff: db-name)
   // arg3 := os.Args[3] //arg3 = (manage: db-name) (scan: all, js, html) (diff: date1)
   // arg4 := os.Args[4] //arg4 = (manage: endpoints.txt) (scan: ) (diff: date2)

    //switch statment to handle args
    switch {
        case arg1 == "manage":
            fmt.Print("\nMode selected: manage")
            if(arg2 == "update"){
                fmt.Print("\nMode selected: update")
            }

        case arg1 == "scan":
            fmt.Print("\nMode selected: scan")
        case arg1 == "diff":
            fmt.Print("\nMode selected: diff")
    }
}
