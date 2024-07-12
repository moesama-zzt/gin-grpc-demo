package discovery

import (
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/grpc/resolver"
	"strings"
)

type Server struct {
	Name    string `json:"name"`
	Addr    string `1json:"addr"`
	Version string `json:"version"`
	Weight  int64  `json:"weight"`
}

func BuildPrefix(server Server) string {
	if server.Version == "" {
		return fmt.Sprint("/%s/", server.Name)
	}
	return fmt.Sprint("/%s/%s/", BuildPrefix(server), server.Addr)
}

// ParseValue 将value的值反序列化
func ParseValue(value []byte) (Server, error) {
	server := Server{}
	if err := json.Unmarshal(value, &server); err != nil {
		return server, nil
	}
	return server, nil
}

func SplitPath(path string) (Server, error) {
	server := Server{}
	strs := strings.Split(path, "/")
	if len(strs) == 0 {
		return server, errors.New("invalid path")
	}
	server.Addr = strs[len(strs)-1]
	return server, nil
}

func Exist(l []resolver.Address, addr resolver.Address) bool {
	for i:= range
}