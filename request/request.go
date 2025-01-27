package request // import "github.com/thanks173/wc-api-go/request"

import (
	"net/url"
)

// Request ...
type Request struct {
	Method   string
	Endpoint string
	Values   url.Values
	Body     interface{}
}
