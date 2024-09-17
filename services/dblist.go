package services

import (
	"chinmay-sawant/Spider-Query/internal"
	"chinmay-sawant/Spider-Query/model"
	"log"

	"gorm.io/gorm"
)

type Dblist struct {
	db *gorm.DB
}

func (dl *Dblist) InitService(database *gorm.DB) {
	dl.db = database
	dl.db.AutoMigrate(&model.Dbs{})
}

func (dl *Dblist) DefaultRow() {
	dbs := &model.Dbs{Host: "localhost", User: "postgres", Password: "abcd@1234", Dbname: "postgres", Port: "5432"}
	dl.db.Create(&dbs)
}

// for getting the dblist from the database i.e only postgres db
func (dl *Dblist) GetAllDbList() []model.Dbs {
	var dblist []model.Dbs
	var tDblist []model.Dbs
	dl.db.Find(&dblist)
	for _, v := range dblist {
		tDb := internal.InitDB(v)
		dl.InitService(tDb)
		tResults := dl.DynamicQuery(`
		SELECT
			r.rolname AS role_name,
			db.datname AS database_name
		FROM
			pg_roles r
		JOIN
			pg_database db ON db.datdba = r.oid
		WHERE
			r.rolcanlogin = true;
		`)
		for _, row := range tResults {
			v.Dbname = row["database_name"].(string)
			v.User = row["role_name"].(string)
			v.Password = row["role_name"].(string)
			tDblist = append(tDblist, v)
		}
	}
	return tDblist
}

func (dl *Dblist) DynamicQuery(query string) []map[string]interface{} {
	// Execute the query and retrieve results as a raw map slice
	var results []map[string]interface{}
	if err := dl.db.Raw(query).Scan(&results).Error; err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}

	// Process the results dynamically
	/*
		for _, row := range results {
			fmt.Println("Row:")
			for key, value := range row {
				fmt.Printf("%s: %v\n", key, value)
			}
		}
	*/
	return results
}
