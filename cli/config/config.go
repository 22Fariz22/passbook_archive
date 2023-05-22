package config

import (
	"path"
)

// UserFilePath путь где будет появлятся файл session.txt с сессионным ключом
var UserFilePath = path.Join("/")

// Key секретный ключ для шифратора
var Key = []byte("passphrasewhichneedstobe32bytes!")

//type Config struct {
//	Text string
//}
//
//func NewConfig() *Config {
//	cfg := &Config{}
//
//	pflag.StringVarP(&cfg.Text, "text", "", "", "file storage path")
//
//}
