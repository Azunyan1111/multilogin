package api

import "github.com/Azunyan1111/multilogin/structs"

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
	JsonResponse JsonResponse         `json:"Status"`
	User         structs.UserResponse `json:"UserResponse"`
}
