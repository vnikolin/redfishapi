package redfishapi

// DellBMCAdapter wraps a DellProvider to satisfy BMCProvider.
type DellBMCAdapter struct {
	Provider DellProvider
}

// NewDellBMC returns a BMCProvider backed by a DellProvider.
func NewDellBMC(p DellProvider) BMCProvider {
	return &DellBMCAdapter{Provider: p}
}

func (a *DellBMCAdapter) GetSystemInfo() (SystemData, error) {
	return a.Provider.GetSystemInfoDell()
}

func (a *DellBMCAdapter) GetFirmware() ([]FirmwareData, error) {
	return a.Provider.GetFirmwareDell()
}

func (a *DellBMCAdapter) StartServer() (string, error) {
	return a.Provider.StartServerDell()
}

func (a *DellBMCAdapter) StopServer() (string, error) {
	return a.Provider.StopServerDell()
}

func (a *DellBMCAdapter) GetServerPowerState() (string, error) {
	return a.Provider.GetServerPowerStateDell()
}

func (a *DellBMCAdapter) GetProcessorHealth() ([]HealthList, error) {
	return a.Provider.GetProcessorHealthDell()
}

func (a *DellBMCAdapter) GetPowerHealth() ([]HealthList, error) {
	return a.Provider.GetPowerHealthDell()
}

func (a *DellBMCAdapter) GetThermalHealth() ([]HealthList, error) {
	return a.Provider.GetSensorsHealthDell()
}

func (a *DellBMCAdapter) GetUserAccounts() ([]Accounts, error) {
	return a.Provider.GetUserAccountsDell()
}

// HPBMCAdapter wraps an HPProvider to satisfy BMCProvider.
type HPBMCAdapter struct {
	Provider HPProvider
}

// NewHPBMC returns a BMCProvider backed by an HPProvider.
func NewHPBMC(p HPProvider) BMCProvider {
	return &HPBMCAdapter{Provider: p}
}

func (a *HPBMCAdapter) GetSystemInfo() (SystemData, error) {
	return a.Provider.GetSystemInfoHP()
}

func (a *HPBMCAdapter) GetFirmware() ([]FirmwareData, error) {
	return a.Provider.GetFirmwareHP()
}

func (a *HPBMCAdapter) StartServer() (string, error) {
	return a.Provider.StartServerHP()
}

func (a *HPBMCAdapter) StopServer() (string, error) {
	return a.Provider.StopServerHP()
}

func (a *HPBMCAdapter) GetServerPowerState() (string, error) {
	return a.Provider.GetServerPowerStateHP()
}

func (a *HPBMCAdapter) GetProcessorHealth() ([]HealthList, error) {
	return a.Provider.GetProcessorHealthHP()
}

func (a *HPBMCAdapter) GetPowerHealth() ([]HealthList, error) {
	return a.Provider.GetPowerHealthHP()
}

func (a *HPBMCAdapter) GetThermalHealth() ([]HealthList, error) {
	return a.Provider.GetThermalHealthHP()
}

func (a *HPBMCAdapter) GetUserAccounts() ([]Accounts, error) {
	return a.Provider.GetUserAccountsHP()
}
