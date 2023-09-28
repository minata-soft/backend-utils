package backend_utils

import (
	"encoding/json"
	"fmt"
)

func FormatJSON(data interface{}) (string, error) {
	// Formatage du JSON avec des indentations et des sauts de ligne
	indentedJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", fmt.Errorf("erreur lors du formatage du JSON : %s", err.Error())
	}

	// Conversion du JSON formaté en une chaîne de caractères
	formattedJSON := string(indentedJSON)

	return formattedJSON, nil
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func PrintPretty(p interface{}) {
	jsonBytes, _ := json.MarshalIndent(p, "", "  ")
	fmt.Println(string(jsonBytes))
}
