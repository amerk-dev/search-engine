package restapi

import (
	"encoding/json"
	"log"
	"net/http"
	"searchSystem/internal/index"
	"strconv"
)

type Server struct {
	Index *index.Index
}

func (s Server) StartServer() {
	handler := http.NewServeMux()
	handler.HandleFunc("/search/", s.SearchHandler)

	log.Println("Сервер запущен на http://localhost:8080")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}

type SearchResponse struct {
	Response string `json:"response"`
}

func (s Server) SearchHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	req := SearchResponse{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	res := s.Index.Search(req.Response)
	log.Println(res)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("найдено результатов - " + strconv.Itoa(len(res))))
}
