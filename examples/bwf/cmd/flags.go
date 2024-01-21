package main

import "flag"

var envFlag = flag.String("env", "", ".env file to load")

var appFlag StringArrFlag

type StringArrFlag []string

func (i *StringArrFlag) String() string {
	return "application(s) to run"
}

func (i *StringArrFlag) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func (i *StringArrFlag) Value() interface{} {
	return *i
}
