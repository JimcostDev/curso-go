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

// Part representa una parte del contenido
type Part struct {
	Text string `json:"text"`
}

// Content representa el contenido principal de la solicitud
type Content struct {
	Parts []Part `json:"parts"`
}

// RequestBody ajustado para generateContent
type RequestBody struct {
	Contents []Content `json:"contents"`
}

// ResponseBody ajustado para generateContent
type ResponseBody struct {
	Candidates []struct {
		Content struct {
			Parts []Part `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

func Gemini() {
	apiKey := os.Getenv("GOOGLE_API_KEY")
	if apiKey == "" {
		log.Fatal("La variable de entorno GOOGLE_API_KEY no está definida")
	}

	// modelo de Gemini que se puede usar
	modelName := "gemini-1.5-flash-latest"
	endpoint := fmt.Sprintf(
		"https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s",
		modelName,
		apiKey,
	)

	reqBody := RequestBody{
		Contents: []Content{
			{
				Parts: []Part{
					{
						Text: "¿Qué hace interesante a Go? Responde en una sola frase.",
					},
				},
			},
		},
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatalf("Error al serializar el cuerpo JSON: %v", err)
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatalf("Error creando la petición: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error enviando la petición: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error leyendo la respuesta: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error de API: %s\n%s", resp.Status, string(bodyBytes))
	}

	var resBody ResponseBody
	if err := json.Unmarshal(bodyBytes, &resBody); err != nil {
		log.Fatalf("Error parseando JSON de respuesta: %v", err)
	}

	if len(resBody.Candidates) > 0 && len(resBody.Candidates[0].Content.Parts) > 0 {
		fmt.Println("Respuesta de Gemini:")
		fmt.Println(resBody.Candidates[0].Content.Parts[0].Text)
	} else {
		fmt.Println("No hubo respuesta de Gemini.")
	}
}
