package bingGPT

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/fasthttp/websocket"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

const Split = `\x1e`

type CreateNewConversationResponse struct {
	ConversationId        string `json:"conversationId"`
	ClientId              string `json:"clientId"`
	ConversationSignature string `json:"conversationSignature"`
	Result                struct {
		Value   string `json:"value"`
		Message any    `json:"message"`
	} `json:"result"`
}

type ConversationResponse struct {
	Type         int `json:"type"`
	InvocationId int `json:"invocationId"`
	Item         struct {
		Messages []struct {
			Text   string `json:"text,omitempty"`
			Author string `json:"author"`
			From   struct {
				Id   string `json:"id"`
				Name any    `json:"name"`
			} `json:"from,omitempty"`
			CreatedAt     time.Time `json:"createdAt"`
			Timestamp     time.Time `json:"timestamp"`
			Locale        string    `json:"locale,omitempty"`
			Market        string    `json:"market,omitempty"`
			Region        string    `json:"region,omitempty"`
			Location      string    `json:"location,omitempty"`
			LocationHints []struct {
				Country           string `json:"country"`
				CountryConfidence int    `json:"countryConfidence"`
				State             string `json:"state"`
				City              string `json:"city"`
				CityConfidence    int    `json:"cityConfidence"`
				ZipCode           string `json:"zipCode"`
				TimeZoneOffset    int    `json:"timeZoneOffset"`
				Dma               int    `json:"dma"`
				SourceType        int    `json:"sourceType"`
				Center            struct {
					Latitude  float64 `json:"latitude"`
					Longitude float64 `json:"longitude"`
					Height    any     `json:"height"`
				} `json:"center"`
				RegionType int `json:"regionType"`
			} `json:"locationHints,omitempty"`
			MessageId uuid.UUID `json:"messageId"`
			RequestId uuid.UUID `json:"requestId"`
			Offense   string    `json:"offense"`
			Feedback  struct {
				Tag       any    `json:"tag"`
				UpdatedOn any    `json:"updatedOn"`
				Type      string `json:"type"`
			} `json:"feedback"`
			ContentOrigin string `json:"contentOrigin"`
			Privacy       any    `json:"privacy"`
			InputMethod   string `json:"inputMethod,omitempty"`
			HiddenText    string `json:"hiddenText,omitempty"`
			MessageType   string `json:"messageType,omitempty"`
			AdaptiveCards []struct {
				Type    string `json:"type"`
				Version string `json:"version"`
				Body    []struct {
					Type    string `json:"type"`
					Inlines []struct {
						Type     string `json:"type"`
						IsSubtle bool   `json:"isSubtle"`
						Italic   bool   `json:"italic"`
						Text     string `json:"text"`
					} `json:"inlines,omitempty"`
					Text string `json:"text,omitempty"`
					Wrap bool   `json:"wrap,omitempty"`
					Size string `json:"size,omitempty"`
				} `json:"body"`
			} `json:"adaptiveCards,omitempty"`
			SourceAttributions []struct {
				ProviderDisplayName string `json:"providerDisplayName"`
				SeeMoreUrl          string `json:"seeMoreUrl"`
				SearchQuery         string `json:"searchQuery"`
			} `json:"sourceAttributions,omitempty"`
			SuggestedResponses []struct {
				Text        string    `json:"text"`
				Author      string    `json:"author"`
				CreatedAt   time.Time `json:"createdAt"`
				Timestamp   time.Time `json:"timestamp"`
				MessageId   string    `json:"messageId"`
				MessageType string    `json:"messageType"`
				Offense     string    `json:"offense"`
				Feedback    struct {
					Tag       any    `json:"tag"`
					UpdatedOn any    `json:"updatedOn"`
					Type      string `json:"type"`
				} `json:"feedback"`
				ContentOrigin string `json:"contentOrigin"`
				Privacy       any    `json:"privacy"`
			} `json:"suggestedResponses,omitempty"`
			SpokenText string `json:"spokenText,omitempty"`
		} `json:"messages"`
		FirstNewMessageIndex   int       `json:"firstNewMessageIndex"`
		SuggestedResponses     any       `json:"suggestedResponses"`
		ConversationId         string    `json:"conversationId"`
		RequestId              string    `json:"requestId"`
		ConversationExpiryTime time.Time `json:"conversationExpiryTime"`
		Telemetry              struct {
			Metrics   any       `json:"metrics"`
			StartTime time.Time `json:"startTime"`
		} `json:"telemetry"`
		ShouldInitiateConversation bool `json:"shouldInitiateConversation"`
		Result                     struct {
			Value          string `json:"value"`
			Message        string `json:"message"`
			ServiceVersion string `json:"serviceVersion"`
		} `json:"result"`
	} `json:"item"`
}

type Option any

