package cities

import (
	"math/rand/v2"
)

var cities = [...]string{
	"New York", "Tokyo", "London", "Paris", "Berlin",
	"Sydney", "Toronto", "Dubai", "Singapore", "Barcelona",
	"Moscow", "Los Angeles", "Rio de Janeiro", "Istanbul", "Seoul",
	"Mexico City", "Bangkok", "Buenos Aires", "Cairo", "Rome",
	"Amsterdam", "Hong Kong", "Madrid", "Lima", "Lisbon",
	"San Francisco", "Chicago", "Kuala Lumpur", "Mumbai", "Jakarta",
	"Vienna", "Brussels", "Athens", "Helsinki", "Oslo",
	"Stockholm", "Copenhagen", "Zurich", "Dublin", "Budapest",
	"Prague", "Warsaw", "Santiago", "Nairobi", "Manila",
	"Lisbon", "Bangalore", "Cape Town", "Tel Aviv", "Abu Dhabi",
	"Doha", "Hanoi", "Colombo", "Tunis", "Lagos",
	"Accra", "Addis Ababa", "Algiers", "Helsinki", "Havana",
	"Valencia", "Belfast", "Edinburgh", "Glasgow", "Montreal",
	"Brisbane", "Perth", "Adelaide", "Wellington", "Christchurch",
	"Osaka", "Nagoya", "Kobe", "Fukuoka", "Kyoto",
	"Seville", "Malaga", "Bilbao", "Granada", "Valencia",
	"Antwerp", "Ghent", "Bruges", "Rotterdam", "Utrecht",
	"Geneva", "Lausanne", "Lucerne", "Stuttgart", "Frankfurt",
	"Munich", "Düsseldorf", "Cologne", "Hamburg", "Leipzig",
}

func GenerateRandomCities() []string {
	citiesCopy := make([]string, len(cities))
	copy(citiesCopy, cities[:])
	rand.Shuffle(len(citiesCopy), func(i, j int) {
		citiesCopy[i], citiesCopy[j] = citiesCopy[j], citiesCopy[i]
	})
	return citiesCopy
}
