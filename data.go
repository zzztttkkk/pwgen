package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Data struct {
	Default  string
	Accounts map[string]string
}

func load() *Data {
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	f, err := os.Open(filepath.Join(dir, ".pwgen"))
	if err != nil {
		if os.IsNotExist(err) {
			return &Data{Accounts: map[string]string{}}
		}
		log.Fatalln(err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}

	v, err := AesDecrypt(data)
	if err != nil {
		log.Fatalln(err)
	}

	m := &Data{}
	err = json.Unmarshal(v, &m)
	if err != nil {
		log.Fatalln(err)
	}
	return m
}

func save(m *Data) {
	data, err := json.Marshal(m)
	if err != nil {
		log.Fatalln(err)
	}
	data, err = AesEncrypt(data)
	if err != nil {
		log.Fatalln(err)
	}
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.OpenFile(filepath.Join(dir, ".pwgen"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		log.Fatalln(err)
	}
}

func getKey(name string) string {
	accounts := load().Accounts
	if len(accounts) > 0 {
		return accounts[name]
	}
	return ""
}

func setKey(name, secret string, isDefault bool) {
	m := load()
	m.Accounts[name] = secret
	if isDefault {
		m.Default = name
	}
	save(m)
}

func delKey(name string) {
	m := load()
	delete(m.Accounts, name)
	if m.Default == name {
		m.Default = ""
	}
	save(m)
}
