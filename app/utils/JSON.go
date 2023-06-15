package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {

	maxBytes := int64(1048576)

	r.Body = http.MaxBytesReader(w, r.Body, maxBytes)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(data)
	if err != nil {
		return err
	}

	err = decoder.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("something went wrong")
	}

	return nil
}

func WriteJSON(w http.ResponseWriter, r *http.Request, data interface{}, status int, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header()[k] = v
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(out)
	return nil

}
