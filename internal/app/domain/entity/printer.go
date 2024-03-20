package entity

type Printer struct {
	Name      string            `json:"name"`
	IsDefault bool              `json:"isDefault"`
	Selected  bool              `json:"selected"`
	Options   map[string]string `json:"options"`
}
