package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Identity(id string) string {
	return fmt.Sprintf("#%s", id)
}

func Api(endpoint string) string {
	return fmt.Sprintf("/api%s", endpoint)
}

func FetchData(endpoint string) (interface{}, error) {
	formattedEndpoint := fmt.Sprintf("http://localhost:2340/api/%s", endpoint)
	resp, err := http.Get(formattedEndpoint)
	fmt.Printf("Respuesta: %v, Error: %v\n", resp, err)
	if err != nil {
		// println("This is an error", err)
		return nil, err
	}
	// !! IMPORTANTE: Buena práctica para dejar los sockets cerrados y evitar leaks !!
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
