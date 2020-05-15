package app

import (
	"database/sql"
	"fmt"
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

	silaDB.Prepare(ctx, conn)

	app := cli.App{
		Commands: []*cli.Command{
			{
				Name: "new",
				Action: func(c *cli.Context) error {
					bug, err := db.CreateBug(ctx, silaDB.CreateBugParams{
						Title:          str(c.Args().Get(0)),
						BugDescription: str(c.Args().Get(1)),
					})
					if err != nil {
						repr.Println(err)
					}
					repr.Println(bug)
					return nil
				},
			},
			{
				Name: "list",
				Action: func(c *cli.Context) error {
					bugs, err := db.ListBugs(ctx, silaDB.ListBugsParams{
						Limit:  0,
						Offset: 50,
					})
					if err != nil {
						repr.Println(err)
					}
					repr.Println(bugs)
					return nil
				},
			},
			{
				Name: "delete",
				Action: func(c *cli.Context) error {
					parsed, err := strconv.ParseInt(c.Args().First(), 10, 64)
					if err != nil {
						repr.Println(err)
						return nil
					}
					err = db.ClearDependenciesForBug(ctx, parsed)
					if err != nil {
						repr.Println(err)
						return nil
					}
					err = db.ClearDependenciesOnBug(ctx, parsed)
					if err != nil {
						repr.Println(err)
						return nil
					}
					err = db.DeleteBug(ctx, parsed)
					if err != nil {
						repr.Println(err)
						return nil
					}
					return nil
				},
			},
			{
				Name: "deps",
				Subcommands: []*cli.Command{
					{
						Name: "of",
						Action: func(c *cli.Context) error {
							parsed, err := strconv.ParseInt(c.Args().First(), 10, 64)
							if err != nil {
								repr.Println(err)
								return nil
							}
							bugs, err := db.Requires(ctx, parsed)
							if err != nil {
								repr.Println(err)
							}
							repr.Println(bugs)
							return nil
						},
					},
					{
						Name: "on",
						Action: func(c *cli.Context) error {
							parsed, err := strconv.ParseInt(c.Args().First(), 10, 64)
							if err != nil {
								repr.Println(err)
								return nil
							}
							bugs, err := db.WhatRequires(ctx, parsed)
							if err != nil {
								repr.Println(err)
							}
							repr.Println(bugs)
							return nil
						},
					},
					{
						Name: "remove",
						Action: func(c *cli.Context) error {
							parsed1, err := strconv.ParseInt(c.Args().Get(0), 10, 64)
							parsed2, err := strconv.ParseInt(c.Args().Get(1), 10, 64)
							if err != nil {
								repr.Println(err)
								return nil
							}
							err = db.RemoveDependency(ctx, silaDB.RemoveDependencyParams{
								RequiredBy: parsed1,
								DependsOn:  parsed2,
							})
							if err != nil {
								repr.Println(err)
							}
							return nil
						},
					},
					{
						Name: "add",
						Action: func(c *cli.Context) error {
							parsed1, err := strconv.ParseInt(c.Args().Get(0), 10, 64)
							parsed2, err := strconv.ParseInt(c.Args().Get(1), 10, 64)
							if err != nil {
								repr.Println(err)
								return nil
							}
							err = db.AddDependency(ctx, silaDB.AddDependencyParams{
								RequiredBy: parsed1,
								DependsOn:  parsed2,
							})
							if err != nil {
								repr.Println(err)
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
