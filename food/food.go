package food

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetRestaurant(lat, lon, keyword string, offset, pageSize int) (FoodResult, error) {
	client := &http.Client{}

	var query = fmt.Sprintf(`{"latlng":"%v,%v","keyword": "%v","offset": %v,"pageSize": %v}`, lat, lon, keyword, offset, pageSize)
	var payload = strings.NewReader(query)
	req, err := http.NewRequest("POST", "https://portal.grab.com/foodweb/v2/search", payload)
	if err != nil {
		log.Fatal(err)
		return FoodResult{}, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Windows x86_64; rv:72.0) Gecko/20100101 Firefox/75.0")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en")
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Set("Origin", "https://food.grab.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "https://food.grab.com/id/en/restaurants")
	req.Header.Set("TE", "Trailers")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return FoodResult{}, err
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return FoodResult{}, err
	}

	data, err := UnmarshalFood(bodyText)
	if err != nil {
		return FoodResult{}, err
	}
	return data, nil
}
