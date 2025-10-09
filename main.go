package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"reflect"
	"syscall"

	"github.com/steampoweredtaco/gotiktoklive"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go <username_tiktok>")
		return
	}

	username := os.Args[1]
	fmt.Printf("Conectando al livestream de: @%s\n", username)

	// Crear instancia de TikTok
	tiktok, err := gotiktoklive.NewTikTok()
	if err != nil {
		log.Fatalf("Error creando cliente TikTok: %v", err)
	}

	// Trackear usuario por username
	live, err := tiktok.TrackUser(username)
	if err != nil {
		log.Fatalf("Error conectando al stream: %v", err)
	}

	fmt.Println("Conectado exitosamente!")
	fmt.Println("==================================================")
	fmt.Println("Escuchando comentarios del chat... (Ctrl+C para salir)")

	// Recibir eventos del livestream
	go func() {
		for event := range live.Events {
			// Usar reflexión para acceder a los campos
			eventValue := reflect.ValueOf(event)
			if eventValue.Kind() == reflect.Ptr {
				eventValue = eventValue.Elem()
			}

			// Buscar el campo Comment
			commentField := eventValue.FieldByName("Comment")
			if commentField.IsValid() && commentField.Kind() == reflect.String && commentField.String() != "" {
				comment := commentField.String()

				// Buscar el campo User
				userField := eventValue.FieldByName("User")
				if userField.IsValid() {
					username := extractUsername(userField)
					if username != "" {
						fmt.Printf("💬 %s: %s\n", username, comment)
					} else {
						fmt.Printf("💬 %s\n", comment)
					}
				} else {
					fmt.Printf("💬 %s\n", comment)
				}
			}
		}
	}()

	// Esperar señal para terminar
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	fmt.Println("\nSaliendo...")
}

func extractUsername(userValue reflect.Value) string {
	if userValue.Kind() == reflect.Ptr {
		if userValue.IsNil() {
			return ""
		}
		userValue = userValue.Elem()
	}

	if userValue.Kind() == reflect.Struct {
		// Intentar campos comunes para username
		possibleFields := []string{"Username", "Nickname", "DisplayName", "UniqueID", "Name"}

		for _, fieldName := range possibleFields {
			field := userValue.FieldByName(fieldName)
			if field.IsValid() && field.Kind() == reflect.String && field.String() != "" {
				return field.String()
			}
		}

		// Si no encontramos campos específicos, explorar todos los campos string
		for i := 0; i < userValue.NumField(); i++ {
			field := userValue.Field(i)
			if field.Kind() == reflect.String && field.String() != "" {
				fieldName := userValue.Type().Field(i).Name
				// Ignorar campos que probablemente no sean usernames
				if fieldName != "ID" && fieldName != "Email" && fieldName != "Avatar" {
					return field.String()
				}
			}
		}
	}

	return ""
}
