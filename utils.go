package main

import "github.com/fatih/color"

func print_good(msg string){
	color.Green("[+] - %s", msg)
}

func print_info(msg string){
	color.Cyan("[*] - %s", msg)
}

func panic_with_msg(msg string, err error) {
	panic(color.RedString(msg) + "\n" + err.Error())
}