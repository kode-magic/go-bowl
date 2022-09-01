package commons

type AddGSTData struct {
	Gst float64 `json:"gst"`
	AmountWithGst float64 `json:"amountWithGst"`
}

type RemoveGSTData struct {
	Gst float64 `json:"gst"`
	AmountWithoutGst float64 `json:"amountWithGst"`
}