package cmd

import (
	"log"

	"github.com/PiotrKozimor/drink"
	"github.com/spf13/cobra"
)

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Get all drinks",
	Run: func(cmd *cobra.Command, args []string) {
		s, err := drink.NewStore(dbPath)
		if err != nil {
			log.Fatal(err)
		}
		defer s.DB.Close()
		drinks, err := s.GetAllDrinks()
		if err != nil {
			log.Fatal(err)
		}
		drink.PrintAll(drinks)
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
}
