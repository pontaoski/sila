package app

import (
	"database/sql"
	"fmt"
	"math"
	"os"
	"strconv"

	"context"

	"github.com/alecthomas/repr"
	_ "github.com/lib/pq"
	silaDB "github.com/pontaoski/sila/app/db"
	"github.com/urfave/cli/v2"
)

const (
	HOST   = "localhost"
	PORT   = 5432
	USER   = "postgres"
	DBNAME = "sila"
)

func Complain(out string) {
	println(out)
	os.Exit(1)
}

func ParseInt(in string) int64 {
	out, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		Complain(err.Error())
	}
	return out
}

func Main() {
	info := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, DBNAME,
	)
	conn, err := sql.Open("postgres", info)
	if err != nil {
		repr.Println(err)
		return
	}
	defer conn.Close()

	err = conn.Ping()
	if err != nil {
		repr.Println(err)
		return
	}

	ctx := context.Background()
	db := silaDB.New(conn)

	str := func(in string) sql.NullString {
		return sql.NullString{String: in, Valid: true}
	}

	sbool := func(in bool) sql.NullBool {
		return sql.NullBool{Bool: in, Valid: true}
	}

	silaDB.Prepare(ctx, conn)

	app := cli.App{
		Commands: []*cli.Command{
			{
				Name: "users",
				Subcommands: []*cli.Command{
					{
						Name: "add",
						Action: func(c *cli.Context) error {
							if c.NArg() < 2 {
								Complain("Please provide an email and a full name")
							}
							user, err := db.CreateUser(ctx, silaDB.CreateUserParams{
								Email:    str(c.Args().Get(0)),
								Fullname: str(c.Args().Get(1)),
							})
							if err != nil {
								Complain(err.Error())
							}
							fmt.Printf("User %d created:\n\tReal Name: %s\n\tEmail: %s\n", user.ID, user.Fullname.String, user.Email.String)
							return nil
						},
					},
					{
						Name: "list",
						Action: func(c *cli.Context) error {
							users, err := db.ListUsers(ctx, silaDB.ListUsersParams{
								Limit:  math.MaxInt32,
								Offset: 0,
							})
							if err != nil {
								Complain(err.Error())
							}
							for _, user := range users {
								fmt.Printf("User %d:\n\tReal Name: %s\n\tEmail: %s\n", user.ID, user.Fullname.String, user.Email.String)
							}
							return nil
						},
					},
				},
			},
			{
				Name: "products",
				Subcommands: []*cli.Command{
					{
						Name: "add",
						Action: func(c *cli.Context) error {
							product, err := db.CreateProduct(ctx, silaDB.CreateProductParams{
								Name:               str(c.Args().Get(0)),
								ProductDescription: str(c.Args().Get(1)),
								Active:             sbool(true),
							})
							if err != nil {
								Complain(err.Error())
							}
							fmt.Printf("Product %d created:\n\tName: %s\n\tDescription: %s\n", product.ID, product.Name.String, product.ProductDescription.String)
							return nil
						},
					},
					{
						Name: "list",
						Action: func(c *cli.Context) error {
							products, err := db.ListProducts(ctx, silaDB.ListProductsParams{
								Limit:  math.MaxInt32,
								Offset: 0,
							})
							if err != nil {
								Complain(err.Error())
							}
							for _, product := range products {
								fmt.Printf("Product %n (%d):\t%s", product.Name.String, product.ID, product.ProductDescription.String)
							}
							return nil
						},
					},
				},
			},
		},
	}
	app.Run(os.Args)
}
