package utils

import "github.com/gofiber/fiber/v2"

func MapSlice[T any, V any](input []T, f func(T) V) []V {
	result := make([]V, len(input))
	for i, v := range input {
		result[i] = f(v)
	}
	return result
}

// Helper function to map error codes to HTTP status codes
func GetStatusCode(code string) int {
	switch code {
	case "NOT_FOUND":
		return fiber.StatusNotFound
	case "DATABASE_ERROR":
		return fiber.StatusInternalServerError
	default:
		return fiber.StatusInternalServerError
	}
}
