package m2go

type Shipment struct {
	BillingAddressID    int                `json:"billing_address_id"`
	Comments            []*ShipmentComment `json:"comments"`
	CreatedAt           string             `json:"created_at"`
	EntityID            int                `json:"entity_id"`
	ExtensionAttributes struct {
		SourceCode string `json:"source_code"`
	} `json:"extension_attributes"`
	IncrementID       string          `json:"increment_id"`
	Items             []*ShipmentItem `json:"items"`
	OrderID           int             `json:"order_id"`
	Packages          []interface{}   `json:"packages"`
	ShippingAddressID int             `json:"shipping_address_id"`
	StoreID           int             `json:"store_id"`
	TotalQty          int             `json:"total_qty"`
	Tracks            []*Track        `json:"tracks"`
	UpdatedAt         string          `json:"updated_at"`
}

type ShipmentItem struct {
	EntityID    int     `json:"entity_id"`
	Name        string  `json:"name"`
	OrderItemID int     `json:"order_item_id"`
	ParentID    int     `json:"parent_id"`
	Price       float64 `json:"price"`
	ProductID   int     `json:"product_id"`
	Qty         int     `json:"qty"`
	Sku         string  `json:"sku"`
	Weight      float64 `json:"weight"`
}

type ShipmentComment struct {
	Comment            string `json:"comment"`
	CreatedAt          string `json:"created_at"`
	EntityID           int    `json:"entity_id"`
	IsCustomerNotified int    `json:"is_customer_notified"`
	IsVisibleOnFront   int    `json:"is_visible_on_front"`
	ParentID           int    `json:"parent_id"`
}

type Track struct {
	CarrierCode string  `json:"carrier_code"`
	CreatedAt   string  `json:"created_at"`
	Description string  `json:"description"`
	EntityID    int     `json:"entity_id"`
	OrderID     int     `json:"order_id"`
	ParentID    int     `json:"parent_id"`
	Qty         int     `json:"qty"`
	Title       string  `json:"title"`
	TrackNumber string  `json:"track_number"`
	UpdatedAt   string  `json:"updated_at"`
	Weight      float64 `json:"weight"`
}
