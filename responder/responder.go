package responder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func JSON(res http.ResponseWriter, req *http.Request, v interface{}, status int) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	if err := enc.Encode(v); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Content-Length", strconv.Itoa(buf.Len()))

	if status == 0 {
		res.WriteHeader(http.StatusOK)
	} else {
		res.WriteHeader(status)
	}

	if _, err := buf.WriteTo(res); err != nil {
		fmt.Errorf("Error writing JSON response: %+v", err)
	}
}
