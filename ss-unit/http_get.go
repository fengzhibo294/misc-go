package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	//生成client 参数为默认

	client := &http.Client{}

	//生成要访问的url

	//url := "http://" + "huiyi.cnstrongwx.cn" + "/index.php/user/update_partner_role/" + "1223006" + "/1"
	url := "http://192.168.20.39:8080"

	//提交请求

	reqest, err := http.NewRequest("POST", url, strings.NewReader("username=kevin&password=*********"))

	if err != nil {

		panic(err)

	}

	reqest.Header.Set("Accept-Charset", "8859_1,utf-8;q=0.7,*;q=0.3")
	//处理返回结果

	response, _ := client.Do(reqest)

	//将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理

	stdout := os.Stdout

	_, err = io.Copy(stdout, response.Body)

	//返回的状态码

	status := response.StatusCode

	fmt.Println(status)

}
