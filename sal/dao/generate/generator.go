package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"miner_core/sal/dao/model"
)

func main() {
	//modelGenerator()
	queryGenerator()
}

func modelGenerator() {
	g := gen.NewGenerator(gen.Config{
		FieldNullable: true,
		OutPath:       "./sal/dao/model",
		Mode:          gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	db, _ := gorm.Open(mysql.Open("root:wxl5211314@tcp(wxl475.cn:32130)/miner?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	g.UseDB(db)

	g.GenerateModelAs("jobs", "JobPO", gen.FieldGenType("deleted_at", "gorm.DeletedAt"))

	g.Execute()
}

func queryGenerator() {
	g := gen.NewGenerator(gen.Config{
		FieldNullable: true,
		OutPath:       "./sal/dao/query",
		Mode:          gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	db, _ := gorm.Open(mysql.Open("root:wxl5211314@tcp(wxl475.cn:32130)/miner?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	g.UseDB(db)

	g.ApplyBasic(model.JobPO{})

	g.Execute()
}
