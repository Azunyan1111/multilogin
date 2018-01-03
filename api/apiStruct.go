package api

type JsonResponse struct {
	StatusCode int	`json:"StatusCode"`
	Message string	`json:"Message"`
}

type SumResponse struct {
	JsonResponse JsonResponse `json:"Status"`
	Value        int          `json:"Value"`
}

type NameResponse struct {
	JsonResponse JsonResponse `json:"Status"`
	Name        string          `json:"Name"`
}
