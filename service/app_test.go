package service_test

import (
	"errors"
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	"io/ioutil"
	"net/http"
	"time"

	_ "github.com/pari-27/GoPhotos/service"
)

func getResponse(url string) ([]byte, error) {
	if len(url) == 0 {
		return nil, errors.New("Invalid URL")
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	c := &http.Client{
		Timeout: 5 * time.Second,
	}
	fmt.Println(req)
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	code := resp.StatusCode
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil && code != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	if code != http.StatusOK {
		return nil, fmt.Errorf("Server status error: %v", http.StatusText(code))
	}
	return body, nil
}

var _ = Describe("App", func() {
	var (
		server     *ghttp.Server
		statusCode int
		body       []byte
		path       string
		addr       string
	)
	BeforeEach(func() {
		// start a test http server
		server = ghttp.NewServer()
	})
	AfterEach(func() {
		server.Close()
	})
	Context("When get request is sent to albums path", func() {
		BeforeEach(func() {
			statusCode = 200
			path = "/albums"
			addr = "http://" + server.Addr() + path
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", path),
					ghttp.RespondWithPtr(&statusCode, &body),
				))
		})
		It("Returns all the albums", func() {
			_, err := getResponse(addr)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(statusCode).Should(BeEquivalentTo(http.StatusOK))
		})
	})
})
