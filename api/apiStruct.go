package api

type JsonResponse struct {
	StatusCode int    `json:"StatusCode"`
	Message    string `json:"Message"`
}

type SumResponse struct {
	JsonResponse JsonResponse `json:"Status"`
	Value        int          `json:"Value"`
}

type NameResponse struct {
	JsonResponse JsonResponse `json:"Status"`
	Name         string       `json:"Name"`
}

type ImageResponse struct {
	JsonResponse JsonResponse `json:"Status"`
	Image        string       `json:"Image"`
}

type AgeResponse struct {
	JsonResponse JsonResponse `json:"Status"`
	Age          string       `json:"Image"`
}

type BirthdayResponse struct {
	JsonResponse JsonResponse `json:"Status"`
	Birthday     string       `json:"Birthday"`
}

type EmailResponse struct {
	JsonResponse JsonResponse `json:"Status"`
	Email        string       `json:"Email"`
}

type PhoneResponse struct {
	JsonResponse JsonResponse `json:"Status"`
	Phone        string       `json:"Phone"`
}

type AddressResponse struct {
	JsonResponse JsonResponse `json:"Status"`
	Address      string       `json:"Address"`
}

type AllResponse struct {
	JsonResponse JsonResponse `json:"Status"`
	User         UserResponse `json:"User"`
}

type UserResponse struct {
	Uid      string `json:"Uuid"`
	UserName string `json:"UserName"`
	Email    string `json:"Email"`
	Image    string `json:"Image"`
	Age      string `json:"Age"`
	Birthday string `json:"Birthday"`
	Phone    string `json:"Phone"`
	Address  string `json:"Address"`
}
