package food

import "encoding/json"

func UnmarshalFood(data []byte) (FoodResult, error) {
	var r FoodResult
	err := json.Unmarshal(data, &r)
	return r, err
}

type FoodResult struct {
	SearchResult SearchResult `json:"searchResult"`
}

type SearchResult struct {
	SearchID        string           `json:"searchID"`
	TotalCount      int64            `json:"totalCount"`
	SearchMerchants []SearchMerchant `json:"searchMerchants"`
	RequestID       string           `json:"requestID"`
	Offset          int64            `json:"offset"`
	HasMore         bool             `json:"hasMore"`
}

type SearchMerchant struct {
	ID                   string                      `json:"id"`
	Address              Address                     `json:"address"`
	Latlng               Latlng                      `json:"latlng"`
	MerchantBrief        SearchMerchantMerchantBrief `json:"merchantBrief"`
	ChainID              *string                     `json:"chainID,omitempty"`
	ChainName            *string                     `json:"chainName,omitempty"`
	EstimatedDeliveryFee EstimatedDeliveryFee        `json:"estimatedDeliveryFee"`
	BranchMerchants      []BranchMerchant            `json:"branchMerchants"`
}

type Address struct {
	Name            string `json:"name"`
	CombinedAddress string `json:"combined_address"`
	CombinedCity    City   `json:"combined_city"`
	City            City   `json:"city"`
}

type BranchMerchant struct {
	ID                   string                      `json:"id"`
	Address              Address                     `json:"address"`
	Latlng               Latlng                      `json:"latlng"`
	MerchantBrief        BranchMerchantMerchantBrief `json:"merchantBrief"`
	ChainID              string                      `json:"chainID"`
	ChainName            string                      `json:"chainName"`
	BranchName           string                      `json:"branchName"`
	EstimatedDeliveryFee *EstimatedDeliveryFee       `json:"estimatedDeliveryFee,omitempty"`
}

type EstimatedDeliveryFee struct {
	Currency     Currency `json:"currency"`
	Price        int64    `json:"price"`
	PriceDisplay string   `json:"priceDisplay"`
	Multiplier   int64    `json:"multiplier"`
}

type Currency struct {
	Code     Code   `json:"code"`
	Symbol   Symbol `json:"symbol"`
	Exponent int64  `json:"exponent"`
}

type Latlng struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type BranchMerchantMerchantBrief struct {
	Cuisine        []string  `json:"cuisine"`
	PhotoHref      string    `json:"photoHref"`
	SmallPhotoHref *string   `json:"smallPhotoHref,omitempty"`
	Halal          bool      `json:"halal"`
	IsIntegrated   *bool     `json:"isIntegrated,omitempty"`
	OpenHours      OpenHours `json:"openHours"`
	DistanceInKM   float64   `json:"distanceInKm"`
	Rating         float64   `json:"rating"`
	VoteCount      int64     `json:"vote_count"`
	Promo          *Promo    `json:"promo,omitempty"`
	DeliverBy      DeliverBy `json:"deliverBy"`
	ClosedText     *string   `json:"closedText,omitempty"`
	ClosedTextType *string   `json:"closedTextType,omitempty"`
	Description    *string   `json:"description,omitempty"`
	IconHref       *string   `json:"iconHref,omitempty"`
}

type OpenHours struct {
	Open           *bool  `json:"open,omitempty"`
	DisplayedHours string `json:"displayedHours"`
	Sun            string `json:"sun"`
	Mon            string `json:"mon"`
	Tue            string `json:"tue"`
	Wed            string `json:"wed"`
	Thu            string `json:"thu"`
	Fri            string `json:"fri"`
	Sat            string `json:"sat"`
	TempClosed     *bool  `json:"tempClosed,omitempty"`
}

type Promo struct {
	HasPromo    bool   `json:"hasPromo"`
	Description string `json:"description"`
}

type SearchMerchantMerchantBrief struct {
	Description    *string   `json:"description,omitempty"`
	Cuisine        []string  `json:"cuisine"`
	PhotoHref      string    `json:"photoHref"`
	IconHref       *string   `json:"iconHref,omitempty"`
	IsIntegrated   *bool     `json:"isIntegrated,omitempty"`
	OpenHours      OpenHours `json:"openHours"`
	DistanceInKM   float64   `json:"distanceInKm"`
	Rating         float64   `json:"rating"`
	VoteCount      int64     `json:"vote_count"`
	DeliverBy      DeliverBy `json:"deliverBy"`
	SmallPhotoHref *string   `json:"smallPhotoHref,omitempty"`
	Halal          *bool     `json:"halal,omitempty"`
	Promo          *Promo    `json:"promo,omitempty"`
}

type City string

const (
	Jakarta City = "Jakarta"
)

type Code string

const (
	Idr Code = "IDR"
)

type Symbol string

const (
	Rp Symbol = "RP"
)

type DeliverBy string

const (
	Grab DeliverBy = "GRAB"
)
