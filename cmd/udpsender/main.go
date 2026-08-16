package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)
func main(){
	network := "udp"
	address := "localhost:42069"
	udpAddress,err := net.ResolveUDPAddr(network,address)
	if err !=nil{
		log.Fatal("error","error",err)
	}
	udpConnection , err :=net.DialUDP(network,nil,udpAddress)
	if err !=nil{
		log.Fatal("error","error",err)
	}
	defer udpConnection.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		//data := make([]byte,8)
		fmt.Print(">")
		line,err := reader.ReadString('\n')
		if err ==io.EOF{
			return
		}
		if err !=nil{
			log.Fatal("error","error",err)
		}
		_,err = udpConnection.Write([]byte(line))
		if err!=nil{
			log.Fatal("error","error",err)
		}
		//fmt.Printf("bytes escritos: %d",n)


	}



}
