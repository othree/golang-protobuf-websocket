package main;

import (
  "log"
  "fmt"
  "net/http"
  "github.com/golang/protobuf/proto"
  "github.com/gorilla/websocket"
  "./message"
)

var data []byte

var upgrader = websocket.Upgrader{
  ReadBufferSize:  1024,
  WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
  conn, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
    return
  }
  for {
    _, _, err := conn.ReadMessage()
    if err != nil {
      return
    }
    fmt.Println("read done")
    fmt.Println("write %s", data)
    if err = conn.WriteMessage(websocket.BinaryMessage, data); err != nil {
      fmt.Println("err %s", err)
      return
    }
    fmt.Println("write done")
  }
}

func serveHome(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.Error(w, "Not found", 404)
    return
  }
  if r.Method != "GET" {
    http.Error(w, "Method not allowed", 405)
    return
  }
  w.Header().Set("Content-Type", "text/html; charset=utf-8")
  w.Write([]byte(homeHTML))
}


func main() {

  msg := &message.Message{
    Id: proto.Int32(17),
    Author: &message.Message_Person{
      Id: proto.Int32(1),
      Name:  proto.String("othree"),
    },
    Text:  proto.String("Hi, this is message."),
  }


  fmt.Println(msg.GetAuthor().GetName() + ": " +msg.GetText())

  data, _ = proto.Marshal(msg)

  fmt.Println("%s", data)

  http.HandleFunc("/", serveHome)
  http.HandleFunc("/ws", handler)
  if err := http.ListenAndServe("127.0.0.1:1337", nil); err != nil {
    log.Fatal(err)
  }
}

const homeHTML = `<!DOCTYPE html>
<html lang="en">
    <head>
        <title>WebSocket Example</title>
    </head>
    <body>
        <script type="text/javascript">
            (function() {
                var data = document.getElementById("fileData");
                var conn = new WebSocket("ws://127.0.0.1:1337/ws");
                conn.onopen = function () {
                  console.log("Opening a connection...");
                  conn.send('MSG');
                  conn.binaryType = "arraybuffer";
                }
                conn.onclose = function(evt) {
                  data.textContent = 'Connection closed';
                }
                conn.onmessage = function(evt) {
                    console.log('file updated');
                    data.textContent = evt.data;
                }
            })();
        </script>
    </body>
</html>
`
