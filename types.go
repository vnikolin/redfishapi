package redfishapi

type SystemDataHp struct {
	PowerState      string `json:"power_state"`
	SerialNumber    string `json:"serial_number"`
	Health          string `json:"health"`
	SystemType      string `json:"system_type"`
	Model           string `json:"model"`
	Memory          int    `json:"memory"`
	Processors      int    `json:"processors"`
	ProcessorFamily string `json:"processor_family"`
}

type FirmwareData struct {
	Name       string `json:"name"`
	Id         string `json:"id"`
	Version    string `json:"version"`
	Updateable bool   `json:"updateable"`
}
