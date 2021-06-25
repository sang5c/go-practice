package main

import "fmt"

type Zip struct {
	Name    string
	address string
}

type SuperZip struct {
	Zipzip Zip
	Level  int
}

type SuperZip2 struct {
	Zip
	Level int
}

func strt() {
	zip := Zip{"dd", "1234"}
	fmt.Println(zip.address, zip.Name)

	value := SuperZip{}
	fmt.Println(value.Zipzip.address, value.Zipzip.Name, value.Level)

	value2 := SuperZip2{Zip{"name", "add"}, 9}
	fmt.Println(value2.address, value2.Name, value2.Zip.Name, value2.Level) // 두 가지 방법으로 모두 접근 가능함

	copyValue := value2
	copyValue.Name = "change"
	fmt.Println(copyValue, value2)
}
