
<div align="center">
  <h1><strong>gopci</strong></h1>
  <p>
		<strong>A minimal and fast library to parse pci device info from sysfs.</strong>
  </p>
  <p>
    <a href="https://goreportcard.com/report/github.com/hertg/gopci">
      <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/hertg/gopci" />
    </a>
    <a href="#">
			<img alt="License Information" src="https://img.shields.io/github/license/hertg/go-pciids">
    </a>
  </p>
</div>

## Description

This library provides methods to get PCI information on a linux system.
It has been created as a minimalistic and more performant alternative to
the pci parsing capabilities of [jaypipes/ghw](https://github.com/jaypipes/ghw).

If you need non-pci or non-linux hardware information,
ghw offers an all-in-one solution. If you simply care about linux pci information,
this library offers a **significant performance benefit** over ghw (see [comparison](#Comparison)).

## Usage

### Install
```shell
go get github.com/hertg/gopci
```

### Scan
```go
devices, _ := pci.Scan()
for _, device := range devices {
	// ...
}
```

The `Device` struct contains information like `Vendor`, `Device`, `Class`, `Subvendor`, `Subdevice`, etc.
More detailed information is made available in the `Config` field which contains the raw parsed PCI config header.

If more information from sysfs shall be processed, the `SysfsPath()` method can be used
to get a direct link to the PCI devices sysfs path which can be used for further custom information gathering.

### Filter
The `Scan()` method allows optional `Filter` arguments to only include matching
devices in the resulting list of devices.

```go
classFilter := func(d *pci.Device) bool { return d.Class.Class == 0x03 }
devices, _ := pci.Scan(classFilter)
```

## Comparison

Device configuration is parsed directly as bytes (from `config`) instead of
reading strings and parsing from there. This prevents unnecessary string
allocationis and significantly improves performance.

It has been found that this library parses at more
than 10x the speed while using 50x less memory compared to
[jaypipes/ghw](https://github.com/jaypipes/ghw).

```text
goos: linux
goarch: amd64
pkg: github.com/hertg/gopci/pkg/pci
cpu: AMD Ryzen 9 5950X 16-Core Processor
BenchmarkGoPci
BenchmarkGoPci-32            518           2554827 ns/op          297709 B/op       5178 allocs/op
BenchmarkGhw
BenchmarkGhw-32               34          33286659 ns/op        15745188 B/op     201189 allocs/op
PASS
ok      github.com/hertg/gopci/pkg/pci  2.802s
```

