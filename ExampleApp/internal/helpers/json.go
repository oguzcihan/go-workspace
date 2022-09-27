package helpers

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

const (
	jsonKey   = "Content-Type"
	jsonValue = "application/json"
)

// aynı anda request gelirse bir sorun olur mu?
func JSON(w http.ResponseWriter, code int, res interface{}) {
	w.Header().Set(jsonKey, jsonValue)
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		Logger.Error("json_encode_error", zap.Error(err))
	}
}
