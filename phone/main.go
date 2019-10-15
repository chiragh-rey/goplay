package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"goplay/phone/db"
	"regexp"
)

const (
	drivername = "mysql"
	host     = "localhost"
	port     = 3306
	user     = "pikachu"
	password = "password"
	dbname   = "phone"
)

func main() {
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/", user, password, host, port)
	must(db.Reset(drivername, mysqlInfo, dbname))

	mysqlInfo = fmt.Sprintf("%s%s", mysqlInfo, dbname)
	must(db.Migrate(drivername, mysqlInfo))

	dbA, err := db.Open(drivername, mysqlInfo)
	must(err)
	defer dbA.Close()

	err = dbA.Seed()
	must(err)

	phones, err := dbA.AllPhones()
	must(err)
	for _, p := range phones {
		fmt.Printf("Working on... %+v\n", p)
		number := normalize(p.Number)
		if number != p.Number {
			fmt.Println("Updating or removing...", number)
			existing, err := dbA.FindPhone(number)
			must(err)
			if existing != nil {
				must(dbA.DeletePhone(p.ID))
			} else {
				p.Number = number
				must(dbA.UpdatePhone(&p))
			}
		} else {
			fmt.Println("No changes required")
		}
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func normalize(phone string) string {
	re := regexp.MustCompile("\\D")
	return re.ReplaceAllString(phone, "")
}

// func normalize(phone string) string {
// 	var buf bytes.Buffer
// 	for _, ch := range phone {
// 		if ch >= '0' && ch <= '9' {
// 			buf.WriteRune(ch)
// 		}
// 	}
// 	return buf.String()
// }