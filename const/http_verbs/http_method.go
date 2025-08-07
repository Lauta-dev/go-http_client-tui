package constants

const (
	MethodGet    = "GET"
	MethodPost   = "POST"
	MethodPut    = "PUT"
	MethodDelete = "DELETE"
)

func MethodList() []string {
	methods := []string{MethodGet, MethodPost, MethodPut, MethodDelete}
	return methods
}
