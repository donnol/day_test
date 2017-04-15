package main

import (
	"fmt"
	"regexp"
)

func CheckQiniuUrl(link string) bool {
	// pattern := `^http(s)?://[image.hongbeibang.com/|video.hongbeibang.com/|audio.hongbeibang.com/]{1}`
	pattern := `[http:\/\/|https:\/\/]*www[\.test]*[2]*\.hongbeibang\.com\/live\/[lesson|room|series]+\/\d+`
	re := regexp.MustCompile(pattern)
	return re.MatchString(link)
}

func main() {
	link := []string{
		"link:http://image.hongbeibang.com/FqjyQAWyk6utJqnJV1lIz39lEtfD?1080X1440",
		"http://image.hongbeibang.com/FqjyQAWyk6utJqnJV1lIz39lEtfD?1080X1440",
		"https://image.hongbeibang.com/FqjyQAWyk6utJqnJV1lIz39lEtfD&1080X1440",
		"httpss://image.hongbeibang.com/FqjyQAWyk6utJqnJV1lIz39lEtfD&1080X1440",
		"http://video.hongbeibang.com/FgLtiBPZ4VTk39KMcfnJkonwDuh0?attname=&e=1489731444&token=M8SB8lTux2QlyQ3Yy44MNGRRXwJUccuFJkMzJ0SJ:XITdpBUl56JPbMp7b9UE-9UTLYA",
		"https://video.hongbeibang.com/FgLtiBPZ4VTk39KMcfnJkonwDuh0?attname=&e=1489731444&token=M8SB8lTux2QlyQ3Yy44MNGRRXwJUccuFJkMzJ0SJ:XITdpBUl56JPbMp7b9UE-9UTLYA",
		"httpss://video.hongbeibang.com/FgLtiBPZ4VTk39KMcfnJkonwDuh0?attname=&e=1489731444&token=M8SB8lTux2QlyQ3Yy44MNGRRXwJUccuFJkMzJ0SJ:XITdpBUl56JPbMp7b9UE-9UTLYA",
		"http://audio.hongbeibang.com/Fg-oGliOcY3c4CuuXdaD8Kflmgct?attname=&e=1489731482&token=M8SB8lTux2QlyQ3Yy44MNGRRXwJUccuFJkMzJ0SJ:QZemzvmj12hLrRRbWuEJVQaGUvE",
		"https://audio.hongbeibang.com/Fg-oGliOcY3c4CuuXdaD8Kflmgct?attname=&e=1489731482&token=M8SB8lTux2QlyQ3Yy44MNGRRXwJUccuFJkMzJ0SJ:QZemzvmj12hLrRRbWuEJVQaGUvE",
		"httpss://audio.hongbeibang.com/Fg-oGliOcY3c4CuuXdaD8Kflmgct?attname=&e=1489731482&token=M8SB8lTux2QlyQ3Yy44MNGRRXwJUccuFJkMzJ0SJ:QZemzvmj12hLrRRbWuEJVQaGUvE",
		"https://v.qq.com/iframe/player.html?vid=m035285ws08&tiny=0&auto=0",
		"http://v.qq.com/iframe/player.html?vid=m035285ws08&tiny=0&auto=0",
		"abcc",
	}

	for _, single := range link {
		fmt.Println(single, "is", CheckQiniuUrl(single))
	}
}
