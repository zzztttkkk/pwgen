package main

import (
	"bufio"
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
	"syscall"
	"time"
)

func _ReadFromTerminal(name string, isRepeat bool) string {
	if !isRepeat {
		fmt.Printf("Enter Your %s:\n", name)
	} else {
		fmt.Printf("Enter Your %s Again:\n", name)
	}
	pwd, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println()
		log.Fatalf("%s; please see: https://github.com/golang/go/issues/11914#issuecomment-613715787\n", err)
	}
	return string(pwd)
}

func _GetAccountSecret(isRepeat bool) string {
	return _ReadFromTerminal("Account Secret", isRepeat)
}

func _GetAppSecret(isRepeat bool) string {
	return _ReadFromTerminal("Application Secret", isRepeat)
}

func _GetUsername(useData bool) string {
	if useData {
		data := load()
		if len(data.Default) > 0 {
			return data.Default
		}
	}

	fmt.Println("Enter Username:")
	reader := bufio.NewReader(os.Stdin)
	name, _, err := reader.ReadLine()
	if err != nil {
		log.Fatalln(err)
	}

	name = bytes.TrimSpace(name)
	if len(name) < 1 {
		log.Fatalln("empty username")
	}
	return string(name)
}

func _GetHostname() string {
	fmt.Println("Enter Hostname:")
	reader := bufio.NewReader(os.Stdin)
	name, _, _ := reader.ReadLine()
	name = bytes.TrimSpace(name)
	if len(name) < 1 {
		log.Fatalln("empty hostname")
	}
	return string(name)
}

func main() {
	fixSecretKey()
	time.Sleep(time.Millisecond)

	if len(os.Args) < 2 {
		os.Args = append(os.Args, "gen")
	} else {
		sub := os.Args[1]
		if sub[0] == '-' {
			var v = make([]string, len(os.Args)-1, len(os.Args)-1)
			copy(v, os.Args[1:])
			os.Args = os.Args[:1]
			os.Args = append(os.Args, "gen")
			os.Args = append(os.Args, v...)
		}
	}

	switch os.Args[1] {
	case "gen":
		cmdGen()
	case "add":
		cmdAdd()
	case "del":
		cmdDel()
	case "ls":
		cmdLs()
	case "help":
		fmt.Print(
			`pwgen: generate password by remap(sha512(username:secret@hostname)).
        -u : username
        -h : hostname

    add <username string> <-d bool> <-u string>
        save a account to the filesystem("~/.pwgen").
        -d : this account will be the default
	
    del <username string> <-u string>
        delete a account from the filesystem.

    ls
        list all saved accounts.

    gen <hostname string> <username string> <-l int> <-p bool> <-u string> <-h string> <-s bool> <-d bool>
        generate a password. if 'username' is empty, use the default account. this is the default sub-command.
        -l : length of the password, in range (1-64)
        -p : do not send to the clipboard, only print. 
        -s : auto save this account.
        -d : auto save this account as the default.
`,
		)
	}
}
