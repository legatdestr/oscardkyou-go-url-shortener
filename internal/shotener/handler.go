package shotener

import (
	"GoUrlShotener/internal/store"
	"encoding/json"
	"net/http"
)

func SaveURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	alias := data["alias"]
	if alias == "" {
		alias = GenerateAlias() // Ensure this function is implemented elsewhere
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"shortened": "http://localhost:8080/" + alias})
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	alias := r.URL.Path[1:]

	url, ok := store.Get(alias)
	if !ok {
		http.Error(w, "не найдено", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}
