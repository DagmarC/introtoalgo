package page

import (
	"bytes"
	"testing"
)

func TestDownloadPage(t *testing.T) {
	var buf bytes.Buffer
	DownloadPage("https://www.google.com", &buf)

	if len(buf.Bytes()) == 0 {
		t.Errorf("No write has been done.")
	}
	if !bytes.Contains(buf.Bytes(), []byte("HTTP/2.0 200 OK")) {
		t.Errorf("Write has been done but not the successful one")
	}
}
