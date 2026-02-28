package config

import (
	"path/filepath"
	"testing"
)

func setupTestConfig(t *testing.T) (*Config, string) {
	// setup a temp directory for testing
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, ".gatorconfig.json")

	return &Config{
		DBURL:           "postgres://testuser@localhost:5432/testdb",
		CurrentUserName: "test_wizard",
	}, configPath

}
func TestReadWriteConfig(t *testing.T) {
	testCfg, configPath := setupTestConfig(t)

	// Test writing
	err := writeToPath(configPath, *testCfg)
	if err != nil {
		t.Fatalf("writeToPath() failed: %v", err)
	}

	//test reading
	gotCfg, err := readFromPath(configPath)
	if err != nil {
		t.Fatalf("readFromPath() failed: %v", err)
	}

	// assert equal
	if gotCfg.DBURL != testCfg.DBURL {
		t.Errorf("expected DBURL %s, got %s", testCfg.DBURL, gotCfg.DBURL)
	}
	if gotCfg.CurrentUserName != testCfg.CurrentUserName {
		t.Errorf("expected iser %s, got %s", testCfg.CurrentUserName, gotCfg.CurrentUserName)
	}
}
