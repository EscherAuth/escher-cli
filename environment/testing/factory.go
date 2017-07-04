package testing

import (
	"os"
	"testing"
)

func SetEnvForTheTest(t testing.TB, key, value string) func() {

	orgEnvValue := os.Getenv(key)
	err := os.Setenv(key, value)

	if err != nil {
		t.Fatal(err)
	}

	return func() {
		err := os.Setenv(key, orgEnvValue)

		if err != nil {
			t.Fatal(err)
		}
	}

}
