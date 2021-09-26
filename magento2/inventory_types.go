package m2go

type InventorySource struct {
	Items          []*Source `json:"items"`
	SearchCriteria struct {
		FilterGroups []interface{} `json:"filter_groups"`
	} `json:"search_criteria"`
	TotalCount int `json:"total_count"`
}

type Source struct {
	SourceCode              string        `json:"source_code"`
	Name                    string        `json:"name"`
	Enabled                 bool          `json:"enabled"`
	CountryID               string        `json:"country_id"`
	Region                  string        `json:"region,omitempty"`
	City                    string        `json:"city,omitempty"`
	Street                  string        `json:"street,omitempty"`
	Postcode                string        `json:"postcode"`
	Phone                   string        `json:"phone,omitempty"`
	UseDefaultCarrierConfig bool          `json:"use_default_carrier_config"`
	CarrierLinks            []interface{} `json:"carrier_links"`
	ExtensionAttributes     struct {
		IsPickupLocationActive bool   `json:"is_pickup_location_active"`
		FrontendName           string `json:"frontend_name"`
	} `json:"extension_attributes"`
	Description string `json:"description,omitempty"`
	Latitude    float64    `json:"latitude,omitempty"`
	Longitude   float64    `json:"longitude,omitempty"`
}

type SourceItem struct {
	Sku        string `json:"sku"`
	SourceCode string `json:"source_code"`
	Quantity   int    `json:"quantity"`
	Status     int8    `json:"status"`
}
