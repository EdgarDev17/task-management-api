package response

// ResponseError estructura estándar para errores
type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
