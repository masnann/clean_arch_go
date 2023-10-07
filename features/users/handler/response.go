package handler

type CreateResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoginResponse struct {
	Name  string `json:"name"`
	Token any    `json:"token"`
}

type GetALLResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
