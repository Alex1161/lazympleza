package decisionTree

import (
	"fmt"
	"strconv"
	"strings"
)

type Request struct {
	request string
	params  []string
}

func CreateRequest(request string, params []string) Request {
	return Request{request, params}
}

func (r Request) ToString() string {
	params := strings.Join(r.params[:], ",")
	return strings.Join([]string{"request", r.request, fmt.Sprint(len(r.params)), params}, ",")
}

func ToRequest(content string) Request {
	l := strings.Split(content, ",")
	request := l[1]
	offset, _ := strconv.Atoi(l[2])
	params := l[3 : 3+offset]

	return Request{request, params}
}
