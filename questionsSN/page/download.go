package page

import (
	"context"
	"io"
	"net/http"
)

func DownloadPage(url string, out io.Writer) error {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp.Write(out)
	return nil
}
