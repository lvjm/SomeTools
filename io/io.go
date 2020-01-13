package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func ReadFrom(reader io.Reader) {

}
func main() {
	// 扫描链码文件
	files, _ := ioutil.ReadDir("./io/chaincode")
	for _, f := range files {
		fi, err := os.Open(fmt.Sprintf("/Users/may/go/src/jm/sometools/io/chaincode/%s",f.Name()))
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		defer fi.Close()
		br := bufio.NewReader(fi)
		for {
			a, _, c := br.ReadLine()
			if c == io.EOF {
				break
			}
			var lines []string
			if strings.Contains(string(a),"function =="){
				lines = append(lines,string(a))
			}
			for _,line := range lines{
				rule,_ := regexp.Compile(`"([^"]+)"`)
				results := rule.FindAllString(line,-1)
				fmt.Println(results[0])
			}
		}
	}
}
