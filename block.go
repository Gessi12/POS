package main

import "sync"

//构建块类型
type Block struct {
	Index   int // 索引
	Timestamp  string //时间
	BPM int //脉搏
	Hash string //哈希
	PrevHash string // 前一个哈希
	Validator string //验证器
}



var (
	Blockchain []Block                 //定义区块
	tempBlocks []Block                 //临时区块
	candidateBlocks = make(chan Block) //每个新区块发送通道
	announcements = make(chan string)  //TCP Server向所有节点广播最新区块
	validators = make(map[string]int)  //使用map存储验证的token
)

var mutex = &sync.Mutex{} // 用于并发读写的锁
