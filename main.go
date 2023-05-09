package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	//在同目录下创建prop.env文件("PORT=8088")
	err := godotenv.Load("prop.env")
	if err != nil {
		log.Fatal(err)
	}

	//构建创世块genesisBlock
	t := time.Now()
	genesisBlock := Block{}
	genesisBlock = Block{0, t.String(), 0, CalculateBlockHash(genesisBlock), "", ""}
	spew.Dump(genesisBlock)

	Blockchain = append(Blockchain, genesisBlock)
	//读取.env文件，获取Server端口8088
	httpPort := os.Getenv("PORT")
	//监听server
	server, err := net.Listen("tcp", ":"+httpPort)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("HTTP Server Listening on port : ", httpPort)
	defer server.Close()

	go func() {
		for candidate := range candidateBlocks {
			mutex.Lock()
			tempBlocks = append(tempBlocks, candidate)
			mutex.Unlock()
		}
	}()

	go func() {
		for {
			PickWinner() //选举winner
		}
	}()

	for {
		conn, err := server.Accept() //开启服务
		if err != nil {
			log.Fatal(err)
		}
		go HandleConn(conn) //处理连接
	}

}