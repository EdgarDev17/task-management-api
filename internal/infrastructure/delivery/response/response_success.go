package response

// ResponseSuccess estructura estándar para respuestas exitosas
type ResponseSuccess struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
