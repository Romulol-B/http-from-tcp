package main
import (
	"log"
	"io"
	"fmt"
	"os"
)

func main(){
	//args:= os.Args
	path := "message.txt"
	file,err:= os.Open(path)
	//fmt.Println(args[1])
	if err!= nil{
		log.Fatal("error" ,"error",err)
	}
	for {
		data := make([]byte,8)
		bytesRead ,err:= file.Read(data)
		if err== io.EOF{
			return
		}
		fmt.Printf("read: %s\n",string(data[:bytesRead]))

	}

}
