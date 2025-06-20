package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Deepseek() {
	// Obtener API key de DeepSeek desde variables de entorno
	apiKey := os.Getenv("DEEPSEEK_API_KEY")
	if apiKey == "" {
		log.Fatal("❌ DEEPSEEK_API_KEY no está configurada")
	}

	url := "https://api.deepseek.com/v1/chat/completions"

	// Crear cuerpo de solicitud
	requestBody, err := json.Marshal(map[string]interface{}{
		"model": "deepseek-chat",
		"messages": []map[string]string{
			{"role": "user", "content": "¿Que hace a Golang tan especial? Responde en una sola frase."},
		},
		"max_tokens": 200,
	})
	if err != nil {
		log.Fatalf("Error creando cuerpo: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalf("Error creando solicitud: %v", err)
	}

	// Agregar headers de autenticación
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error en solicitud HTTP: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error leyendo respuesta: %v", err)
	}

	// Verificar estado de respuesta
	if resp.StatusCode != http.StatusOK {
		log.Printf("❌ Estado inesperado: %s", resp.Status)
		log.Printf("Respuesta cruda: %s", body)
		return
	}

	// Parsear respuesta JSON
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Error parseando JSON: %v", err)
	}

	// Extraer contenido de manera segura
	if choices, ok := result["choices"].([]interface{}); ok && len(choices) > 0 {
		if firstChoice, ok := choices[0].(map[string]interface{}); ok {
			if message, ok := firstChoice["message"].(map[string]interface{}); ok {
				if content, ok := message["content"].(string); ok {
					fmt.Println(content)
					return
				}
			}
		}
	}

	log.Println("❌ Estructura de respuesta inesperada")
	log.Printf("Respuesta completa: %s", body)
}
