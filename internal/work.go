package unsealer

const STATUSCHECK = 100
const UNSEALVAULT = 101

// this can be removed later
type WorkRequest struct {
    Name string
    Url string
    Action int
    Params map[string]string
}

type StatusCheckRequest struct {
    Name string
    Url string
}

type UnsealRequest struct {
    Name string
    Url string
    KeyNumber int
}

