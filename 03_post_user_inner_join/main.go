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

	post, err := models.Posts(
		qm.InnerJoin("users on users.id = posts.user_id"),
		models.UserWhere.Username.EQ("john_doe"),
	).One(ctx, db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("========================================\n")
	fmt.Printf("Post Relations: %+v\n", post.R)
	fmt.Printf("\tPost ID: %s\n", post.ID)
	fmt.Printf("\tTitle: %s\n", post.Title)
	fmt.Printf("\tContent: %s\n", post.Content.String)
	fmt.Printf("========================================\n")

	type PostWithUser struct {
		models.Post `boil:"posts,bind"`
		models.User `boil:"users,bind"`
	}

	var userWithPost PostWithUser

	err = models.NewQuery(
		qm.Select("users.*", "posts.*"),
		qm.From(models.TableNames.Posts),
		qm.InnerJoin("users on users.id = posts.user_id"),
		models.UserWhere.Username.EQ("john_doe"),
	).Bind(ctx, db, &userWithPost)

	fmt.Printf("========================================\n")
	fmt.Printf("User ID: %s\n", userWithPost.User.ID)
	fmt.Printf("Username: %s\n", userWithPost.User.Username)
	fmt.Printf("Email: %s\n", userWithPost.User.Email)
	fmt.Printf("\tPost ID: %s\n", userWithPost.Post.ID)
	fmt.Printf("\tTitle: %s\n", userWithPost.Post.Title)
	fmt.Printf("\tContent: %s\n", userWithPost.Post.Content.String)
	fmt.Printf("========================================\n")
}
