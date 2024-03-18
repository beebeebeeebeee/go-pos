package entity

type Printer struct {
	Name      string
	IsDefault bool
	Selected  bool
	Options   map[string]string
}
