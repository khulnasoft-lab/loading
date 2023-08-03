package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/khulnasoft-labs/loading/pkg/bar"
)

func main() {

	fmt.Println("Demo: Customised download bar")

	// create the bar
	loadingBar := bar.New(
		bar.OptionWithStatsFuncs(
			bar.StatsBytesComplete,
		),
	)

	// download a file
	resp, _ := http.DefaultClient.Get("http://speedtest.ftp.otenet.gr/files/test10Mb.db")
	defer func() { _ = resp.Body.Close() }()

	// ...and create a temp file to save it to locally
	f, _ := ioutil.TempFile(os.TempDir(), "loading-example")

	// tell the loading bar how big the total download is
	loadingBar.SetTotalInt64(resp.ContentLength)

	// download the data to the temporary file, and update the loading bar along the way
	_, _ = io.Copy(io.MultiWriter(f, loadingBar), resp.Body)

}
