package environment

import (
	"bytes"
	"errors"
	"fmt"
)

// Defined Environment as int8.
type Environment int8

const (
	Local                     Environment = iota // Local development environment. (LOCAL)
	Development                                  // Development environment. (DEV)
	PreProduction                                // Pre-production environment. (PRE-PROD)
	Production                                   // Production environment. (PRD/PROD)
	UnitTest                                     // Unit test environment. (UT)
	SystemIntegrationTest                        // System Integration test environment. (SIT/IT)
	SystemTest                                   // System test environment. (ST)
	UserAcceptanceTest                           // User acceptance test environment. (UAT)
	PerformanceEvaluationTest                    // Performance evaluation test environment. (PET)
)

// String returns a lower-case ASCII representation of the environment.
func (env Environment) String() string {
	switch env {
	case Local:
		return "Local"
	case Development:
		return "Development"
	case PreProduction:
		return "Pre-Production"
	case Production:
		return "Production"
	case UnitTest:
		return "UnitTest"
	case SystemIntegrationTest:
		return "SystemIntegrationTest"
	case SystemTest:
		return "SystemTest"
	case UserAcceptanceTest:
		return "UserAcceptanceTest"
	case PerformanceEvaluationTest:
		return "PerformanceEvaluationTest"
	default:
		return fmt.Sprintf("Env(%d)", env)
	}
}

// MarshalText marshals the Environment to text. Note that the text representation
// drops the -Level suffix (see example).
func (env Environment) MarshalText() ([]byte, error) {
	return []byte(env.String()), nil
}

// UnmarshalText unmarshals text to a environment. Like MarshalText, UnmarshalText
// expects the text representation of a Environment to drop the -Level suffix (see
// example).
//
// In particular, this makes it easy to configure logging levels using YAML,
// TOML, or JSON files.
func (env *Environment) UnmarshalText(text []byte) error {
	if env == nil {
		return errors.New("can't unmarshal a nil *Level")
	}
	if !env.unmarshalText(text) && !env.unmarshalText(bytes.ToLower(text)) {
		return fmt.Errorf("unrecognized level: %q", text)
	}
	return nil
}

func (env *Environment) unmarshalText(text []byte) bool {
	switch string(bytes.ToLower(text)) {
	case "local":
		*env = Local
	case "development", "dev", "": // make the zero value useful
		*env = Development
	case "pre-production", "pre-prod", "pre-prd":
		*env = PreProduction
	case "production", "prod", "prd":
		*env = Production
	case "unittest", "unit-test", "ut":
		*env = UnitTest
	case "systemintegrationtest", "system-integration-test", "integration-test", "sit", "it":
		*env = SystemIntegrationTest
	case "systemtest", "system-test", "st":
		*env = SystemTest
	case "useracceptancetest", "user-acceptance-test", "uat":
		*env = UserAcceptanceTest
	case "performanceevaluationtest", "performance-evaluation-test", "pet":
		*env = PerformanceEvaluationTest
	default:
		return false
	}
	return true
}

// Set sets the environment for the flag.Value interface.
func (env *Environment) Set(s string) error {
	return env.UnmarshalText([]byte(s))
}

// Get gets the environment for the flag.Getter interface.
func (env *Environment) Get() interface{} {
	return *env
}
