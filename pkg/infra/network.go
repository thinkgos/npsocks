package infra

import (
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

type Network struct {
	Name         string // net card name
	HardwareAddr net.HardwareAddr
	IPNet        *net.IPNet
}

func ActiveNetwork() (*Network, error) {
	ifcs, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, ifc := range ifcs {
		if (ifc.Flags & net.FlagUp) != 0 {
			addrs, _ := ifc.Addrs()
			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return &Network{ifc.Name, ifc.HardwareAddr, ipnet}, nil
					}
				}
			}
		}
	}
	return nil, errors.New("not any active net interfaces")
}

func GatewayIP(netCard string) (string, error) {
	if runtime.GOOS == "windows" {
		out, err := exec.Command("cmd.exe", "route print -4").Output()
		if err != nil {
			return "", err
		}
		v := regexp.MustCompile(`0\\.0\\.0\\.0 *(0|128)\\.0\\.0\\.0 *([0-9\\.]*) *([0-9\\.]*) *([0-9]*)`).Find(out)
		return string(v), nil
	}
	arg := fmt.Sprintf("route -n | grep %s | grep UG | awk '{print $2}'", netCard)
	out, err := exec.Command("/bin/sh", "-c", arg).Output()
	if err != nil {
		return "", err
	}
	return strings.Trim(string(out), "\n"), nil
}

// NetInformation 网络信息
type NetInformation struct {
	Name         string // 网卡名
	HardwareAddr net.HardwareAddr
	Mac          string
	IP           string
	Mask         string
	GatewayIP    string
}

// GetNetInformation 通过网卡获得 MAC IP IPMask GatewayIP
func GetNetInformation() (*NetInformation, error) {
	anet, err := ActiveNetwork()
	if err != nil {
		return nil, err
	}

	ip, err := GatewayIP(anet.Name)
	if err != nil {
		return nil, err
	}
	return &NetInformation{
		anet.Name,
		anet.HardwareAddr,
		hex.EncodeToString(anet.HardwareAddr),
		anet.IPNet.IP.String(),
		net.IP(anet.IPNet.Mask).String(),
		ip,
	}, nil
}

// LanIP 局域网ip地址
func LanIP() string {
	an, err := ActiveNetwork()
	if err != nil {
		return ""
	}
	return an.IPNet.IP.String()
}
