package service

func Create() (interface{}, error) {
	// Implementing connection to gRPC, HTTP or any other microservice...
	return struct {
		ID string `json:"id"`
	}{
		ID: "123",
	}, nil
}
