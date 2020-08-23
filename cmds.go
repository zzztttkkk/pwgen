package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"pwgen/clipboard"
)

func cmdAdd() {
	var set = flag.NewFlagSet("", flag.ExitOnError)
	var unPtr = set.String("u", "", "username")
	var idPtr = set.Bool("d", false, "make default")
	var i int
	var v string
	for i, v = range os.Args {
		if v[0] == '-' {
			break
		}
	}
	err := set.Parse(os.Args[i:])
	if err != nil {
		log.Fatalln(err)
	}

	var username string
	if len(os.Args) >= 3 && noOptBefore(2) { // pwgen add username ...
		username = os.Args[2]
	}
	if len(username) < 1 || username[0] == '-' {
		username = *unPtr
		if len(username) < 1 {
			username = _GetUsername(false)
		}
	}

	s1 := _GetSecret(false)
	s2 := _GetSecret(true)
	if s1 != s2 {
		log.Fatalln("the two secret keys are inconsistent")
	}
	setKey(username, s1, *idPtr)
}

func cmdDel() {
	var set = flag.NewFlagSet("", flag.ExitOnError)
	var unPtr = set.String("u", "", "username")

	var username string
	if len(os.Args) >= 3 && noOptBefore(2) { // pwgen del username ...
		username = os.Args[2]
	}
	if len(username) < 1 || username[0] == '-' {
		username = *unPtr
		if len(username) < 1 {
			username = _GetUsername(false)
		}
	}
	delKey(username)
}

func cmdLs() {
	d := load()
	if len(d.Accounts) < 1 {
		fmt.Println("nil data")
	}

	for k, _ := range d.Accounts {
		if k == d.Default {
			fmt.Printf("\t%s default\n", k)
		} else {
			fmt.Printf("\t%s\n", k)
		}
	}
}

func noOptBefore(ind int) bool {
	for i := 0; i <= ind; i++ {
		v := os.Args[i]
		if v[0] == '-' {
			return false
		}
	}
	return true
}

func cmdGen() {
	var set = flag.NewFlagSet("gen", flag.ExitOnError)
	var lengthPtr = set.Int("l", 0, "password length(in range 1-64)")
	var dpPtr = set.Bool("p", false, "print password")
	var unPtr = set.String("u", "", "username")
	var hnPtr = set.String("h", "", "hostname")
	var autoSavePtr = set.Bool("s", false, "auto save this account")
	var isDefaultPtr = set.Bool("d", false, "make default")
	var i int
	var v string
	for i, v = range os.Args {
		if v[0] == '-' {
			break
		}
	}
	err := set.Parse(os.Args[i:])
	if err != nil {
		log.Fatalln(err)
	}

	var hostname string
	if len(os.Args) >= 3 && noOptBefore(2) { // pwgen gen hostname ...
		hostname = os.Args[2]
	}
	if len(hostname) < 1 || hostname[0] == '-' {
		hostname = *hnPtr
		if len(hostname) < 1 {
			hostname = _GetHostname()
		}
	}

	var username string
	if len(os.Args) >= 4 && noOptBefore(3) { // pwgen gen hostname username ...
		username = os.Args[3]
	}
	if len(username) < 1 || username[0] == '-' {
		username = *unPtr
		if len(username) < 1 {
			username = _GetUsername(true)
		}
	}

	key := getKey(username)
	if len(key) < 1 {
		if *autoSavePtr {
			s1 := _GetSecret(false)
			s2 := _GetSecret(true)
			if s1 != s2 {
				log.Fatalln("the two secret keys are inconsistent")
			}
			setKey(username, s1, *isDefaultPtr)
			key = s1
		} else {
			key = _GetSecret(false)
		}
	}

	l := *lengthPtr
	if l == 0 {
		l = getLength(hostname)
	} else {
		setLength(hostname, l)
	}
	pwd := generate(hostname, username, key, l)
	if *dpPtr {
		fmt.Println(pwd)
		return
	}
	err = clipboard.WriteAll(pwd)
	if err != nil {
		log.Println("failed to access clipboard, use `-p` to print password")
		return
	}
	log.Println("the password in your clipboard now")
}
