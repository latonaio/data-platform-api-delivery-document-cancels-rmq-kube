package requests

type Item struct {
	DeliveryDocumentID int     `json:"DeliveryDocumentID"`
	DeliveryDocument   int     `json:"DeliveryDocument"`
	ItemDeliveryStatus *string `json:"ItemDeliveryStatus"`
	IsCancelled        *bool   `json:"IsCancelled"`
}
