package modals

type Request struct {
	Method  string
	Path    string
	Headers map[string]string
	Raw     string
}

