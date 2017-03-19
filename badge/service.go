package badge

import "fmt"

// Service interface
type Service interface {
	Badge(num float64, status int) string
}

type service struct{}

// NewService func
func NewService() Service {
	return &service{}
}

// Badge generate url badge
func (s *service) Badge(num float64, status int) string {
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
		}
		return fmt.Sprintf("https://img.shields.io/badge/coverage-%.1f%%25-brightgreen.svg?style=flat", num)
	case 2:
		return fmt.Sprintf("https://img.shields.io/badge/coverage-error-red.svg?style=flat")
	default:
		return fmt.Sprintf("https://img.shields.io/badge/coverage-unknown-lightgrey.svg?style=flat")
	}
}
