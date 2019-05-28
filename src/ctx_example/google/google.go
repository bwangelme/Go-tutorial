package google

import (
	"context"
	"ctx_example/userip"
	"fmt"
	"io"
	"net/http"

	"encoding/json"
)

type Result struct {
	Title, URL string
}

// Console: https://developers.google.com/apis-explorer/#p/customsearch/v1/search.cse.list?q=%25E7%259F%25A5%25E4%25B9%258E&cx=017618265042162313223%253Ax79ncj1o0jk&_h=2&
// Doc: https://developers.google.com/custom-search/v1/using_rest
const searchApiKey = ""
const searchEngineId = ""

var url = fmt.Sprintf("https://www.googleapis.com/customsearch/v1?key=%s&cx=%s&fields=kind,items(title,link,snippet,pagemap(cse_image))", searchApiKey, searchEngineId)

type Results []Result

func DecodeResp(reader io.ReadCloser) (results Results, err error) {
	defer reader.Close()
	decoder := json.NewDecoder(reader)
	var data struct {
		Kind  string `json:"kind"`
		Items []struct {
			Title   string `json:"title"`
			Link    string `json:"link"`
			Snippet string `json:"snippet"`
			PageMap struct {
				CSEImage []struct {
					Src string `json:"src"`
				} `json:"cse_image"`
			} `json:"pagemap"`
		} `json:"items"`
	}

	if err = decoder.Decode(&data); err != nil {
		return results, err
	}
	for _, res := range data.Items {
		results = append(results, Result{
			Title: res.Title,
			URL:   res.Link,
		})
	}

	return results, nil
}

func Search(ctx context.Context, query string) (Results, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("q", query)

	if userIP, ok := userip.FromContext(ctx); ok {
		q.Set("userIp", userIP.String())
	}
	req.URL.RawQuery = q.Encode()

	var results Results
	err = httpDo(ctx, req, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}

		// Resp Template: https://gist.github.com/bwangelme/788cc61e135aa642836675fc8de8230e
		results, err = DecodeResp(resp.Body)
		if err != nil {
			return err
		}

		return nil
	})

	return results, err
}

func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
	c := make(chan error, 1)
	req = req.WithContext(ctx)

	go func() {
		c <- f(http.DefaultClient.Do(req))
	}()

	select {
	case <-ctx.Done():
		<-c
		return ctx.Err()
	case err := <-c:
		return err
	}
}
