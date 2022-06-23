package inverter

type Appliance struct {
	Name string `json:"name"`
	Watt int    `json:"watt"`
}
type Appliances struct {
	AllAppliance []Appliance `json:"appliances"`
}
type UserNeed struct {
	ApplianceName  string `json:"name"`
	ApplianceCount int    `json:"count"`
}
type UserInput struct {
	ApplianceDetails []UserNeed `json:"applianceDetails"`
	BackupTime       int        `json:"backupTime"`
}
type Results struct {
	TotalLoad          int     `json:"totalLoad"`
	PowerFactor        float32 `json:"powerFactor"`
	VARatingOfInverter float32 `json:"VARatingOfInverter"`
	BackupTime         int     `json:"backupTime"`
	BatterySize        float32 `json:"batterySize"`
	InverterPower      float32 `json:"inverterPower"`
	TotalVAh           float32 `json:"totalVAh"`
}
