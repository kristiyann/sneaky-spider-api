package httpclient

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"golang.org/x/exp/slog"
)

type httpClientImpl struct {
	client     http.Client
	baseUrl    string
	clientPool *sync.Pool
	maxRetries int
	retryDelay time.Duration
	log        *log.Logger
}

type HttpClient interface {
	Get(context.Context, string, *map[string]string, *map[string]string) (*http.Response, error)
}

func New(baseUrl string) HttpClient {
	logger := log.New(os.Stdout, "["+baseUrl+"] ", log.LstdFlags)

	return &httpClientImpl{
		client: http.Client{
			Timeout: time.Minute,
		},
		baseUrl:    baseUrl,
		log:        logger,
		maxRetries: 1,
		retryDelay: 2 & time.Second,
		clientPool: nil,
	}
}

func NewWithoutLog(baseUrl string) HttpClient {
	return &httpClientImpl{
		client: http.Client{
			Timeout: time.Minute,
		},
		baseUrl:    baseUrl,
		maxRetries: 1,
		retryDelay: 2 & time.Second,
		log:        nil,
		clientPool: nil,
	}
}

func NewWithClientPool(baseUrl string, maxRetries int, logger *log.Logger) HttpClient {
	clientPool := sync.Pool{
		New: func() interface{} {
			return &http.Client{}
		},
	}

	return &httpClientImpl{
		client: http.Client{
			Timeout: time.Minute,
		},
		baseUrl:    baseUrl,
		log:        logger,
		clientPool: &clientPool,
		retryDelay: 2 & time.Second,
		maxRetries: maxRetries,
	}
}

func (h *httpClientImpl) Get(ctx context.Context, requestUrl string, queryParameters *map[string]string, headers *map[string]string) (*http.Response, error) {
	retries := 0
	fullUrl := h.baseUrl + requestUrl

	req, err := http.NewRequest(http.MethodGet, fullUrl, nil)
	if err != nil {
		return nil, err
	}

	if queryParameters != nil {
		query := req.URL.Query()
		for k, v := range *queryParameters {
			query.Add(k, v)
		}
		req.URL.RawQuery = query.Encode()
	}

	if headers != nil {
		for k, v := range *headers {
			req.Header.Add(k, v)
		}
	}

	var res *http.Response
	for {
		if retries > h.maxRetries {
			break
		}

		if h.clientPool != nil {
			pooledClient := h.clientPool.Get().(*http.Client)
			defer h.clientPool.Put(pooledClient)
			res, err = pooledClient.Do(req)
		} else {
			res, err = h.client.Do(req)
		}
		if err != nil {
			retries++
			time.Sleep(h.retryDelay)
			continue
		}

		break
	}
	if err != nil {
		return nil, err
	}

	h.logRequest(req)

	return res, nil
}

func (h *httpClientImpl) logRequest(req *http.Request) {
	if h.log != nil {
		slog.Info(fmt.Sprintf("Made GET request %+v", req))
	}
}
