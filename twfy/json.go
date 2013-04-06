package twfy

type (
	jsonMPsResponse []Member
	jsonResponse    struct {
		Error Error
	}
	JsonHansard struct{
		Rows []JsonRow `json:"rows"`
	}

	JsonRow struct {
		Body string `json:"body"`
	}
)
