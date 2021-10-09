package drink

import (
	"testing"
	"time"
)

func TestPrintDrinks(t *testing.T) {
	testDrinks := []Drink{
		{
			HowMuch: 240,
			When:    time.Date(2021, 1, 1, 12, 1, 1, 0, time.UTC),
		},
		{
			HowMuch: 300,
			When:    time.Date(2021, 1, 1, 1, 1, 1, 0, time.UTC),
		},
	}
	Print(testDrinks)
}

func TestPrintAllDrinks(t *testing.T) {
	testDrinks := []DailyDrinks{
		{
			Day: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			Drinks: []Drink{
				{
					HowMuch: 240,
					When:    time.Date(2021, 1, 1, 12, 1, 1, 0, time.UTC),
				},
				{
					HowMuch: 300,
					When:    time.Date(2021, 1, 1, 1, 1, 1, 0, time.UTC),
				},
			},
		},
		{
			Day: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
			Drinks: []Drink{
				{
					HowMuch: 240,
					When:    time.Date(2021, 1, 2, 12, 1, 1, 0, time.UTC),
				},
			},
		},
	}
	PrintAll(testDrinks)

}
