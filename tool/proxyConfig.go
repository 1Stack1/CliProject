package tool

import (
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"net/url"
	"regexp"
)

/*
*
从配置文件读取代理地址
*/
func proxyConfig(v *viper.Viper) (*http.Transport, error) {
	proxyURL := ConfigReadProxy(v)
	if proxyURL == "" {
		return nil, nil
	}
	if !isValidURL(proxyURL) {
		return nil, fmt.Errorf("配置文件proxy_url不符合url规范")
	}
	proxy, err := url.Parse(proxyURL)
	if err != nil {
		return nil, err
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxy),
	}
	return transport, nil
}

func isValidURL(url string) bool {
	pattern := `^(https?://)?((localhost|[\w-]+(\.[\w-]+)+)|(\d{1,3}\.){3}\d{1,3})(:\d{1,5})?(/[\w./-]*)?(\?[^#\s]*)?(#\S*)?$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(url)
}
