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
	args, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		panic(err)
	}

	sensors, err := lmsensors.Get(true)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	output := []string{}
	for _, arg := range args[1:] {
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

		reading := sensors.ChipsMap[addrs[0]].SensorsMap[addrs[1]]
		var sb strings.Builder
		if opts.Type {
			sb.WriteString(reading.SensorType.String() + " ")
		}
		if opts.Name {
			sb.WriteString(label + " ")
		}
		sb.WriteString(reading.Value)
		if opts.Unit {
			sb.WriteString(reading.Unit)
		}

		output = append(output, sb.String())
	}

	fmt.Print(strings.Join(output, ", "))
}
