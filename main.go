package main

import (
	"fmt"
	"github.com/bmamha/gator/internal/config"
)

func main(){
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("error in reading user:%w\n", err)
		return
	}


	err = cfg.SetUser("bereket")
	if err != nil {
		fmt.Printf("error in setting new user:%w\n", err)
		return 
	}
	newCfg, err := config.Read()
	if err != nil {
		fmt.Printf("error in reading new user:%w\n", err)
		return
	}

	fmt.Println(newCfg.DbURL) 
	fmt.Println(newCfg.CurrentUserName)


}
