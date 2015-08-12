package main;

import (
  "fmt"
  "github.com/golang/protobuf/proto"
  "./message"
)

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

  data, _ := proto.Marshal(msg)

  fmt.Println("%s", data)
}
