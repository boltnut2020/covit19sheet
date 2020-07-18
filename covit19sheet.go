package main

import(
    "fmt"
    "time"
    "net/http"
    "io/ioutil"
    "encoding/json"
	"sort"
)

type Country struct {
	Country string `json:"Country"`
	Slug string `json:"Slug"`
	CountryCode string `json:"CountryCode"`
	NewConfirmed int `json:"NewConfirmed"`
	NewDeaths int `json:"NewDeaths"`
	Date string `json:"Date"`
}

type Countries struct {
	Countries []Country `json:"Countries"`
}

var apiUrl string = "https://api.covid19api.com/summary"
var asia = []string{"JP", "ID", "IN", "BD", "NP", "CN", "KR", "TH", "PH", "VN", "LA", "MM", "TW", "HK", "GE", "RU", "KH", "MY", "SG"}

func main() {

    var countries Countries
    resp, err := http.Get(apiUrl)

    if err != nil {
		fmt.Println(err)
		return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
		fmt.Println(err)
		return
    }


    err = json.Unmarshal(body, &countries)
    if err != nil {
		fmt.Println(err)
		return
    }

    sort.SliceStable(countries.Countries, func(i, j int) bool {
        return countries.Countries[i].NewConfirmed > countries.Countries[j].NewConfirmed
    })

	// output buffer
	t, _ := time.Parse(time.RFC3339, countries.Countries[0].Date)
	loc, err := time.LoadLocation("Asia/Tokyo")
	date := t.In(loc)

	fmt.Println("Asia + Russia & Georgea NewConfirmed Ranking", date)
	fmt.Printf("|%-28v|%-12v|%-12v|\n", "Country", "NewConfirmed", "NewDeaths")
	for _, c := range countries.Countries {
		if contains(asia, c.CountryCode) {
			fmt.Printf("|%-28v|%-12v|%-12v|\n", c.Country, c.NewConfirmed, c.NewDeaths)
		}
	}

	fmt.Println("\n")
	fmt.Println("World NewConfirmed Ranking", date)
	fmt.Printf("|%-28v|%-12v|%-12v|\n", "Country", "NewConfirmed", "NewDeaths")

	n := 20
	for i :=0; i < n ; i++ {
		fmt.Printf("|%-28v|%-12v|%-12v|\n", countries.Countries[i].Country, countries.Countries[i].NewConfirmed, countries.Countries[i].NewDeaths)
	}

	fmt.Println("Data from covid19api.com")
}

func contains(arr []string, str string) bool {
	for _, a := range arr{
		if a == str {
			return true
		}
	}
	return false
}
