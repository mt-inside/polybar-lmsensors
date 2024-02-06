package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"

	"github.com/mt-inside/go-lmsensors"
)

func main() {
	var opts struct {
		Type       bool `short:"t" long:"type" description:"Print the type of the sensor"`
		TypeSymbol bool `short:"T" long:"type-symbol" description:"Print the type of the sensor, as a unicode symbol (font-dependant)"`
		Unit       bool `short:"u" long:"unit" description:"Print the unit of the reading"`
		Name       bool `short:"n" long:"name" description:"Print the name of the sensor"`
	}
	args, err := flags.Parse(&opts)
	if err != nil {
		panic(err)
	}

	if err := lmsensors.Init(); err != nil {
		panic(err)
	}

	sensors, err := lmsensors.Get()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	output := []string{}
	for _, arg := range args {
		var label string

		labels := strings.Split(arg, "=")
		if len(labels) > 2 {
			fmt.Println("Error: invalid sensor address")
			os.Exit(1)
		}

		addrs := strings.Split(labels[0], "/")
		if len(addrs) > 2 {
			fmt.Println("Error: invalid sensor address")
			os.Exit(1)
		}

		if len(labels) == 2 {
			label = labels[1]
		} else {
			label = addrs[1]
		}

		reading := sensors.Chips[addrs[0]].Sensors[addrs[1]]
		if reading == nil {
			fmt.Println("Error: can't find Sensor")
			os.Exit(1)
		}

		var sb strings.Builder
		if opts.TypeSymbol {
			switch reading.(type) {
			case *lmsensors.TempSensor:
				sb.WriteString(" ")
			case *lmsensors.FanSensor:
				sb.WriteString(" ")
			case *lmsensors.VoltageSensor:
				sb.WriteString("⚡ ")
			default:
				sb.WriteString("TODO ")
			}
		}
		if opts.Type {
			switch reading.(type) {
			case *lmsensors.TempSensor:
				sb.WriteString("temp ")
			case *lmsensors.FanSensor:
				sb.WriteString("fan ")
			case *lmsensors.VoltageSensor:
				sb.WriteString("volt ")
			default:
				sb.WriteString("TODO ")
			}
		}
		if opts.Name {
			sb.WriteString(label + " ")
		}
		sb.WriteString(reading.Rendered())
		if opts.Unit {
			sb.WriteString(reading.Unit())
		}

		output = append(output, sb.String())
	}

	fmt.Print(strings.Join(output, ", "))
}
