package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestContainerStats(t *testing.T) {
	expectedURL := "/containers/dc26aecdcbf8/stats"
	cases := []struct {
		stream         bool
		expectedStream string
	}{
		{
			expectedStream: "0",
		},
		{
			stream:         true,
			expectedStream: "1",
		},
	}
	for _, c := range cases {
		client := &Client{
			client: newMockClient(func(r *http.Request) (*http.Response, error) {
				if !strings.HasPrefix(r.URL.Path, expectedURL) {
					return nil, fmt.Errorf("Expected URL '%s', got '%s'", expectedURL, r.URL)
				}
				query := r.URL.Query()
				stream := query.Get("stream")
				if stream != c.expectedStream {
					return nil, fmt.Errorf("stream not set in URL query properly. Expected '%s', got %s", c.expectedStream, stream)
				}

				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewReader([]byte("response"))),
				}, nil
			}),
		}
		resp, err := client.ContainerStats(context.Background(), "dc26aecdcbf8", c.stream)

		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()
		content, err := io.ReadAll(resp.Body)
		contentstr := string(content)
		println(contentstr)
		if err != nil {
			t.Fatal(err)
		}
		if string(content) != "response" {
			t.Fatalf("expected response to contain 'response', got %s", string(content))
		}

	}
}
