package main

import (
	"fmt"
	"regexp"
)

func CheckQiniuUrl(link string) bool {
	// pattern := `[http|https]://[image|video|audio][.]hongbeibang[.]com/.*?`
	// pattern := "[http|https]://[image|video|audio][\\.]hongbeibang"
	pattern := `^http(s*)://([image.hongbeibang.com/|video.hongbeibang.com/|audio.hongbeibang.com/]+).*`
	re := regexp.MustCompile(pattern)
	return re.MatchString(link)
}

func main() {
	link := []string{
		"link:http://image.hongbeibang.com/FqjyQAWyk6utJqnJV1lIz39lEtfD?1080X1440",
		"http://image.hongbeibang.com/FqjyQAWyk6utJqnJV1lIz39lEtfD?1080X1440",
		"http://image.hongbeibang.com/FqjyQAWyk6utJqnJV1lIz39lEtfD&1080X1440",
		"https://image.hongbeibang.com/FqjyQAWyk6utJqnJV1lIz39lEtfD&1080X1440",
		"https://v.qq.com/iframe/player.html?vid=m035285ws08&tiny=0&auto=0",
		"http://v.qq.com/iframe/player.html?vid=m035285ws08&tiny=0&auto=0",
		"abcc",
	}

	for _, single := range link {
		fmt.Println(CheckQiniuUrl(single))
	}
}
