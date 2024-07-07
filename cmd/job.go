package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

var jobCmd = &cobra.Command{
	Use:   "job",
	Short: "A job is a task that can be run",
	Long:  `A job is a task that can be run`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Background job running...")
		go backgroundJob()

		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

		// Block until a signal is received
		<-sigChan
		log.Print("Shutting down background job...")
	},
}

func backgroundJob() {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	for {
		time.Sleep(5 * time.Second)

		// Start a transaction
		tx, err := db.Begin()
		if err != nil {
			log.Printf("Failed to begin transaction: %v", err)
			continue
		}

		result, err := tx.Exec("DELETE FROM STORE WHERE expires_at <= $1", time.Now().Local())
		if err != nil {
			log.Printf("Failed to delete expired keys: %v", err)
			tx.Rollback()
			continue
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Printf("Failed to get rows affected: %v", err)
			tx.Rollback()
			continue
		}

		log.Printf("Deleted %d expired keys", rowsAffected)

		// Commit the transaction
		if err := tx.Commit(); err != nil {
			log.Printf("Failed to commit transaction: %v", err)
			continue
		}
	}
}
