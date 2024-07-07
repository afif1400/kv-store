package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "GET",
	Short: "Retrieve a key-value pair from the store",
	Long:  `Retrieve a key-value pair from the store`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		key := args[0]

		connStr := os.Getenv("DATABASE_URL")
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)
			return err
		}

		defer db.Close()

		rows, err := db.Query("SELECT value FROM STORE WHERE key = $1 AND expires_at > $2", key, time.Now().Local())
		if err != nil {
			log.Fatal(err)
			return err
		}

		defer rows.Close()

		if !rows.Next() {
			fmt.Println("Key not found")
			return nil
		}

		for rows.Next() {
			var value string
			if err := rows.Scan(&value); err != nil {
				log.Fatal(err)
				return err
			}
			fmt.Println(value)
		}

		return nil

	},
}
