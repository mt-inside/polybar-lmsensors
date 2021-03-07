# polybar-lmsensors
Simple programme that prints the values of the specified lmsensors.

[![Checks](https://github.com/mt-inside/polybar-lmsensors/actions/workflows/checks.yaml/badge.svg)](https://github.com/mt-inside/polybar-lmsensors/actions/workflows/checks.yaml)
[![GitHub Issues](https://img.shields.io/github/issues-raw/mt-inside/polybar-lmsensors)](https://github.com/mt-inside/polybar-lmsensors/issues)

[![Go Reference](https://pkg.go.dev/badge/github.com/mt-inside/polybar-lmsensors.svg)](https://pkg.go.dev/github.com/mt-inside/polybar-lmsensors)

Uses the [lm-sensors](https://github.com/lm-sensors/lm-sensors) (linux monitoring sensors) pacakge, on top of the [hwmon](https://hwmon.wiki.kernel.org) kernel feature.

## Setup
* Install _lm-sensors_
  * Ubuntu: `sudo apt install lm-sensors libsensors-dev`
  * Arch: `pacman -S lm_sensors`
* Configure _lm-sensors_
  * Run `sensors-detect`
  * Make any [necessary adjustments](https://hwmon.wiki.kernel.org/faq) to the [configuration](https://linux.die.net/man/5/sensors3.conf) in `/etc/sensors3.conf`, using `/etc/sensors.d/*`
* Download `polybar-lmsensors` binary from TODO
  * Or, with the _go_ toolchain installed, build it manually with `go get github.com/mt-inside/polybar-lmsensors`

## Example

### Manual Execution
```
$ polybar-lmsensors -t -n -u it8688-isa-0a40/SYS_FAN2=sys k10temp-pci-00c3/Tdie
Fan sys 825/min, Temp Tdie 41°C
```

### Polybar Config
```
[module/sensors-cpu]
type = custom/script
interval = 2
format-prefix = " "
exec = $GOPATH/bin/polybar-lmsensors -u -n -t k10temp-pci-00c3/Tdie it8688-isa-0a40/CPU_FAN=cpu
```
