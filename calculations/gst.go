package calculations

import "github.com/kode-magic/go-bowl/calculations/commons"

func CalculateGST(amount, gstPercent float64) commons.AddGSTData {
	gst := (amount * gstPercent) / 100
	amountWithGst := amount + gst
	return commons.AddGSTData{
		Gst: gst,
		AmountWithGst: amountWithGst,
	}
}



//179126