package config

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/kelseyhightower/envconfig"
)

type EnvVars struct {
	OriginalAnimalsDataFile  string `envconfig:"ORIGINAL_DATA_FILE" required:"true"`
	ToReplaceAnimalsDataFile string `envconfig:"TO_REPLACE_DATA_FILE" required:"true"`
	CredentialsFile          string `envconfig:"CREDENTIALS_FILE" required:"true"`
	TokenFile                string `envconfig:"TOKEN_FILE" required:"true"`
	SpreadSheetID            string `envconfig:"SPREADSHEET_ID" required:"true"`
	SpreadSheetRange         string `envconfig:"SPREADSHEET_RANGE" required:"true"`
}

func MustNew() *EnvVars {
	env := mustGetEnvVars()
	return &env
}

func mustGetEnvVars() EnvVars {
	var env EnvVars
	err := envconfig.Process("", &env)
	if err != nil {
		panic(err)
	}
	mustHaveNoBlankEnvVars(env)
	return env
}
func mustHaveNoBlankEnvVars(env EnvVars) {
	blanks := findBlankEnvVars(env)
	if len(blanks) > 0 {
		panic(fmt.Errorf("the following environment variables are blank: %s", strings.Join(blanks, ", ")))
	}
}

func findBlankEnvVars(env EnvVars) []string {
	var blanks []string
	valueOfStruct := reflect.ValueOf(env)
	typeOfStruct := valueOfStruct.Type()
	for i := 0; i < valueOfStruct.NumField(); i++ {
		if valueOfStruct.Field(i).Interface() == "" {
			blanks = append(blanks, typeOfStruct.Field(i).Name)
		}
	}
	return blanks
}
