package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/gargath/homecloud/pkg/sis"
)

func main() {
	log.Printf("Homecloud System Information Service %s\n", version())

	viper.SetEnvPrefix("HCSIS")
	viper.AutomaticEnv()

	//	flag.String("listenAddr", "0.0.0.0:8080", "address to listen on")
	flag.Bool("help", false, "print this help and exit")

	flag.Parse()
	viper.BindPFlags(flag.CommandLine)

	if viper.GetBool("help") {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, flag.CommandLine.FlagUsages())
		os.Exit(0)
	}

	if runtime.GOOS != "linux" {
		fmt.Fprintf(os.Stderr, "ERROR: %s requires linux operating system. Yours is %s\n", os.Args[0], runtime.GOOS)
		os.Exit(1)
	}

	info, err := sis.Collect(VERSION)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to collect sys info: %v", err)
		os.Exit(1)
	}
	out, err := json.Marshal(info)
	fmt.Println(string(out))
}
