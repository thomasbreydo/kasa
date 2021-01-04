package device

import (
	"encoding/json"

	"github.com/cgreenhalgh/hs1xxplug"
)

type SysInfoJSON = map[string]interface{}

func GetSysInfo(plug *hs1xxplug.Hs1xxPlug) (SysInfoJSON, error) {
	sysInfoJson, err := (*plug).SystemInfo() // format json
	if err != nil {
		return nil, err
	}
	var sysInfo SysInfoJSON
	jsonErr := json.Unmarshal([]byte(sysInfoJson), &sysInfo)
	return sysInfo, jsonErr
}

func TurnOff(plug *hs1xxplug.Hs1xxPlug) error {
	return (*plug).TurnOff()
}

func TurnOn(plug *hs1xxplug.Hs1xxPlug) error {
	return (*plug).TurnOn()
}

func IsOff(plug *hs1xxplug.Hs1xxPlug) (bool, error) {
	sysInfo, err := GetSysInfo(plug)
	isOff := sysInfo["system"].(SysInfoJSON)["get_sysinfo"].(SysInfoJSON)["relay_state"] == 0.0
	return isOff, err
}

func IsOn(plug *hs1xxplug.Hs1xxPlug) (bool, error) {
	isOff, err := IsOff(plug)
	return !isOff, err
}
