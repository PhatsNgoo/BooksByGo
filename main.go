package main

import "./Config"
func main() {
	c:=Config.New("localhost","root","123456","bookseller")
	c.GetConfig()
}
