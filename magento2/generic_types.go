package m2go

type CustomAttributes struct {
	AttributeCode string `json:"attribute_code"`
	Value         interface{} `json:"value"`
}

type Payment struct {
	AccountStatus             string   `json:"account_status"`
	AdditionalData            string   `json:"additional_data"`
	AdditionalInformation     []string `json:"additional_information"`
	AddressStatus             string   `json:"address_status"`
	AmountAuthorized          float64  `json:"amount_authorized"`
	AmountCanceled            float64  `json:"amount_canceled"`
	AmountOrdered             float64  `json:"amount_ordered"`
	AmountPaID                int64    `json:"amount_paid"`
	AmountRefunded            float64  `json:"amount_refunded"`
	AnetTransMethod           string   `json:"anet_trans_method"`
	BaseAmountAuthorized      float64  `json:"base_amount_authorized"`
	BaseAmountCanceled        float64  `json:"base_amount_canceled"`
	BaseAmountOrdered         float64  `json:"base_amount_ordered"`
	BaseAmountPaID            int64    `json:"base_amount_paid"`
	BaseAmountPaidOnline      float64  `json:"base_amount_paid_online"`
	BaseAmountRefunded        float64  `json:"base_amount_refunded"`
	BaseAmountRefundedOnline  float64  `json:"base_amount_refunded_online"`
	BaseShippingAmount        float64  `json:"base_shipping_amount"`
	BaseShippingCaptured      float64  `json:"base_shipping_captured"`
	BaseShippingRefunded      float64  `json:"base_shipping_refunded"`
	CcApproval                string   `json:"cc_approval"`
	CcAvsStatus               string   `json:"cc_avs_status"`
	CcCidStatus               string   `json:"cc_cid_status"`
	CcDebugRequestBody        string   `json:"cc_debug_request_body"`
	CcDebugResponseBody       string   `json:"cc_debug_response_body"`
	CcDebugResponseSerialized string   `json:"cc_debug_response_serialized"`
	CcExpMonth                string   `json:"cc_exp_month"`
	CcExpYear                 string   `json:"cc_exp_year"`
	CcLast4                   string   `json:"cc_last4"`
	CcNumberEnc               string   `json:"cc_number_enc"`
	CcOwner                   string   `json:"cc_owner"`
	CcSecureVerify            string   `json:"cc_secure_verify"`
	CcSsIssue                 string   `json:"cc_ss_issue"`
	CcSsStartMonth            string   `json:"cc_ss_start_month"`
	CcSsStartYear             string   `json:"cc_ss_start_year"`
	CcStatus                  string   `json:"cc_status"`
	CcStatusDescription       string   `json:"cc_status_description"`
	CcTransID                 string   `json:"cc_trans_id"`
	CcType                    string   `json:"cc_type"`
	EcheckAccountName         string   `json:"echeck_account_name"`
	EcheckAccountType         string   `json:"echeck_account_type"`
	EcheckBankName            string   `json:"echeck_bank_name"`
	EcheckRoutingNumber       string   `json:"echeck_routing_number"`
	EcheckType                string   `json:"echeck_type"`
	EntityID                  int64    `json:"entity_id"`
	LastTransID               string   `json:"last_trans_id"`
	Method                    string   `json:"method"`
	ParentID                  int64    `json:"parent_id"`
	PoNumber                  string   `json:"po_number"`
	ProtectionEligibility     string   `json:"protection_eligibility"`
	QuotePaymentID            int64    `json:"quote_payment_id"`
	ShippingAmount            float64  `json:"shipping_amount"`
	ShippingCaptured          float64  `json:"shipping_captured"`
	ShippingRefunded          float64  `json:"shipping_refunded"`
	ExtensionAttributes       struct {
		VaultPaymentToken struct {
			EntityID          int64  `json:"entity_id"`
			CustomerID        int64  `json:"customer_id"`
			PublicHash        string `json:"public_hash"`
			PaymentMethodCode string `json:"payment_method_code"`
			Type              string `json:"type"`
			CreatedAt         string `json:"created_at"`
			ExpiresAt         string `json:"expires_at"`
			GatewayToken      string `json:"gateway_token"`
			TokenDetails      string `json:"token_details"`
			IsActive          bool   `json:"is_active"`
			IsVisible         bool   `json:"is_visible"`
		} `json:"vault_payment_token"`
	} `json:"extension_attributes"`
}