func CreateNewConversation(ctx context.Context, cookies string) (response CreateNewConversationResponse, err error) {
	const URL = "https://www.bing.com/turing/conversation/create"

	request, err := http.NewRequestWithContext(ctx, "GET", URL, nil)
	if err != nil {
		return
	}
	UUID, _ := uuid.NewUUID()
	request.Header.Set("accept", "application/json")
	request.Header.Set("content-type", "application/json")
	request.Header.Set("accept-language", "en-US,en;q=0.9")
	request.Header.Set("sec-ch-ua", `"Not_A Brand";v="99", "Microsoft Edge";v="109", "Chromium";v="109"`)
	request.Header.Set("sec-ch-ua-arch", `"x86"`)
	request.Header.Set("sec-ch-ua-bitness", `"64"`)
	request.Header.Set("sec-ch-ua-full-version", `"109.0.1518.78"`)
	request.Header.Set("sec-ch-ua-full-version-list", `"Not_A Brand";v="99.0.0.0", "Microsoft Edge";v="109.0.1518.78", "Chromium";v="109.0.5414.120"`)
	request.Header.Set("sec-ch-ua-mobile", `?0`)
	request.Header.Set("sec-ch-ua-model", ``)
	request.Header.Set("sec-ch-ua-platform", `"Windows"`)
	request.Header.Set("sec-ch-ua-platform-version", `"15.0.0"`)
	request.Header.Set("sec-fetch-dest", `empty`)
	request.Header.Set("sec-fetch-mode", `cors`)
	request.Header.Set("sec-fetch-site", `same-origin`)
	request.Header.Set("x-ms-client-request-id", UUID.String())
	request.Header.Set("x-ms-useragent", "azsdk-js-api-client-factory/1.0.0-beta.1 core-rest-pipeline/1.10.0 OS/Win32")
	request.Header.Set("cookie", cookies)
	request.Header.Set("Referer", "https://www.bing.com/search?q=Bing+AI&showconv=1&FORM=hpcodx")
	request.Header.Set("Referrer-Policy", "origin-when-cross-origin")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("CURL %s http.StatusCode = %d", URL, resp.StatusCode))
		return
	}

	body, _ := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	if response.Result.Value != "Success" {
		err = errors.New(fmt.Sprintf("CURL %s : %s", URL, string(body)))
		return
	}

	return
}

func CreateWebSocketConnection(ctx context.Context) (conn *websocket.Conn, err error) {
	const URL = "wss://sydney.bing.com/sydney/ChatHub"
	// 建立WebSocket连接
	conn, _, err = websocket.DefaultDialer.Dial(URL, nil)
	if err != nil {
		return
	}
	err = SendProtocolMessage(ctx, conn)
	return
}

func ResetWebSocketConnection(ctx context.Context, conn *websocket.Conn) (newConn *websocket.Conn, err error) {
	_ = CloseWebSocketConnection(ctx, conn)
	return CreateWebSocketConnection(ctx)
}

func CloseWebSocketConnection(ctx context.Context, conn *websocket.Conn) error {
	return conn.Close()
}

func SendMessageWebSocket(ctx context.Context, conn *websocket.Conn, message []byte) (err error) {
	err = conn.WriteMessage(websocket.TextMessage, append(message, []byte(Split)...))
	if err == nil {
		log.Println("SendMessageWebSocket success !")
	}
	return
}

func SendProtocolMessage(ctx context.Context, conn *websocket.Conn) (err error) {
	return SendMessageWebSocket(ctx, conn, []byte(`{"protocol":"json","version":1}`))
}

func SendPongMessage(ctx context.Context, conn *websocket.Conn) (err error) {
	return SendMessageWebSocket(ctx, conn, []byte(`{"type":6}`))
}

func SendConversation(ctx context.Context, conn *websocket.Conn, clientId, conversationId, conversationSignature string, invocationId int, message string, options ...Option) error {
	data := map[string]any{
		"arguments": []any{
			map[string]any{
				"source": "cib",
				"optionsSets": []any{
					"nlu_direct_response_filter",
					"deepleo",
					"enable_debug_commands",
					"disable_emoji_spoken_text",
					"responsible_ai_policy_235",
					"enablemm",
				},
				"isStartOfSession": invocationId == 0,
				"message": map[string]any{
					"author":      "user",
					"inputMethod": "Keyboard",
					"text":        message,
					"messageType": "Chat",
				},
				"conversationSignature": conversationSignature,
				"participant": map[string]any{
					"id": clientId,
				},
				"conversationId": conversationId,
			},
		},
		"invocationId": strconv.Itoa(invocationId),
		"target":       "chat",
		"type":         4,
	}
	marshal, _ := json.Marshal(data)
	return SendMessageWebSocket(ctx, conn, marshal)
}

func ListenWebSocketConnection(ctx context.Context, conn *websocket.Conn, messageChan chan string) error {
	for {
		select {
		case <-ctx.Done():
			return CloseWebSocketConnection(ctx, conn)
		default:
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				return err
			}
			switch messageType {
			case websocket.TextMessage:
				msg := string(message)
				splits := strings.Split(msg, Split)
				splits = lo.Filter(splits, func(item string, index int) bool {
					return strings.TrimSpace(item) != ""
				})
				for i := range splits {
					messageChan <- splits[i]
				}
			}
		}
	}
}
