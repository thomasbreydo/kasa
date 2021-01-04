package cmd

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/cgreenhalgh/hs1xxplug"
	"github.com/spf13/cobra"
	"github.com/thomasbreydo/kasa/device"
)

func setup() {
	log.SetFlags(0) // no timestamp on logging (e.g. log.Fatal(...))
	flag.Usage = func() {

		flag.Usage()
	}
}

func main() {
	setup()
	var on, off bool
	var ip string
	flag.BoolVar(&on, "on", false, "turn the device on")
	flag.BoolVar(&off, "off", false, "turn the device off")
	flag.StringVar(&ip, "ip", "", "IP address or hostname of the device")
	flag.Parse()
	checkValidInput(on, off, ip)
	plug := hs1xxplug.Hs1xxPlug{IPAddress: ip}
	if on {
		err := device.TurnOn(&plug)
		handleError(err)
	} else if off {
		err := device.TurnOff(&plug)
		handleError(err)
	} else {
		isOff, err := device.IsOff(&plug)
		handleError(err)
		switch isOff {
		case true:
			fmt.Println("The plug is off.")
		case false:
			fmt.Println("The plug is on.")
		}
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkValidInput(on, off bool, ip string) {
	if on && off {
		handleError(errors.New("both --on and --off cannot be set"))
	} else if ip == "" {
		handleError(errors.New("--ip must be set"))
	} else if flag.NArg() != 0 {
		flag.Usage()
		//handleError(errors.New(getUsage()))
	}
}

var DeviceIP string
var rootCmd = &cobra.Command{
	Use:   "kasa",
	Short: "A command-line utility to interact with Kasa smart devices",
	Long:  "A command-line utility to interact with Kasa smart devices.",
}

func init() {
	rootCmd.AddCommand(OnCmd)
	rootCmd.AddCommand(OffCmd)
	rootCmd.AddCommand(StatusCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
