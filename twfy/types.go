/*
	Package twfy impliments basic Go bindings
	for the TheyWorkForYou API.
*/
package twfy

const (
	Liberal_Democrat Party = "liberal democrat"
	Conservative Party = "conservative"
	Labour Party = "labour"
)

type (
	//Represents one person- an MP usually
	PersonID uint64
	MemberID uint64
	//Your TWFY API key
	ApiKey string
	Error string

	API struct {
		Key ApiKey
	}

	Party string
	Constituency string

	Member struct {
		MemberID  MemberID `json:"member_id"`
		PersonID PersonID `json:"person_id"`


		Name string `json:"name"`
		Party Party `json:"party"`
		Constituency Constituency `json:"constituency"`
	}

)

func (e Error) Error() string {
	return string(e)
}

