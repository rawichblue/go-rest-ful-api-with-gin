package cmd

import (
	"app/models"
	"app/modules"
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var transactions = &cobra.Command{
	Use:     "migrate-transactions",
	Aliases: []string{"addition"},
	Short:   "Start server",
	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Migrating transactions")
		db := modules.Get().DB
		if _, err := db.NewCreateTable().Model((*models.Transactions)(nil)).Exec(context.Background()); err != nil {
			log.Panic(err)
			os.Exit(1)
			return
		}
		log.Println("Model transactions up success")
	},
}

func init() {
	rootCmd.AddCommand(transactions)
}
