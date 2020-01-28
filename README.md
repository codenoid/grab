# grab [WIP]
Collection of numerous GRAB's API in Go

## Installation

`go get -u github.com/codenoid/grab`

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
	resNearMe, err := food.GetRestaurant("-6.1328296", "106.8140347", "geprek bensu", 32, 32)
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

## Legal

This code is in no way affiliated with, authorized, maintained, sponsored or endorsed by Grab Holdings Inc or any of its affiliates or subsidiaries. This is an independent and unofficial software. Use at your own risk.

## License

The MIT License (MIT)

Copyright (c) 2018

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
