package alphavantage

import (
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	t.Parallel()

	t.Run("errors with no API key", func(t *testing.T) {
		wantErr := errors.New("no API key provided")
		wantClient := (*Client)(nil)

		gotClient, gotErr := NewClient(ClientOptions{})

		assert.Equal(t, wantErr, gotErr)
		assert.Equal(t, wantClient, gotClient)
	})

	tt := []struct {
		name        string
		giveOptions ClientOptions
		wantClient  Client
	}{
		{
			name: "given no base URL defaults base URL",
			giveOptions: ClientOptions{
				APIKey: "test",
			},
			wantClient: Client{
				apiKey:  "test",
				baseURL: defaultBaseURL,
				httpClient: &http.Client{
					Timeout: defaultTimeout,
				},
			},
		},
		{
			name: "given base URL sets base URL",
			giveOptions: ClientOptions{
				APIKey:  "test",
				BaseURL: "https://example.com/query",
			},
			wantClient: Client{
				apiKey:  "test",
				baseURL: "https://example.com/query",
				httpClient: &http.Client{
					Timeout: defaultTimeout,
				},
			},
		},
		{
			name: "given retries sets retries",
			giveOptions: ClientOptions{
				APIKey:  "test",
				Retries: 1,
			},
			wantClient: Client{
				apiKey:  "test",
				baseURL: defaultBaseURL,
				retries: 1,
				httpClient: &http.Client{
					Timeout: defaultTimeout,
				},
			},
		},
		{
			name: "given HTTP client sets HTTP client",
			giveOptions: ClientOptions{
				APIKey: "test",
				HTTPClient: &http.Client{
					Timeout: 1 * time.Second,
				},
			},
			wantClient: Client{
				apiKey:  "test",
				baseURL: defaultBaseURL,
				httpClient: &http.Client{
					Timeout: 1 * time.Second,
				},
			},
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			wantErr := (error)(nil)

			gotClient, gotErr := NewClient(test.giveOptions)

			assert.Equal(t, wantErr, gotErr)
			assert.Equal(t, test.wantClient, *gotClient)
		})
	}
}

func TestClient_Do(t *testing.T) {
	t.Parallel()

	t.Run("successfully adds api key and datatype", func(t *testing.T) {
		mux := http.NewServeMux()
		server := httptest.NewServer(mux)
		defer server.Close()
		client := &Client{
			apiKey:     "test-123",
			baseURL:    server.URL,
			httpClient: &http.Client{},
		}

		wantQuery := url.Values{}
		wantQuery.Add("apikey", "test-123")
		wantQuery.Add("datatype", "json")
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, wantQuery, r.URL.Query())

			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `{"test":"example"}`)
		})

		gotResponse := make(map[string]interface{})
		gotErr := client.Do(context.TODO(), url.Values{}, &gotResponse)
		assert.Nil(t, gotErr)
		assert.Equal(t, "example", gotResponse["test"])
	})

	t.Run("returns error when receiving error message", func(t *testing.T) {
		mux := http.NewServeMux()
		server := httptest.NewServer(mux)
		defer server.Close()
		client := &Client{
			apiKey:     "test-123",
			baseURL:    server.URL,
			httpClient: &http.Client{},
		}

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `{"Error Message":"failure"}`)
		})

		gotErr := client.Do(context.TODO(), url.Values{}, nil)
		assert.Equal(t, "failure", gotErr.Error())
	})

	t.Run("retries when failing to make request", func(t *testing.T) {
		t.Skip("TODO")
	})

	t.Run("errors when failing to unmarshal", func(t *testing.T) {
		t.Skip("TODO")
	})

	t.Run("cancels when context cancelled", func(t *testing.T) {
		t.Skip("TODO")
	})
}
