package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"net/url"
)

var addr = flag.String("addr", "192.168.70.141:80", "http service address")

func main() {
	flag.Parse()

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/app"}
	fmt.Println("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println("websocket fail :", err)
		return
	}
	defer c.Close()

	sendStr := `{"m":1001,"p":{"isOnline":0,"hand":false,"cid":1024453,"userid":"141823-15049707","character":3,"audio":false,"videoUrl":null,"audioUrl":null,"audioSynId":"141823-15049707--2016-08-04-14-07-35","videoTime":0,"desktopSynId":"141823-15049707--2016-08-04-14-07-35","name":".........","videoSynId":"141823-15049707--2016-08-04-14-07-35","setB":false,"state":3,"desktopUrl":null,"desktop":false,"video":false}}`
	err = c.WriteMessage(websocket.TextMessage, []byte(sendStr))
	if err != nil {
		fmt.Println("WriteMessage fail:", err)
		return
	}

	for {
		iMsgType, recvBytes, err := c.ReadMessage()
		if err != nil {
			fmt.Println("ReadMessage fail: ", err)
			return
		} else {
			fmt.Println("msgType: ", iMsgType, "read message: ", string(recvBytes))
		}
	}
}
