
package main

import "database/sql"

var (
	name ="root"
	pword =""
	host ="localhost"
	port ="3306"
	dbname ="test-user"
)

var DB *sql.DB

type user struct {
	id int
	pwd int
	money int
}

func main()  {

}

func IntoDB(){
	dsn :=name +":"+pword+"@tcp("+host+":"+port+")/"+dbname
	db , err:=sql.Open("mysql",dsn)
	if err !=nil{
		panic(err)
	}
	DB = db

}

func LoginIn(){

}
func Regist ()  {

}