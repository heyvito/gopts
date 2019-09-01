package gopts

import (
	"os"
	"testing"
)

type T struct {
	Username     string
	SecretKey    string `default:"s3cr37"`
	AutoRestart  bool   `default:"true"`
	IgnoredField string `gopts:"-"`
}

type U struct {
	BoolField   bool
	IntField    int64
	SliceField  []string
	StringField string
}

func TestLoadingWithoutPrefix(t *testing.T) {
	os.Setenv("USERNAME", "rob")
	os.Setenv("AUTO_RESTART", "false")

	result := LoadEnvs(T{}).(T)
	if result.AutoRestart {
		t.Error("Failed to override default boolean value")
	}
	if result.Username != "rob" {
		t.Error("Failed to read string value from env")
	}
	if result.SecretKey != "s3cr37" {
		t.Errorf("Failed to copy default value. Expected: `%s\" Current: `%s\"", "s3cr37", result.SecretKey)
	}
}

func TestLoadingWithPrefix(t *testing.T) {
	os.Setenv("PREF_USERNAME", "rob")
	os.Setenv("PREF_AUTO_RESTART", "false")

	result := LoadEnvsWithPrefix("pref", T{}).(T)
	if result.AutoRestart {
		t.Error("Failed to override default boolean value")
	}
	if result.Username != "rob" {
		t.Error("Failed to read string value from env")
	}
	if result.SecretKey != "s3cr37" {
		t.Errorf("Failed to copy default value. Expected: `%s\" Current: `%s\"", "s3cr37", result.SecretKey)
	}
}

func TestIgnoredFields(t *testing.T) {
	os.Setenv("PREF_TEST_USERNAME", "rob")
	os.Setenv("PREF_TEST_AUTO_RESTART", "false")
	os.Setenv("PREF_TEST_IGNORED_FIELD", "hello")

	result := LoadEnvsWithPrefix("pref_test", T{}).(T)
	if result.AutoRestart {
		t.Error("Failed to override default boolean value")
	}
	if result.Username != "rob" {
		t.Error("Failed to read string value from env")
	}
	if result.SecretKey != "s3cr37" {
		t.Errorf("Failed to copy default value. Expected: `%s\" Current: `%s\"", "s3cr37", result.SecretKey)
	}
	if result.IgnoredField != "" {
		t.Errorf("Failed to ignore tagged field")
	}
}

func TestFieldTypes(t *testing.T) {
	os.Setenv("BOOL_FIELD", "1")
	os.Setenv("INT_FIELD", "27")
	os.Setenv("SLICE_FIELD", "27,C,G")
	os.Setenv("STRING_FIELD", "fnord")

	result := LoadEnvs(U{}).(U)

	if !result.BoolField {
		t.Error("Failed reading bool field")
	}

	if result.IntField != 27 {
		t.Error("Failed reading int field")
	}

	expectedSlice := []string{"27", "C", "G"}
	for _, i := range expectedSlice {
		hasItem := false
		for _, j := range result.SliceField {
			if i == j {
				hasItem = true
				break
			}
		}
		if !hasItem {
			t.Errorf("Failed reading slice field. %s should be present.", i)
		}

	}
	if result.StringField != "fnord" {
		t.Error("Failed reading string field")
	}
}

func TestEmptySlices(t *testing.T) {
	os.Setenv("BOOL_FIELD", "")
	os.Setenv("INT_FIELD", "")
	os.Setenv("SLICE_FIELD", "")
	os.Setenv("STRING_FIELD", "")
	result := LoadEnvs(U{}).(U)

	if len(result.SliceField) != 0 {
		t.Errorf("Expected SliceField to be empty, instead got len %d", len(result.SliceField))
	}
}
