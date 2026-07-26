package entity

import "testing"

func TestIsValidTechCategory(t *testing.T) {
	validCategories := []string{
		TechCategoryFrontend,
		TechCategoryBackend,
		TechCategoryDatabase,
		TechCategoryInfrastructure,
	}
	for _, category := range validCategories {
		if !IsValidTechCategory(category) {
			t.Errorf("expected category %q to be valid", category)
		}
	}

	if IsValidTechCategory("Newsletter") {
		t.Error("expected Newsletter to be invalid")
	}
}
