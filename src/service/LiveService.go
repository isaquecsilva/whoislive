package service

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/isaquecsilva/whoislive/src/domain"
	"github.com/redis/go-redis/v9"
)

type LiveServiceInterface interface {
	GetLiveChannels() (domain.Channels, error)
}

type LiveService struct {
	endpoint       string
	payload        string
	clientIdHeader string
	cache          *redis.Client
}

func (ls *LiveService) GetLiveChannels() (domain.Channels, error) {
	scmd := ls.cache.Get(context.Background(), "LIVE")
	buf, err := scmd.Bytes()
	if err !=  nil {
		slog.Warn("REDIS_CACHE", "err_message", "`key` does not exist")
	}

	var channels domain.Channels

	if buf != nil {
		slog.Info("REDIS_CACHE", "message", "`key` found")
		switch err := json.Unmarshal(buf, &channels); err {
		case nil:
			return channels, nil
		default:
			slog.Warn("JSON_UNMARSHAL_REDIS", "message", err.Error())	
		}
	}

	req, err := http.NewRequest(http.MethodPost, ls.endpoint, strings.NewReader(ls.payload))

	if err != nil {
		slog.Error("HTTP_REQUEST", "message", err.Error())
		return nil, err
	}

	req.Header.Add("Client-Id", ls.clientIdHeader)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error("HTTP_RESPONSE", "message", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	
	channels, err = ls.parseData(resp.Body)
	if err != nil {
		return nil, err
	}

	if buf, err := json.Marshal(channels); err != nil {
		slog.Error("JSON_CHANNELS_MARSHALING", "message", err.Error())
	} else {
		ls.cache.Set(context.Background(), "LIVE", buf, time.Minute * 2)
	}

	return channels, nil	
}

func (ls *LiveService) parseData(data io.Reader) (domain.Channels, error) {
	var jsonData []map[string]any

	if err := json.NewDecoder(data).Decode(&jsonData); err != nil {
		slog.Error("JSON_RAW_PARSING", "message", err.Error())
		return nil, err
	}

	if len(jsonData) == 0 {
		var err error = errors.New("twitch response array is zero-length")
		slog.Error("JSON_DATA LENGTH", "message", err.Error())
		return nil, err
	}

	itemList, ok := jsonData[0]["data"].(map[string]any)["personalSections"].([]any)[0].(map[string]any)["items"].([]any)

	if !ok {
		var err error = errors.New("getting `items` key from json: invalid cast")
		slog.Error("RESPONSE_JSON", "message", err.Error())
		return nil, err
	}

	jsonData = nil
	channels := make(domain.Channels, 0, len(itemList))

	for _, item := range itemList {
		user := item.(map[string]interface{})["user"].(map[string]interface{})

		channels = append(channels, domain.NewChannel(user["login"].(string), user["displayName"].(string), user["profileImageURL"].(string)))
	}

	return channels, nil
}

func NewLiveService() *LiveService {
	cache := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "whoislive",
		DB:       0,
	})

	status, err := cache.Ping(context.Background()).Result()
	log.Println(status, err)

	return &LiveService{
		endpoint:       "https://gql.twitch.tv/gql",
		payload:        `[{"operationName":"PersonalSections","variables":{"input":{"sectionInputs":["RECOMMENDED_SECTION"],"recommendationContext":{"platform":"web"}},"creatorAnniversariesFeature":false},"extensions":{"persistedQuery":{"version":1,"sha256Hash":"1ad4a64ba2d092adc054604d66d3f4fb899ebddbcfcd931fee5796fa6bb597c0"}}}]`,
		clientIdHeader: "kimne78kx3ncx6brgo4mv6wki5h1ko",
		cache:          cache,
	}
}
