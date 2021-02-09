package tour

import (
	"fmt"
	"net/http"
)

func Entry() {
	t, err := ContentType("https://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(t)
	}
}

func ContentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	cType := resp.Header.Get("Content-Type")
	if cType == "" {
		return "", fmt.Errorf("cannot find Content-Type header")
	}
	return cType, nil
}
