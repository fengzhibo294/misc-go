package main

import (
	"fmt"
	sjson "github.com/bitly/go-simplejson"
)

func main() {
	str := `{"m":1001,"p":{"isOnline":0,"hand":false,"cid":1024453,"userid":"141823-15049707","character":3,"audio":false,"videoUrl":null,"audioUrl":null,"audioSynId":"141823-15049707--2016-08-04-14-07-35","videoTime":0,"desktopSynId":"141823-15049707--2016-08-04-14-07-35","name":".........","videoSynId":"141823-15049707--2016-08-04-14-07-35","setB":false,"state":3,"desktopUrl":null,"desktop":false,"video":false}}`

	strJs, _ := sjson.NewJson([]byte(str))
	fmt.Println("str: ", strJs)

	bytes, _ := strJs.MarshalJSON()
	fmt.Println("after str:", string(by))
}
