package library

import "testing"

func TestEnvUsesStagingPrefixOutsideProduction(t *testing.T) {
	t.Setenv("IS_PRODUCTION", "false")
	t.Setenv("DATABASE_URL", "production")
	t.Setenv("STAGING_DATABASE_URL", "staging")

	if got := Env("DATABASE_URL"); got != "staging" {
		t.Fatalf("Env(DATABASE_URL) = %q, want staging", got)
	}
}

func TestEnvUsesProductionValueInProduction(t *testing.T) {
	t.Setenv("IS_PRODUCTION", "true")
	t.Setenv("DATABASE_URL", "production")
	t.Setenv("STAGING_DATABASE_URL", "staging")

	if got := Env("DATABASE_URL"); got != "production" {
		t.Fatalf("Env(DATABASE_URL) = %q, want production", got)
	}
}
