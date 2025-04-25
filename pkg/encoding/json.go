package encoding

import (
	"encoding/json"
	"io"
)

func WriteJSON(w io.Writer, data any) error {
	return json.NewEncoder(w).Encode(data)
}
