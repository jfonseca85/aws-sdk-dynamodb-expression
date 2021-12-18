package app

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

type Param struct {
	Name     string
	Required bool
	Type     string
	Default  string
}

/*
type Command struct {
	Handler func(ctx context.Context, cfg *config.Config, args map[string]string) (string, error)
	Params  []*Param
}
*/
func ValidateParams(args map[string]string, params []*Param) error {

	for _, param := range params {

		// Try to extract the value
		paramValue, ok := args[param.Name]

		// Enforce required param
		if param.Required && !ok {
			return fmt.Errorf("Invalid request for command, required param %s not found", param.Name)
		}

		// Skip optional param param not found
		if !param.Required && !ok {
			args[param.Name] = param.Default
			continue
		}

		switch param.Type {
		case "int":
			{
				if _, err := strconv.Atoi(paramValue); err != nil {
					return fmt.Errorf("Invalid request, param %s of type %s is invalid", param.Name, param.Type)
				}
			}
		case "bool":
			{
				if _, err := strconv.ParseBool(paramValue); err != nil {
					return fmt.Errorf("Invalid request, param %s of type %s is invalid", param.Name, param.Type)
				}
			}
		case "string":
			{
				if len(strings.Trim(paramValue, " ")) == 0 {
					return fmt.Errorf("Invalid request, param %s of type %s is empty", param.Name, param.Type)
				}

				// Apply placeholders
				paramValue = ApplyPlaceholderString(paramValue, args)
			}
		case "yaml":
			{
				if len(strings.Trim(paramValue, " ")) == 0 {
					return fmt.Errorf("Invalid request, param %s of type %s is empty", param.Name, param.Type)
				}

				paramValueDecoded, err := decodeBase64(paramValue)
				if err != nil {
					return fmt.Errorf("Invalid request, param %s of type %s is invalid : %v", param.Name, param.Type, err.Error())
				}
				buffer := new(bytes.Buffer)
				_, err = buffer.ReadFrom(strings.NewReader(string(paramValueDecoded)))
				if err != nil {
					return fmt.Errorf("Invalid request, param %s of type %s is invalid : %v", param.Name, param.Type, err.Error())
				}
				d := make(map[string]interface{})
				err = yaml.Unmarshal(buffer.Bytes(), &d)
				if err != nil {
					return fmt.Errorf("Invalid request, param %s of type %s is invalid : %v", param.Name, param.Type, err.Error())
				}

				// Apply placeholders
				paramValueDecoded = ApplyPlaceholderString(string(paramValueDecoded), args)

				fmt.Println(paramValueDecoded)
			}
		case "tgz":
			{

			}
		case "json":
			{
				if len(strings.Trim(paramValue, " ")) == 0 {
					return fmt.Errorf("Invalid request, param %s of type %s is empty", param.Name, param.Type)
				}

				if len(strings.Trim(paramValue, " ")) == 0 {
					return fmt.Errorf("Invalid request, param %s of type %s is empty", param.Name, param.Type)
				}

				paramValueDecoded, err := decodeBase64(paramValue)
				if err != nil {
					return fmt.Errorf("Invalid request, param %s of type %s is invalid : %v", param.Name, param.Type, err.Error())
				}

				d := make(map[string]interface{})

				err = json.Unmarshal([]byte(paramValueDecoded), &d)
				if err != nil {
					return fmt.Errorf("Invalid request, param %s of type %s is invalid : %v", param.Name, param.Type, err.Error())
				}

				// Apply placeholders
				paramValueDecoded = ApplyPlaceholderString(string(paramValueDecoded), args)
			}
		}
	}

	// No errors
	return nil
}

func decodeBase64(value string) (string, error) {

	dString, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return "", err
	}

	return string(dString), nil
}
