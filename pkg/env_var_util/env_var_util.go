package env_var_util

import (
	"fmt"
	"os"
)

func CheckEnvVars(requiredEnvVars []string) error {
	notFoundEnvVars := 0
	for _, envVar := range requiredEnvVars {
		if len(os.Getenv(envVar)) == 0 {
			notFoundEnvVars++
			fmt.Printf("Env variable %s is empty\n", envVar)
		}
	}
	if notFoundEnvVars == 0 {
		return nil
	}
	return fmt.Errorf("%d envVars not found", notFoundEnvVars)
}
