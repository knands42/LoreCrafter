package integration

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Setup integration test environment
	err := SetupIntegrationTest()
	if err != nil {
		log.Fatalf("Failed to setup integration test environment: %v", err)
	}

	// Run tests
	exitCode := m.Run()

	// Teardown integration test environment
	TeardownIntegrationTest()

	// Exit with the same code as the tests
	os.Exit(exitCode)
}