package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

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

	users, err := models.Users().All(ctx, db)
	if err != nil {
		return
	}

	for _, user := range users {
		fmt.Printf("========================================\n")
		fmt.Printf("User ID: %s\n", user.ID)
		fmt.Printf("Username: %s\n", user.Username)
		fmt.Printf("Email: %s\n", user.Email)
	}

	complexQueryUsers, err := models.Users(
		models.UserWhere.Email.EQ("example@example.com"),
		models.UserWhere.CreatedAt.GT(null.TimeFrom(time.Now())),
		models.UserWhere.Username.LIKE("example%"),
		models.UserWhere.UpdatedAt.LT(null.TimeFrom(time.Now())),
		models.UserWhere.ID.IN([]string{"1", "2", "3"}),
	).All(ctx, db)
	if err != nil {
		return
	}

	for _, user := range complexQueryUsers {
		fmt.Printf("========================================\n")
		fmt.Printf("User ID: %s\n", user.ID)
		fmt.Printf("Username: %s\n", user.Username)
		fmt.Printf("Email: %s\n", user.Email)
	}
}
