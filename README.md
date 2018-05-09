# tt-dev-test

tt-dev-test implemented a basic RESTful HTTP service in Go for a simplified Tantan backend: Adding users and swiping other people in order to find a match. 

## env && dependency

Golang Vsersion: 1.10.2

Beego Version : 1.9.2 (beego is an open-source, high-performance web framework for the Go programming language. )

Beego Install: go get github.com/astaxie/beego

## DB

DB  use PostgreSQL. The table schema is in db_tables folder. The db connection configuration is in conf/db.conf.

## Run

 If the environment is configured successfully,Runing  `go build`  command to generate binary file.  Then executing binary file to start service.

