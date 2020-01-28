# grab [WIP]
Collection of numerous GRAB's API in Go

## Food Data

[explore here](https://github.com/codenoid/grab/blob/master/food/structs.go) for more data

```
package main

import (
	"fmt"
	"math"

	"github.com/codenoid/grab/food"
)

func main() {
	resNearMe, err := food.GetRestaurant("-6.1328296", "106.8140347", "geprek bensu")
	if err != nil {
		fmt.Println(err)
	}

	for index, restauran := range resNearMe.SearchResult.SearchMerchants {
		index += 1
		fmt.Printf("%v. %v berjarak %vkm dari posisi anda \n", index, restauran.Address.Name, math.Round(restauran.MerchantBrief.DistanceInKM))
	}
	// 1. Ayam Geprek Sambel Rondo - Maphar berjarak 2km dari posisi anda
	// 2. Ayam Geprek Bagors - Cempaka Putih Barat berjarak 7km dari posisi anda
	// 3. Ayam Geprek Koboy - Kota Bambu Selatan berjarak 6km dari posisi anda
	// 4. Geprek Sumo - Ancol GrabKitchen berjarak 1km dari posisi anda
	// 5. Ayam Goreng Bakar Geprek Gendeng - Hadiah berjarak 5km dari posisi anda
}
```