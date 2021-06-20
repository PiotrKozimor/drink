package cmd

import (
	"log"
	"time"

	"github.com/PiotrKozimor/drink"
	"github.com/spf13/cobra"
)

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Get todays drinks",
	Run: func(cmd *cobra.Command, args []string) {
		s, err := drink.NewStore(dbPath)
		if err != nil {
			log.Fatal(err)
		}
		defer s.DB.Close()
		drinks, err := s.GetTodaysDrinks(time.Now())
		if err != nil {
			log.Fatal(err)
		}
		drink.Print(drinks)
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)
}
