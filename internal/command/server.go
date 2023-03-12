package command

import (
	httpApi "BingGPT/internal/http"
	websocketApi "BingGPT/internal/websocket"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
)

func postOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/" && r.Method != "POST" {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	})
}

var Server = &cli.Command{
	Name:  "server",
	Usage: "BingGPT server api or ws",
	Action: func(c *cli.Context) error {
		mux := http.NewServeMux()
		mux.HandleFunc("/", websocketApi.Conversation)
		mux.HandleFunc("/conversation", httpApi.Conversation)
		mux.HandleFunc("/conversation/create", httpApi.NewConversation)
		fmt.Println("Listening on :12527")
		if err := http.ListenAndServe(":12527", postOnly(mux)); err != nil {
			log.Fatal(err)
		}
		return nil
	},
}
