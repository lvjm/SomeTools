package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func fetch (url string) string {
	fmt.Println("Fetch Url", url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Http get err:", err)
		return ""
	}
	if resp.StatusCode != 200 {
		fmt.Println("Http status code:", resp.StatusCode)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error", err)
		return ""
	}
	return string(body)
}

func parseUrls(url string) {
	body := fetch(url)
	body = strings.Replace(body, "\n", "", -1)
	rp := regexp.MustCompile(`<div class="hd">(.*?)</div>`)
	titleRe := regexp.MustCompile(`<span class="title">(.*?)</span>`)
	idRe := regexp.MustCompile(`<a href="https://movie.douban.com/subject/(\d+)/"`)
	items := rp.FindAllStringSubmatch(body, -1)
	for _, item := range items {
		fmt.Println(idRe.FindStringSubmatch(item[1])[1],
			titleRe.FindStringSubmatch(item[1])[1])
	}
}

func main()  {
		start := time.Now()
		for i := 0; i < 10; i++ {
			parseUrls("https://movie.douban.com/top250?start=" + strconv.Itoa(25 * i))
		}
		elapsed := time.Since(start)
		fmt.Printf("Took %s", elapsed)
	}


/*
输出结果
Fetch Url https://movie.douban.com/top250?start=0
1292052 肖申克的救赎
1291546 霸王别姬
1292720 阿甘正传
1295644 这个杀手不太冷
1292063 美丽人生
1292722 泰坦尼克号
1291561 千与千寻
1295124 辛德勒的名单
3541415 盗梦空间
3011091 忠犬八公的故事
1292001 海上钢琴师
2131459 机器人总动员
3793023 三傻大闹宝莱坞
1292064 楚门的世界
1291549 放牛班的春天
1889243 星际穿越
1292213 大话西游之大圣娶亲
*/



