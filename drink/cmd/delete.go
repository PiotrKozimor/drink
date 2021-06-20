package cmd

import (
	"log"

	"github.com/PiotrKozimor/drink"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete last drink",
	Run: func(cmd *cobra.Command, args []string) {
		s, err := drink.NewStore(dbPath)
		if err != nil {
			log.Fatal(err)
		}
		defer s.DB.Close()
		err = s.RemoveLastDrink()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
