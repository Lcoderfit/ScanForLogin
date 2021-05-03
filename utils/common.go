package utils

import (
	"net"
	"net/url"
	"path"
)

// UrlJoin 对URL进行拼接
func UrlJoin(paths ...string) (string, error) {
	baseUrl, err := url.Parse(paths[0])
	if err != nil {
		return "", err
	}
	baseUrl.Path = path.Join(paths[1:]...)
	return baseUrl.String(), nil
}

// 获取本地IP(内网IP)地址
func GetIntranetIp() string {
	addressList, err := net.InterfaceAddrs()
	if err != nil {
		Logger.Error("IP地址获取失败, ", err)
	}

	var results []string
	for _, address := range addressList {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				results = append(results, ipNet.IP.String())
			}
		}
	}
	if len(results) == 0 {
		return ""
	}
	return results[len(results)-2]
}
