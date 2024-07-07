package cmd

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "DEL",
	Short: "Delete a key-value pair from the store",
	Long:  `Delete a key-value pair from the store`,
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

		// do a delete
		_, err = db.Query("UPDATE STORE SET expires_at = $1 WHERE key = $2", "1970-01-01T00:00:00Z", key)

		if err != nil {
			log.Fatal(err)
			return err
		}

		return nil
	},
}
