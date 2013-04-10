package level

type Error uint8

const (
	Already_Open Error = iota
)

func (e Error) String() (o string){
	switch e{
	case Already_Open:
		o = "Database was already open."
	}
	return
}
