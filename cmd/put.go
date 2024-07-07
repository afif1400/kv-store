package cmd

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

var putCmd = &cobra.Command{
	Use:   "PUT",
	Short: "Insert a key-value pair into the store",
	Long:  `Insert a key-value pair into the store`,
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {

		key := args[0]
		value := args[1]
		seconds, err := strconv.Atoi(args[2])
		if err != nil {
			log.Fatal(err)
			return err
		}
		expires_at, err := time.Now().Local().Add(time.Second * time.Duration(seconds)).MarshalText()
		if err != nil {
			log.Fatal(err)
			return err
		}

		connStr := os.Getenv("DATABASE_URL")
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer db.Close()

		// do a upsert
		_, err = db.Query("INSERT INTO STORE (key, value, expires_at) VALUES ($1, $2, $3) ON CONFLICT (key) DO UPDATE SET value = $2, expires_at = $3", key, value, expires_at)

		if err != nil {
			log.Fatal(err)
			return err
		}

		return nil
	},
}
