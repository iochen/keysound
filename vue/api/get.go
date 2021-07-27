package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetQuizHandler(w http.ResponseWriter, r *http.Request) {
	if (r.URL.Path != "/api/get" && r.URL.Path != "/get") ||
		r.Method != http.MethodPost {
		fmt.Printf("unknown request with path: %s, method: %s\n",
			r.URL.Path, r.Method)
		w.WriteHeader(400)
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	id := r.PostForm.Get("id")
	id = strings.TrimSpace(id)
	fmt.Println(id)
	plr, err := MongoGetQuiz(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(404)
			return
		}
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	err = json.NewEncoder(w).Encode(NewResponse{
		ID:     id,
		Base:   os.Getenv("BASE_URL"),
		Pieces: plr,
	})
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
}
