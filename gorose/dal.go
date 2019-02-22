package main

import (
    "github.com/gohouse/gorose"
    _ "github.com/gohouse/gorose/driver/mysql"
    "fmt"
    "os"
)

type user struct {
    Id         string `orm:"id"`
    Username   string `orm:"username"`
    CreateTime  string `orm:"create_time"`
}

// DB Config.(Recommend to use configuration file to import)
var dbConfig = &gorose.DbConfigSingle{
    Driver:          "mysql", // driver: mysql/sqlite/oracle/mssql/postgres
    EnableQueryLog:  true,    // if enable sql logs
    SetMaxOpenConns: 0,       // connection pool of max Open connections, default zero
    SetMaxIdleConns: 0,       // connection pool of max sleep connections
    Prefix:          "",      // prefix of table
    Dsn:             "root:123456@tcp(127.0.0.1:3306)/sql_test?charset=utf8", // db dsn
}

func main() {
    connection, err := gorose.Open(dbConfig)
    if err != nil {
        fmt.Println(err)
        return
    }
    db := connection.NewSession()
    var users []user
    var order_by = os.Args[1]
    //db.Table(&users).Group(order_by).Get()
    db.Table(&users).Order(order_by).Get()
    fmt.Println(db.LastSql)
    fmt.Println(users)
}
