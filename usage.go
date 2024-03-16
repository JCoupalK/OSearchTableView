package main

import (
	"fmt"
	"os"
)

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "General options:")
	fmt.Fprintln(os.Stderr, "  -u,    --url             	OpenSearch URL")
	fmt.Fprintln(os.Stderr, "  -U,    --user             	OpenSearch user")
	fmt.Fprintln(os.Stderr, "  -p,    --password        	OpenSearch password")
	fmt.Fprintln(os.Stderr, "  -i,    --index             	Index name")
	fmt.Fprintln(os.Stderr, "  -s,    --size         	Size limit for the number of rows to fetch (Default is 10, Maximum is 10000)")
	fmt.Fprintln(os.Stderr, "  -c,    --config         	Config file path (replaces above arguments)")
	fmt.Fprintln(os.Stderr, "")
	os.Exit(1)
}
