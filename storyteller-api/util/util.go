package util

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
)

// TODO: Dotenv does not provide functionality for auto searching in all
// parental directories. This code is just from stackOverFlow, better to rewrite
func LoadEnv() {
	projectName := regexp.MustCompile(`^(.*` + "fablescope-backend" + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func ReadJSON(w http.ResponseWriter, r *http.Request, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewDecoder(r.Body).Decode(v)
}

// Might be good to make it return error
func SliceFieldToString(slice any, fieldName string) string {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return ""
	}

	var builder strings.Builder
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		field := elem.FieldByName(fieldName)
		if !field.IsValid() {
			continue
		}
		builder.WriteString(fmt.Sprintf("%v", field.Interface()))
		if i < v.Len()-1 {
			builder.WriteString(", ")
		}
	}

	return builder.String()
}
