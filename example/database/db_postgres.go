package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	DB_USER   = "postgres"
	DB_PASSWD = "1234"
	DB_NAME   = "postgres"
	DB_HOST   = "localhost"
	DB_PORT   = "5432"
)

type st struct {
	Name string
	Age  int
}

// 에러처리는 제외하고 예제를 위해 작성
func main() {
	datasource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWD, DB_NAME)
	database, _ := sql.Open("postgres", datasource)
	defer database.Close()

	// INSERT, UPDATE, DELETE의 경우 Exec 사용
	database.Exec("INSERT INTO st VALUES($1, $2)", "name", 99)

	// PreparedStatement
	stmt, _ := database.Prepare("SELECT * FROM go_pg WHERE name=$1")
	u := st{}

	// QueryRow 한줄을 반환한다.
	_ = stmt.QueryRow("tester").Scan(&u.Name, &u.Age)

	fmt.Println(u)

	// Query는 rows를 반환.
	result, _ := stmt.Query("tester")
	defer result.Close()

	var name string
	var age int
	for result.Next() {
		_ = result.Scan(&name, &age)
		fmt.Println(name, age)
	}

}
