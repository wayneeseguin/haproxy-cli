package haproxy

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
)

type Haproxy struct {
	Socket string
}

type Stats struct {
	Pxname         string `json:"pxname"`
	Svname         string `json:"svname"`
	Qcur           string `json:"qcur"`
	Qmax           string `json:"qmax"`
	Scur           string `json:"scur"`
	Smax           string `json:"smax"`
	Slim           string `json:"slim"`
	Stot           string `json:"stot"`
	Bin            string `json:"bin"`
	Bout           string `json:"bout"`
	Dreq           string `json:"dreq"`
	Dresp          string `json:"dresp"`
	Ereq           string `json:"ereq"`
	Econ           string `json:"econ"`
	Eresp          string `json:"eresp"`
	Wretr          string `json:"wretr"`
	Wredis         string `json:"wredis"`
	Status         string `json:"status"`
	Weight         string `json:"weight"`
	Act            string `json:"act"`
	Bck            string `json:"bck"`
	Chkfail        string `json:"chkfail"`
	Chkdown        string `json:"chkdown"`
	Lastchg        string `json:"lastchg"`
	Downtime       string `json:"downtime"`
	Qlimit         string `json:"qlimit"`
	Pid            string `json:"pid"`
	Iid            string `json:"iid"`
	Sid            string `json:"sid"`
	Throttle       string `json:"throttle"`
	Lbtot          string `json:"lbtot"`
	Tracked        string `json:"tracked"`
	_Type          string `json:"type"`
	Rate           string `json:"rate"`
	Rate_lim       string `json:"rate_lim"`
	Rate_max       string `json:"rate_max"`
	Check_status   string `json:"check_status"`
	Check_code     string `json:"check_code"`
	Check_duration string `json:"check_duration"`
	Hrsp_1xx       string `json:"hrsp_1xx"`
	Hrsp_2xx       string `json:"hrsp_2xx"`
	Hrsp_3xx       string `json:"hrsp_3xx"`
	Hrsp_4xx       string `json:"hrsp_4xx"`
	Hrsp_5xx       string `json:"hrsp_5xx"`
	Hrsp_other     string `json:"hrsp_other"`
	Hanafail       string `json:"hanafail"`
	Req_rate       string `json:"req_rate"`
	Req_rate_max   string `json:"req_rate_max"`
	Req_tot        string `json:"req_tot"`
	Cli_abrt       string `json:"cli_abrt"`
	Srv_abrt       string `json:"srv_abrt"`
	Comp_in        string `json:"comp_in"`
	Comp_out       string `json:"comp_out"`
	Comp_byp       string `json:"comp_byp"`
	Comp_rsp       string `json:"comp_rsp"`
	Lastsess       string `json:"lastsess"`
	Last_chk       string `json:"last_chk"`
	Last_agt       string `json:"last_agt"`
	Qtime          string `json:"qtime"`
	Ctime          string `json:"ctime"`
	Rtime          string `json:"rtime"`
	Ttime          string `json:"ttime"`
}

type Info struct {
	Name                        string `json:"Name"`
	Version                     string `json:"Version"`
	Release_date                string `json:"Release_date"`
	Nbproc                      string `json:"Nbproc"`
	Process_num                 string `json:"Process_num"`
	Pid                         string `json:"Pid"`
	Uptime                      string `json:"Uptime"`
	Uptime_sec                  string `json:"Uptime_sec"`
	Memmax_MB                   string `json:"Memmax_MB"`
	Ulimitn                     string `json:"Ulimit-n"`
	Maxsock                     string `json:"Maxsock"`
	Maxconn                     string `json:"Maxconn"`
	Hard_maxconn                string `json:"Hard_maxconn"`
	CurrConns                   string `json:"CurrConns"`
	CumConns                    string `json:"CumConns"`
	CumReq                      string `json:"CumReq"`
	MaxSslConns                 string `json:"MaxSslConns"`
	CurrSslConns                string `json:"CurrSslConns"`
	CumSslConns                 string `json:"CumSslConns"`
	Maxpipes                    string `json:"Maxpipes"`
	PipesUsed                   string `json:"PipesUsed"`
	PipesFree                   string `json:"PipesFree"`
	ConnRate                    string `json:"ConnRate"`
	ConnRateLimit               string `json:"ConnRateLimit"`
	MaxConnRate                 string `json:"MaxConnRate"`
	SessRate                    string `json:"SessRate"`
	SessRateLimit               string `json:"SessRateLimit"`
	MaxSessRate                 string `json:"MaxSessRate"`
	SslRate                     string `json:"SslRate"`
	SslRateLimit                string `json:"SslRateLimit"`
	MaxSslRate                  string `json:"MaxSslRate"`
	SslFrontendKeyRate          string `json:"SslFrontendKeyRate"`
	SslFrontendMaxKeyRate       string `json:"SslFrontendMaxKeyRate"`
	SslFrontendSessionReuse_pct string `json:"SslFrontendSessionReuse_pct"`
	SslBackendKeyRate           string `json:"SslBackendKeyRate"`
	SslBackendMaxKeyRate        string `json:"SslBackendMaxKeyRate"`
	SslCacheLookups             string `json:"SslCacheLookups"`
	SslCacheMisses              string `json:"SslCacheMisses"`
	CompressBpsIn               string `json:"CompressBpsIn"`
	CompressBpsOut              string `json:"CompressBpsOut"`
	CompressBpsRateLim          string `json:"CompressBpsRateLim"`
	ZlibMemUsage                string `json:"ZlibMemUsage"`
	MaxZlibMemUsage             string `json:"MaxZlibMemUsage"`
	Tasks                       string `json:"Tasks"`
	Run_queue                   string `json:"Run_queue"`
	Idle_pct                    string `json:"Idle_pct"`
	node                        string `json:"node"`
	description                 string `json:"description"`
}

