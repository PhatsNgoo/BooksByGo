package Config

import "fmt"

type Config struct {
	serverName string
	userName string
	password string
	databaseName string
}

func New(serverName string,userName string,password string,databaseName string) Config{
	e := Config{serverName,userName,password,databaseName}
	return e
}
func (c Config) GetConfig() {
	fmt.Printf("Connect to server %s with username %s ,password %s to database %s \n",c.serverName,c.userName,c.password,c.databaseName)
}

