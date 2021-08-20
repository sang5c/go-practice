package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string `gorm:"default:999"`
	Price uint
}

func main() {
	dsn := "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Seoul"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Db 연결에 실패하였습니다.")
	}

	// 테이블 자동 생성
	db.AutoMigrate(&Product{})

	// 생성
	value := &Product{Price: 100}
	fmt.Println(value)
	db.Create(value)

	//value := &Product{}
	//db.First(&value,"id = ?", "16")
	//value.Code = "1234576"
	//db.Save(value)

	//// 수정 - product의 price를 200으로
	//db.Model(&product).Update("Price", 200)
	//// 수정 - 여러개의 필드를 수정하기
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//fmt.Println(product)

	// 삭제 - product 삭제하기
	//db.Delete(&product, 1)
}
