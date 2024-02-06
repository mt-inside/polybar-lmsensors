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
		Type bool `short:"t" long:"type" description:"Print the type of the sensor"`
		// TODO: lookup go-lmsensor's "sensorType" and translate to , , etc
		TypeSymbol bool `short:"T" long:"type-symbol" description:"Print the type of the sensor, as a unicode symbol (font-dependant)"`
		Unit       bool `short:"u" long:"unit" description:"Print the unit of the reading"`
		Name       bool `short:"n" long:"name" description:"Print the name of the sensor"`
		// TODO: implement
		Fans bool `short:"f" long:"fans" description:"Print all non-zero-speed fans (ignores positional arguments)"`
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
		var sb strings.Builder
		if opts.Type {
			sb.WriteString(reading.SensorType.String() + " ")
		}
		if opts.Name {
			sb.WriteString(label + " ")
		}
		sb.WriteString(reading.Rendered)
		if opts.Unit {
			sb.WriteString(reading.Unit)
		}

		output = append(output, sb.String())
	}

	fmt.Print(strings.Join(output, ", "))
}
