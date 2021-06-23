
## Parameter PlaceHolder
- MySQL
  * `SELECT * FROM table WHERE name=?`
  * 물음표로 표현 
- PostgreSQL
  * `SELECT * FROM table WHERE name=$1`
  * $1, $2 ...
- Oracle
  * `SELECT * FROM table WHERE name=:name`
  * 콜론 변수명으로 표현한다.

## 참고
- http://go-database-sql.org/prepared.html