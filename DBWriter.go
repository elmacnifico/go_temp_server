package main

import (
    "database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
    "time"
)

type DBWriter struct {
	Db        *sql.DB
	cache     *Cache
}

func (self *DBWriter) init( cache *Cache ) {
	var err error
    self.Db, err = sql.Open("sqlite3", "/tmp/sqlite")
    if err != nil {
        log.Println( err )
    }
    self.cache = cache
    self.initScheme()
}

func (self *DBWriter) processInput() {
    for {
        now := time.Now().Minute()
        for k,v := range self.cache.Content {
            for i := 0; i<now; i++ {
                if v.Minutes[i].written == false {
                    //write MinuteCache v to db
                    v.Minutes[i].written = true
                }
            }
        }
        time.Sleep( time.Second )
    }
}

func (self *DBWriter) initScheme() {
    sqls := []string{
        "create table hosts (host_id integer not null primary key, host_name text)",
        "create table checks (check_id integer not null primary key, check_name text, check_host integer,FOREIGN KEY(check_host) REFERENCES hosts(host_id))",
        "create table data (data_id integer not null primary key, data_name text, data_check_id integer,    FOREIGN KEY(data_check_id) REFERENCES checks(check_id))",
    }
    for _, sql := range sqls {
        _, err := self.Db.Exec(sql)
        if err != nil {
            log.Println(err)
        }
    }
}

func (self *DBWriter) insertDataPoint() {
    tx, err := db.Begin()
    if err != nil {
        log.Println(err)
    }

    stmt, err := tx.Prepare( "insert into test(id, text) values(?,?)")
    if err != nil {
        log.Println(err)
    }
    defer stmt.Close()

    for i:=0; i<100000; i++ {
        _, err = stmt.Exec(i, "blah")
        if err != nil {
            log.Println(err)
        }
    }
    tx.Commit()
}
