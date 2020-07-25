package agent

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

//noinspection GoUnhandledErrorResult
func NewReverseProxy(host, token string) http.Handler {
	proxy := &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			r.URL.Scheme = "https"
			r.URL.Host = "npm.pkg.github.com"
			r.Host = r.URL.Host
			r.Header.Set("Authorization", "Bearer "+token)
		},
		ModifyResponse: func(r *http.Response) (err error) {
			defer r.Body.Close()
			data, err := ioutil.ReadAll(r.Body)
			if err != nil {
				return
			}
			data = bytes.ReplaceAll(
				data,
				[]byte("https://npm.pkg.github.com"),
				[]byte(host),
			)
			r.Body = ioutil.NopCloser(bytes.NewReader(data))
			return
		},
	}
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(rw, "Not Allowed", http.StatusMethodNotAllowed)
		} else {
			proxy.ServeHTTP(rw, r)
		}
	})
}
