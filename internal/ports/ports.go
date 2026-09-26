package ports

import (
	"sort"

	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
)

type PortInfo struct {
	Port    uint32
	PID     int32
	Proto   string
	State   string
	Process string
	Path    string
}

func List() ([]PortInfo, error) {
	conns, err := net.Connections("all")
	if err != nil {
		return nil, err
	}

	seen := make(map[string]bool)
	var out []PortInfo

	for _, c := range conns {
		if c.Status != "LISTEN" && c.Status != "LISTENING" {
			continue
		}

		key := c.Status + "|" + protoName(c.Type) + "|" + uitoa(c.Laddr.Port)
		if seen[key] {
			continue
		}
		seen[key] = true

		name, path := "<protected>", ""
		if c.Pid != 0 {
			if p, err := process.NewProcess(c.Pid); err == nil {
				if n, err := p.Name(); err == nil {
					name = n
				}
				if e, err := p.Exe(); err == nil {
					path = e
				}
			}
		}

		out = append(out, PortInfo{
			Port:    c.Laddr.Port,
			PID:     c.Pid,
			Proto:   protoName(c.Type),
			State:   c.Status,
			Process: name,
			Path:    path,
		})
	}

	sort.Slice(out, func(i, j int) bool { return out[i].Port < out[j].Port })
	return out, nil
}

func protoName(t uint32) string {
	switch t {
	case 1:
		return "tcp"
	case 2:
		return "udp"
	default:
		return "?"
	}
}

func uitoa(n uint32) string {
	if n == 0 {
		return "0"
	}
	var buf [10]byte
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	return string(buf[i:])
}