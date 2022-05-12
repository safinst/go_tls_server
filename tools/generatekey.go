package main

import (
	"bufio"
	"log"
	"os"

	"github.com/lucasepe/codename"
)

//produce rand string for req
func main() {
	file, err := os.OpenFile("key.txt", os.O_WRONLY|os.O_CREATE|os.O_CREATE, 0666)
	if err != nil {
		log.Println("文件打开失败", err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	for i := 0; i < 10000; i++ {
		rng, err := codename.DefaultRNG()
		if err != nil {
			log.Println("DefaultRNG:", err.Error())
			return
		}
		key := codename.Generate(rng, 1000)
		write.WriteString(key)
		write.WriteString("\n")
	}
	write.Flush()
}
