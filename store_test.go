package drink

import (
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/matryer/is"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func mustNewTestStore(is *is.I) *Store {
	randPostfix := strconv.Itoa(rand.Int())
	s, err := NewStore(os.TempDir() + "/bolt-" + randPostfix)
	is.NoErr(err)
	return s
}

var (
	timeEarly  = time.Date(2021, 1, 12, 0, 0, 1, 0, time.UTC)
	timeLate   = time.Date(2021, 1, 12, 23, 59, 59, 0, time.UTC)
	testDrinks = []Drink{
		{
			HowMuch: 240,
			When:    time.Date(2021, 1, 11, 23, 59, 59, 0, time.UTC),
		},
		{
			HowMuch: 240,
			When:    timeEarly,
		},
		{
			HowMuch: 240,
			When:    timeLate,
		},
		{
			HowMuch: 240,
			When:    time.Date(2021, 1, 13, 0, 0, 1, 0, time.UTC),
		},
	}
)

func mustDrinkMany(s *Store, is *is.I) {
	for _, drink := range testDrinks {
		err := s.DrinkIt(&drink)
		is.NoErr(err)
	}
}

func TestDrinkItOne(t *testing.T) {
	testDrink := Drink{
		HowMuch: 240,
		When:    time.Date(2021, 1, 1, 12, 1, 1, 0, time.UTC),
	}
	is := is.New(t)
	s := mustNewTestStore(is)
	err := s.DrinkIt(&testDrink)
	is.NoErr(err)
	drinks, err := s.GetAllDrinks()
	is.NoErr(err)
	is.Equal(len(drinks), 1)
	testDay := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	is.Equal(drinks[testDay][0].When, testDrink.When)
	is.Equal(drinks[testDay][0].HowMuch, testDrink.HowMuch)
}

func TestDrinkItMany(t *testing.T) {
	is := is.New(t)
	s := mustNewTestStore(is)
	mustDrinkMany(s, is)
	drinks, err := s.GetAllDrinks()
	is.NoErr(err)
	is.Equal(len(drinks), len(testDrinks)-1) // we have two drinks for one day
	is.Equal(drinks[time.Date(2021, 1, 11, 0, 0, 0, 0, time.UTC)][0], testDrinks[0])
	is.Equal(drinks[time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC)][0], testDrinks[1])
	is.Equal(drinks[time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC)][1], testDrinks[2])
	is.Equal(drinks[time.Date(2021, 1, 13, 0, 0, 0, 0, time.UTC)][0], testDrinks[3])
}

func TestGetTodaysDrinks(t *testing.T) {
	is := is.New(t)
	s := mustNewTestStore(is)
	mustDrinkMany(s, is)
	drinks, err := s.GetTodaysDrinks(timeEarly)
	is.NoErr(err)
	is.Equal(len(drinks), 2)
	is.Equal(drinks[0], testDrinks[1])
	is.Equal(drinks[1], testDrinks[2])
}

func TestMarshall(t *testing.T) {
	testDrink := Drink{
		HowMuch: 240,
		When:    time.Date(2021, 1, 1, 12, 1, 1, 0, time.UTC),
	}
	k, v := marshall(&testDrink)
	is := is.New(t)
	is.Equal(v, []byte{240, 0, 0, 0})
	is.Equal(k, []byte("2021-01-01T12:01:01Z"))
}
