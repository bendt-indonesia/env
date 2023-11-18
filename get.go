// This package forked from godotenv (https://github.com/joho/godotenv)
// Commit: c40e9c6392b05ba58e6fea50091ce35a1ef020e7
//
// Examples/readme can be found on the github page at https://github.com/joho/godotenv
//
// The TL;DR is that you make a .env file that looks something like
//
// 		SOME_ENV_VAR=somevalue
//
// and then in your go code you can call
//
// 		env.Load()
//
// and all the env vars declared in .env will be available through os.Getenv("SOME_ENV_VAR")
package env

import (
	"os"
)

func GetEnvConfig(key string) string {
	env := os.Getenv("APP_ENV")
	postfix := ""
	if env == "PROD" || env == "STAGE" {
		postfix += "_"+env
	}

	return os.Getenv(key+postfix)
}

func Get(key string) string {
	return GetEnvConfig(key)
}

