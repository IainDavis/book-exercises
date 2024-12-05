package main

type Logic interface {
	SayHello(userId string) (string, error)
}
