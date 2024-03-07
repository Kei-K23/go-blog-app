package lib

import (
	"database/sql"
	"fmt"
	"log"
)

type DBLib struct {}

func (d *DBLib) CreateBlogTable(db *sql.DB)  {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS blogs (
		id VARCHAR(255) PRIMARY KEY NOT NULL,
		title VARCHAR(255) NOT NULL,
		description TEXT NOT NULL,
		body TEXT NOT NULL,
		is_public BOOLEAN DEFAULT false,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Fatal("Error creating blogs table:", err)
	}
	fmt.Println("Blogs table created successfully")
}