package helpers

import "github.com/jadson-medeiros/dc-a01/domain"

type JsonResponse struct {
	Type    string      `json:"type"`
	Data    domain.Item `json:"data"`
	Message string      `json:"message"`
}
