package internal

type Target struct {
	Path    string `json:"path"`
	Exclude string `json:"exclude,omitempty"`
}

type PacketFile struct {
	Name    string        `json:"name"`
	Version string        `json:"ver"`
	Targets []interface{} `json:"targets"`
	Packets []Dependency  `json:"packets"`
}

type Dependency struct {
	Name    string `json:"name"`
	Version string `json:"ver"`
}

type PackagesFile struct {
	Packages []Dependency `json:"packages"`
}
