package internal

// StatusCheckRequest - The type for a status check request
type StatusCheckRequest struct {
	Name   string
	Url    string
	Domain string
}

// UnsealRequest - The type for unsealing requests
type UnsealRequest struct {
	Name      string
	Url       string
	KeyNumber int
}
