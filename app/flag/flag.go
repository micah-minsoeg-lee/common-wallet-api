package flag

import "flag"

var config_flag = flag.String("config", "", "config file path")

type Flags struct {
	configFlag string
}

func Parse() *Flags {
	flag.Parse()

	return &Flags{
		configFlag: *config_flag,
	}
}

func (f *Flags) ConfigFlag() string {
	return f.configFlag
}
