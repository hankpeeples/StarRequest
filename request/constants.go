package request

// Method is the request method type
type Method int

const (
	// GET is a GET request
	GET Method = iota
	// PUT is a PUT request
	PUT
	// POST is a POST request
	POST
	// DELETE is a DELETE request
	DELETE
)
