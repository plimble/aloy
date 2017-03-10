package badge

import "fmt"

type ServiceInterface interface {
	Badge(num float64, status int) string
}

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Badge(num float64, status int) string {
	switch status {
	case 0:
		return fmt.Sprintf("https://img.shields.io/badge/coverage-pending-lightgrey.svg?style=flat")
	case 1:
		if num < 25.0 {
			return fmt.Sprintf("https://img.shields.io/badge/coverage-%.1f%%25-red.svg?style=flat", num)
		} else if num < 50.0 {
			return fmt.Sprintf("https://img.shields.io/badge/coverage-%.1f%%25-orange.svg?style=flat", num)
		} else if num < 75.0 {
			return fmt.Sprintf("https://img.shields.io/badge/coverage-%.1f%%25-green.svg?style=flat", num)
		} else {
			return fmt.Sprintf("https://img.shields.io/badge/coverage-%.1f%%25-brightgreen.svg?style=flat", num)
		}
	case 2:
		return fmt.Sprintf("https://img.shields.io/badge/coverage-error-red.svg?style=flat")
	default:
		return fmt.Sprintf("https://img.shields.io/badge/coverage-unknown-lightgrey.svg?style=flat")
	}
}
