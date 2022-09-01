package calculations

import "github.com/kode-magic/go-bowl/calculations/commons"

func CalculateGST(amount, gstPercent float64) commons.AddGSTData {
	gst := (amount * gstPercent) / 100
	amountWithGst := amount + gst
	
	return commons.AddGSTData{
		Gst:           gst,
		AmountWithGst: amountWithGst,
	}
}

func RemoveGST(gstAmount, gstPercent float64) commons.RemoveGSTData {
	gst := gstAmount - (gstAmount * (100 / (100 + gstPercent)))
	amountWithoutGst := gstAmount - gst

	return commons.RemoveGSTData{
		Gst:              gst,
		AmountWithoutGst: amountWithoutGst,
	}
}

//179126
