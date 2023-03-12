package http

import (
	"BingGPT/internal/bingGPT"
	"encoding/json"
	"fmt"
	"net/http"
)

func Conversation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	decoder := json.NewDecoder(r.Body)
	req := &bingGPT.ConversationRequest{}
	err := decoder.Decode(&req)
	if err != nil {
		panic(err)
	}
	res, err := bingGPT.NewChat().Conversation(r.Context(), req)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(res)
}

func NewConversation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	decoder := json.NewDecoder(r.Body)
	req := &bingGPT.ConversationRequest{}
	err := decoder.Decode(&req)
	if err != nil {
		panic(err)
	}
	res, err := bingGPT.NewChat().NewConversation(r.Context(), req)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(res)
}
