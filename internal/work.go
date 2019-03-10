package unsealer

// this can be removed later
type WorkRequest struct {
    Name string
    Call func
    Params map[string]string
}

// Check request type
type CheckRequest struct {
    Name string
    Url string
}

type UnsealRequest struct {
    Name string
    Url string
}

