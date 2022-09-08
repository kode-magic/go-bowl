package demographics

type (
	Demographics struct {
		Demographics []struct {
			District string `json:"district"`
			Region   string `json:"region"`
		} `json:"demographics"`
	}

	District struct {
		Name   string `json:"name"`
		Region string `json:"region"`
	}
)
