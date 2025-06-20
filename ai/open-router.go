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

func OpenRouter() {
	// Obtener API key de variable de entorno (mejor práctica)
	apiKey := os.Getenv("OPENROUTER_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENROUTER_API_KEY no está configurada")
	}

	url := "https://openrouter.ai/api/v1/chat/completions"

	// Configurar solicitud
	payload := map[string]interface{}{
		"model": "mistralai/mistral-7b-instruct:free",
		"messages": []map[string]string{
			{"role": "user", "content": "¿Que hace a Golang tan especial? Responde en una sola frase."},
		},
	}

	jsonPayload, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatal("Error creando solicitud:", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("HTTP-Referer", "http://localhost")
	req.Header.Set("X-Title", "UFC Query App")

	// Enviar solicitud
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error enviando solicitud:", err)
	}
	defer resp.Body.Close()

	// Leer respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error leyendo respuesta:", err)
	}

	// Verificar estado HTTP
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error %d: %s", resp.StatusCode, string(body))
		return
	}

	// Parsear respuesta
	var response struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
		Error struct {
			Message string `json:"message"`
		} `json:"error"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatal("Error parseando JSON:", err)
	}

	if response.Error.Message != "" {
		log.Printf("Error en API: %s", response.Error.Message)
	} else if len(response.Choices) > 0 {
		fmt.Println(response.Choices[0].Message.Content)
	} else {
		fmt.Println("No se obtuvo respuesta")
	}
}
