package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"sqlboilerpresentation/models"
	"sqlboilerpresentation/util"
)

func main() {
	ctx := boil.WithDebug(context.Background(), true)

	conn := util.DBConnString()
	log.Printf(conn)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	user, err := models.Users(
		qm.Where(models.UserColumns.Username+"=?", "john_doe"),
		qm.Load(qm.Rels(models.UserRels.Posts)),
	).One(ctx, db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("========================================\n")
	fmt.Printf("User ID: %s\n", user.ID)
	fmt.Printf("Username: %s\n", user.Username)
	fmt.Printf("Email: %s\n", user.Email)

	for _, post := range user.R.Posts {
		fmt.Printf("\tPost ID: %s\n", post.ID)
		fmt.Printf("\tTitle: %s\n", post.Title)
		fmt.Printf("\tContent: %s\n", post.Content.String)
	}
}
