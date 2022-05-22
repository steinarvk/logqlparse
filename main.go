package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/grafana/loki/pkg/logql/syntax"
)

type parseFailed struct {
	query string
	err   error
}

func (e parseFailed) Error() string {
	return fmt.Sprintf("Parse failed as Loki LokQL expression:\n\t%s\n\nExpression was:\n\t%s", e.err.Error(), e.query)
}

func evaluate(query string) (string, error) {
	query = strings.TrimSpace(query)

	parsed, err := syntax.ParseExpr(query)
	if err == nil {
		canonical := parsed.String()
		return canonical, nil
	}

	return "", parseFailed{query, err}
}

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read error: %v", err)
		os.Exit(1)
	}

	if _, err := evaluate(string(data)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
