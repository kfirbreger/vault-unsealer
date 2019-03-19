package internal

type StatusCheckRequest struct {
	Name string
	Url  string
}

type UnsealRequest struct {
	Name      string
	Url       string
	KeyNumber int
}
