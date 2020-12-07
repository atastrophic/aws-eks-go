package main

import (
	"fmt"

	"github.com/atastrophic/go-ms-with-eks/pkg/application"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				fmt.Println(err.Error())
			} else {
				panic(r)
			}
		}
	}()

	application.NewApplication(
		application.NewAppDep(),
	).
		Start()
}