func (h *Haproxy) Cmd(cmd string) (response string, err error) {
	conn, err := net.Dial("unix", h.Socket)
	defer conn.Close()
	response = ""
	if err != nil {
		err = errors.New("Unable to connect to Haproxy socket")
		return
	}
	fmt.Fprint(conn, cmd)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		response += (scanner.Text() + "\n")
	}
	err = scanner.Err()
	return
}

func (h *Haproxy) Stats(statsType string) ([]Stats, error) {
	var stats []Stats
	var cmd string

	switch statsType {
	case "all":
		cmd = "show stat -1\n"
	case "backend":
		cmd = "show stat -1 2 -1\n"
	case "frontend":
		cmd = "show stat -1 1 -1\n"
	case "server":
		cmd = "show stat -1 4 -1\n"
	}

	response, err := h.Cmd(cmd)
	if err != nil {
		return stats, err
	} else {
		response, err := csv2json(strings.Trim(response, "# "))
		if err != nil {
			return stats, err
		} else {
			err := json.Unmarshal([]byte(response), &stats)
			if err != nil {
				return stats, err
			} else {
				return stats, nil
			}
		}
	}
}

func (h *Haproxy) Info() (Info, error) {
	var Info Info
	response, err := h.Cmd("show info \n")
	if err != nil {
		return Info, err
	} else {
		response, err := kvlines2json(response)
		if err != nil {
			return Info, err
		} else {
			err := json.Unmarshal([]byte(response), &Info)
			if err != nil {
				return Info, err
			} else {
				return Info, nil
			}
		}
	}
}

func csv2json(input string) (string, error) {
	csvReader := csv.NewReader(strings.NewReader(input))
	var headers []string
	var output bytes.Buffer
	var element bytes.Buffer
	output.WriteString("[")
	index := 0
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			output.Truncate(int(len(output.String()) - 1))
			output.WriteString("]\n")
			break
		} else if err != nil {
			fmt.Printf("Error: %s\n", err)
			return "", err
		}
		switch {
		case index == 0:
			headers = record[:]
			index += 1
		case index > 0:
			element.WriteString("{")
			for i := 0; i < len(headers); i++ {
				element.WriteString("\"" + headers[i] + "\": \"" + record[i] + "\"")
				if i == (len(headers) - 1) {
					element.WriteString("}")
				} else {
					element.WriteString(",")
				}
			}
			output.WriteString(element.String() + ",")
			element.Reset()
			index += 1
		}
	}
	return output.String(), nil
}

func kvlines2json(lines string) (string, error) {
	var response bytes.Buffer
	lines = strings.TrimSpace(lines)
	numLines := strings.Count(lines, "\n")
	reader := bufio.NewReader(strings.NewReader(lines))
	response.WriteString("{")
	for i := 0; i <= numLines; i++ {
		line, err := (reader.ReadString('\n'))
		if err != nil {
			break
		} else {
			kv := strings.Split(line, ": ")
			response.WriteString("\"" + kv[0] + "\" : \"" + strings.Trim(kv[1], "\n") + "\",")
		}
	}
	response.Truncate(int(len(response.String()) - 1)) // Ignore the final ,
	response.WriteString("}\n")
	return response.String(), nil
}
