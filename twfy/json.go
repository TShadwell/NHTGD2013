package twfy

type (
	jsonMPsResponse []Member
	jsonResponse    struct {
		Error Error
	}
	jsonHansard struct {
		Rows []jsonRow `json:"row"`
	}

	jsonRow struct {
		Body string `json:"body"`
	}
)
