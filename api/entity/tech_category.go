package entity

const (
	TechCategoryFrontend       = "Frontend"
	TechCategoryBackend        = "Backend"
	TechCategoryDatabase       = "Database"
	TechCategoryInfrastructure = "Infrastructure"
)

// IsValidTechCategory は条件を判定します。
func IsValidTechCategory(category string) bool {
	switch category {
	case TechCategoryFrontend, TechCategoryBackend, TechCategoryDatabase, TechCategoryInfrastructure:
		return true
	default:
		return false
	}
}
