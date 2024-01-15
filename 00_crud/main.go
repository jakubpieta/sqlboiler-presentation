package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"sqlboilerpresentation/models"
	"sqlboilerpresentation/util"
)

func main() {
	ctx := boil.WithDebug(context.Background(), true)

	conn := util.DBConnString()
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

	user := &models.User{
		Username: "example_username",
		Email:    "example@example.com",
	}

	if err := user.Insert(ctx, db, boil.Infer()); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("========================================\n")
	fmt.Printf("User ID: %s\n", user.ID)
	fmt.Printf("Username: %s\n", user.Username)
	fmt.Printf("Email: %s\n", user.Email)
	fmt.Printf("========================================\n")

	user.Username = "new_example_username"
	updatedCount, err := user.Update(ctx, db, boil.Infer())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("========================================\n")
	fmt.Printf("Updated rows count: %d\n", updatedCount)
	fmt.Printf("User ID: %s\n", user.ID)
	fmt.Printf("Username: %s\n", user.Username)
	fmt.Printf("Email: %s\n", user.Email)
	fmt.Printf("========================================\n")

	// similar to Users.All()
	user, err = models.Users(
		models.UserWhere.Username.EQ("new_example_username"),
	).One(ctx, db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("========================================\n")
	fmt.Printf("User ID: %s\n", user.ID)
	fmt.Printf("Username: %s\n", user.Username)
	fmt.Printf("Email: %s\n", user.Email)
	fmt.Printf("========================================\n")

	// similar to Users.DeleteAll(), but we need to pass a WHERE query to arguments
	rowsDeleted, err := user.Delete(ctx, db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("========================================\n")
	fmt.Printf("Deleted rows count: %d\n", rowsDeleted)
	fmt.Printf("========================================\n")
}
