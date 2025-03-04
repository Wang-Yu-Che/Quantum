package main

import (
	"fmt"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

const MySQLDSN = "root:459440374a@(localhost:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	// 连接数据库
	db, err := gorm.Open(mysql.Open(MySQLDSN))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}

	// 生成实例
	g := gen.NewGenerator(gen.Config{
		OutPath:          "./douyin_model/douyin_query/",
		Mode:             gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
		FieldNullable:    true,
		FieldCoverable:   false,
		FieldSignable:    false,
		FieldWithTypeTag: true,
	})

	// 设置目标 db
	g.UseDB(db)

	// 自定义字段标签
	jsonField := gen.FieldJSONTagWithNS(func(columnName string) string {
		toStringField := `balance,`
		if strings.Contains(toStringField, columnName) {
			return columnName + ",string"
		}
		return columnName
	})
	softDeleteField := gen.FieldType("delete_time", "gorm.DeletedAt")
	fieldOpts := []gen.ModelOpt{jsonField, softDeleteField}

	// 生成所有表的结构体
	allModel := g.GenerateAllTable(fieldOpts...)

	// 生成查询方法
	g.ApplyBasic(allModel...)

	// 执行生成
	g.Execute()
}
