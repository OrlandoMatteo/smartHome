package lib

// Type string, Unit string, Interval int16, Name string
type Sensor struct {
	Type           string `yaml:"Type"`
	Unit           string `yaml:"Unit"`
	Interval       int64  `yaml:"Interval"`
	Name           string `yaml:"Name"`
	Life           int    `yaml:"Life"`
	Owner          string `yaml:"Owner"`
	ID             int    `yaml:"ID"`
	CatalogAddress string `yaml:"CatalogAddress"`
}

type Service struct {
	Container_Name string
	Image          string
	Expose         []string
	Ports          []string
}

type DeviceDocker struct {
	Image          string
	Volumes        [1]string
	Container_Name string
	Environment    map[string]string
	Depends_On     []string
	Links          []string
	// Expose        string
}
type Docker struct {
	Services []DeviceDocker
}

type House struct {
	ID       int
	Name     string
	Quantity int
}

type Owner struct {
	Name   string
	ID     int
	Houses []House
}

type Configs struct {
	CatalogAddress string
	Owners         []Owner
	Services       []Service
}
