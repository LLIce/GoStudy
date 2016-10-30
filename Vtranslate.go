package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/axgle/mahonia"
	"github.com/pquerna/ffjson/ffjson"
)

func main() {
	str := ""
	for i := 1; i < len(os.Args); i++ {
		str += os.Args[i] + " "
	}
	jsonParse(get(str))
}

func printjson(dat map[string]interface{}) {
	for k, v := range dat {
		switch k {
		case "translation":
			fmt.Println("translation:", v)
		case "basic":
			switch vv := v.(type) {
			case map[string]interface{}:
				printjson(vv)
			}
		case "query":
			fmt.Println("query:", v)
		case "explains":
			fmt.Println("explains:", v)
		default:
		}
	}
}

func jsonParse(body []byte) {
	dat := make(map[string]interface{})
	if err := ffjson.Unmarshal(body, &dat); err == nil {
		printjson(dat)
	}
}

func get(strn string) []byte {
	enc := mahonia.NewEncoder("utf8")
	str := "http://fanyi.youdao.com/openapi.do?keyfrom=mlovez-dev&key=1341364669&type=data&doctype=json&version=1.1&q=" + enc.ConvertString(strn)

	//进行encode编码
	u, _ := url.Parse(str)
	q := u.Query()
	u.RawQuery = q.Encode()
	response, _ := http.Get(u.String())
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	if response.StatusCode != 200 {
		fmt.Println("error")
	}
	return body
}
