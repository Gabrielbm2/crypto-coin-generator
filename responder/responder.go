package responder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// O package "responder" contém uma função JSON que é responsável por serializar dados em formato JSON e enviá-los como resposta HTTP. Essa função é usada para encapsular o processo de serialização de JSON e garantir que a resposta HTTP seja configurada corretamente.

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
