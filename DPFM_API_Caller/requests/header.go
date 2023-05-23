package requests

type Header struct {
	DeliveryDocumentID   int     `json:"DeliveryDocumentID"`
	HeaderDeliveryStatus *string `json:"HeaderDeliveryStatus"`
	IsCancelled          *bool   `json:"IsCancelled"`
}
