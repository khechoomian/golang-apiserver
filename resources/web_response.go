package resources

// Web Response Data Type
type WebResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Results interface{} `json:"results,omitempty"`
}
