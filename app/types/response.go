package types

type Response struct {
	StatusCode int
	Body       string
	Headers    map[string]string
}
