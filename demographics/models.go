package demographics

type (
	Demographics struct {
		Demographics []District `json:"demographics"`
		//Demographics []struct {
		//	Region    string   `json:"region"`
		//	District  string   `json:"district"`
		//	Chiefdoms []string `json:"chiefdoms"`
		//} `json:"demographics"`
	}

	District struct {
		Region    string   `json:"region"`
		District  string   `json:"district"`
		Chiefdoms []string `json:"chiefdoms"`
	}
)
