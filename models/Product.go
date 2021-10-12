package models

type Product struct {
	Name            string  `json:"name,omitempty"`
	Category        string  `json:"category,omitempty"`
	Unitprice       float64 `json:"unitprice,omitempty"`
	UnitWeight      string  `json:"unitweight,omitempty"`
	PocketType      string  `json:"pockettype,omitempty"`
	SeasonalProduct bool    `json:"seasonalproduct,omitempty"`
	Description     string  `json:"description,omitempty"`
	Stock           bool    `json:"stock,omitempty"`
	Quantity        int     `json:"quantity,omitempty"`
}
