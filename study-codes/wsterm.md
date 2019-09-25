- main.go

```go
package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"strings"
)

var (
	Downlinks     = make(chan Msg)
	LogChan       = make(chan string)
	jwt           = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjdhMTlkOGNlLWU0OTItNDczMy1iNDI4LTE1MTYwODM3ZTQ1MSIsImFwcElkIjoiMGNhMWVmMDEtZmE3MS00NDc1LThiZTUtNzBhZjc1ODY1NTg5IiwiaWF0IjoxNTYwNTMyODI2fQ.69OlvCsKeeCQbn_9QZSzCNzBDpDD3W0ofx6Ra6sx_gw"
	Data          = NewDevicePropData()
	Logger        = NewLogger(10000, 67+10)
	devices, hide string
	isStop        bool
)

func main() {
	// handle the cli arguments
	var address string
	flag.StringVar(&address, "vm", "", "type the address such as https://10.5.2.200")
	flag.StringVar(&jwt, "jwt", jwt, "type the JWT")
	flag.StringVar(&devices, "f", "", "only follow the devices separated by comma")
	flag.StringVar(&hide, "hide", "", "hide the subscribed channels separated by comma")

	flag.Parse()

	log.SetFlags(0)

	EnvJwt := strings.TrimSpace(os.Getenv("JWT"))
	if len(EnvJwt) > 0 {
		jwt = EnvJwt
	}
	EnvVM := strings.TrimSpace(os.Getenv("VM"))
	if len(EnvVM) > 0 {
		address = EnvVM
	}

	addr, err := url.Parse(address)
	if err != nil {
		log.Fatalf("You have a wrong url addres: %s", err)
	}
	schema := "ws"
	if addr.Scheme == "https" {
		schema = "wss"
	}

	u := url.URL{Scheme: schema, Host: addr.Host, Path: "/ws/", RawQuery: "transport=websocket&token=" + jwt}
	c := Connecting(u.String())

	go Listen(c)
	go Ping(c)

	// https://bitbucket.org/greensmithenergy/gems-hmi-server/src/1b177d4aeceb1fec23daf929f3245c3b85dad610/src/devices/websocket-address.js?at=master
	Subscribe(c, NewDeviceSubcribleList(devices))
	Subscribe(c, `42["v1.operations.subscribe"]`)
	Subscribe(c, `42["v1.alarms.subscribe"]`)
	Subscribe(c, `42["v1.events.subscribe"]`)
	Subscribe(c, `42["v1.reports.subscribe"]`)
	Subscribe(c, `42["v1.ui.subscribe"]`)
	Subscribe(c, `42["v1.dcm.subscribe"]`)

	DrawUI()
}

```

- parser.go
```go
package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

type DevicesDataObject []struct {
	SiteID    string `json:"siteId"`
	TimeStamp int    `json:"timestamp"`
	Delta     []struct {
		DeviceID string          `json:"deviceId"`
		Status   [][]interface{} `json:"status"`
	} `json:"delta"`
}

type Device struct {
	Status map[string]string
	ID     string
}

type DevicePropData struct {
	DeviceSlice         []*Device
	DevicePropMap       map[string][]string // directly for the second panel
	CurrentDeviceRowNum int
}

func NewDevicePropData() *DevicePropData {
	res := DevicePropData{DevicePropMap: make(map[string][]string), DeviceSlice: []*Device{}}
	return &res
}

func (p *DevicePropData) GetDeviceIdxFromDeviceSlice(id string) int {
	for i, device := range p.DeviceSlice {
		if device.ID == id {
			return i
		}
	}
	return -1
}

func (p *DevicePropData) CreateDeviceLists() {
	for _, device := range p.DeviceSlice {
		var propList []string
		for k, v := range device.Status {
			current := fmt.Sprintf(`%s: %s`, k, RemoveSuffixZero(v))
			propList = append(propList, current)
		}
		sort.Strings(propList)
		for i, s := range propList {
			propList[i] = fmt.Sprintf("[%d] %s", i+1, s)
		}
		p.DevicePropMap[device.ID] = propList
	}
}

func (p *DevicePropData) GetDeviceList() []string {
	var res []string
	for k := range p.DevicePropMap {
		res = append(res, k)
	}
	sort.Strings(res)
	return res
}

func (p *DevicePropData) GetDevicePropList() []string {
	current := p.GetCurrentDevice()
	if current == "" {
		return []string{}
	}
	return p.DevicePropMap[current]
}

func (p *DevicePropData) GetCurrentDevice() string {
	current := p.GetDeviceList()
	if len(current) == 0 {
		return ""
	}
	return current[p.CurrentDeviceRowNum]
}

func (p *DevicePropData) CreateOrUpdateDeviceSliceFromMessage(s string) {
	var dto DevicesDataObject
	err := json.Unmarshal([]byte(s), &dto)
	if err != nil {
		return
	}

	for _, outer := range dto {
		for _, inner := range outer.Delta {
			current := &Device{Status: make(map[string]string), ID: inner.DeviceID}
			for _, status := range inner.Status {
				prop, err1 := ToStr(status[0])
				value, err2 := ToStr(status[1])
				if err1 == nil && err2 == nil {
					current.Status[prop] = value
				}
				for k, v := range current.Status {
					current.Status[k] = v
				}
			}
			idx := p.GetDeviceIdxFromDeviceSlice(inner.DeviceID)
			if idx == -1 { // create device
				p.DeviceSlice = append(p.DeviceSlice, current)
			} else { // update device
				old := p.DeviceSlice[idx]
				for k, v := range current.Status {
					old.Status[k] = v
				}
			}
		}
	}
}

```
- network.go
```go
package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"time"
)

type Msg struct {
	Label   string
	Message string
}

type WsLogger struct {
	dataList  []string
	maxLogNum int
	colSize   int
}

func NewLogger(maxLogNum int, colSize int) *WsLogger {
	return &WsLogger{[]string{}, maxLogNum, colSize}
}

func (l *WsLogger) Add(text string) {
	current := WrapText(text, l.colSize)
	l.dataList = append(current, l.dataList...)
	if len(l.dataList) > l.maxLogNum {
		l.dataList = l.dataList[:l.maxLogNum]
	}
}

func (l *WsLogger) WriteToTxt() string {
	fileName := GetTimeStamp() + ".txt"
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("cannot create txt: %s", err)
	}
	w := bufio.NewWriter(f)
	for _, line := range l.dataList {
		_, err = w.WriteString(line + "\n")
		if err != nil {
			log.Fatalf("cannot write to txt: %s", err)
		}
	}
	err = w.Flush()
	if err != nil {
		log.Fatalf("cannot close file: %s", err)
	}
	return fileName
}

func NewMsg(bs []byte) Msg {
	bs = bytes.TrimSpace(bs)
	commaIdx := bytes.IndexByte(bs, ',')
	bracketIdx := bytes.IndexByte(bs, '[')
	label := bs[bracketIdx+2 : commaIdx-1]
	message := bs[commaIdx+1 : len(bs)-1]
	return Msg{string(label), string(message)}
}

func Subscribe(conn *websocket.Conn, text string) {
	err := conn.WriteMessage(websocket.TextMessage, []byte(text))
	if err != nil {
		log.Fatalf("cannot subscribe due to %s", err)
	}
}

func Listen(c *websocket.Conn) {
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			LogChan <- fmt.Sprintf("read %s", err)
		}
		if bytes.Contains(message, []byte(`v1.`)) {
			Downlinks <- NewMsg(message)
		}
		LogChan <- string(message)
	}
}

func Ping(c *websocket.Conn) {
	ticker := time.NewTicker(time.Second * 15)
	defer ticker.Stop()
	for {
		<-ticker.C
		err := c.WriteMessage(websocket.TextMessage, []byte{'2'})
		LogChan <- fmt.Sprint("send: 2")
		if err != nil {
			LogChan <- fmt.Sprintf("Error: %s", err)
		}
	}
}

func Connecting(url string) *websocket.Conn {
	dialer := websocket.DefaultDialer
	dialer.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	c, resp, err := dialer.Dial(url, nil)
	if err != nil {
		log.Fatalf("connection falied: %s", err)
	}
	if status := resp.StatusCode; status != 101 {
		log.Fatalf("connection falied: status code %d", resp.StatusCode)
	}
	return c
}
```

- ui.go
```go
package main

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"log"
	"strings"
)

var panelNum = 0

func DrawUI() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	// first list
	l1 := widgets.NewList()
	l1.Title = "Devices"
	l1.Rows = Data.GetDeviceList()
	l1.WrapText = true
	l1.SetRect(0, 0, 15, 45)
	l1.BorderStyle.Fg = ui.ColorCyan
	l1.TitleStyle.Fg = ui.ColorCyan
	l1.TextStyle.Fg = ui.ColorYellow

	// second list
	l2 := widgets.NewList()
	l2.Title = "Properties: " + Data.GetCurrentDevice()
	l2.WrapText = true
	l2.Rows = Data.GetDevicePropList()
	l2.SetRect(15, 0, 55, 45)
	l2.TextStyle.Fg = ui.ColorYellow

	// third list
	l3 := widgets.NewList()
	l3.SetRect(55, 0, 135, 45)
	l3.Title = "Logging"
	l3.Rows = Logger.dataList
	l3.TextStyle.Fg = ui.ColorYellow

	p := widgets.NewParagraph()
	p.Text = fmt.Sprintf(`[WSTERM %s] Use arrows to control panels; use q to quit; use h to scroll top and e to bottom, use p to print, use s to switch stop`, _VERSION)
	p.SetRect(0, 45, 135, 48)
	p.TextStyle.Fg = ui.ColorYellow
	p.BorderStyle.Fg = ui.ColorYellow

	uiEvents := ui.PollEvents()
	ui.Render(l1, l2, l3, p)

	var panels = []*ui.Block{&l1.Block, &l2.Block, &l3.Block}
	switchPanel := func(idx int) {
		for i, e := range panels {
			if i == idx {
				e.BorderStyle.Fg = ui.ColorCyan
				e.TitleStyle.Fg = ui.ColorCyan
			} else {
				e.BorderStyle.Fg = ui.ColorYellow
				e.TitleStyle.Fg = ui.ColorYellow
			}
		}
	}

	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "s":
				isStop = !isStop
			case "<Down>":
				switch panelNum {
				case 0:
					l1.ScrollDown()
					Data.CurrentDeviceRowNum = l1.SelectedRow
					l2.Rows = Data.GetDevicePropList()
					l2.Title = "Properties: " + Data.GetCurrentDevice()
				case 1:
					l2.ScrollDown()
				case 2:
					l3.ScrollDown()
				}
			case "p":
				f := Logger.WriteToTxt()
				p.Text = fmt.Sprintf("  A text file %s has been written", f)
			case "<Up>":
				switch panelNum {
				case 0:
					l1.ScrollUp()
					Data.CurrentDeviceRowNum = l1.SelectedRow
					l2.Rows = Data.GetDevicePropList()
					l2.Title = "Properties: " + Data.GetCurrentDevice()
				case 1:
					l2.ScrollUp()
				case 2:
					l3.ScrollUp()
				}
			case "h", "<Home>":
				switch panelNum {
				case 0:
					l1.ScrollTop()
				case 1:
					l2.ScrollTop()
				case 2:
					l3.ScrollTop()
				}
			case "e", "<End>":
				switch panelNum {
				case 0:
					l1.ScrollBottom()
				case 1:
					l2.ScrollBottom()
				case 2:
					l3.ScrollBottom()
				}
			case "<Left>":
				panelNum--
				if panelNum == -1 {
					panelNum = 2
				}
			case "<Right>":
				panelNum++
				if panelNum == 3 {
					panelNum = 0
				}

			case "q", "<C-c>":
				return
			}
			switchPanel(panelNum)
			ui.Render(l1, l2, l3, p)

		case m := <-Downlinks:
			if m.Label == "v1.devices.update" {
				Data.CreateOrUpdateDeviceSliceFromMessage(m.Message)
				Data.CreateDeviceLists()
				l1.Rows = Data.GetDeviceList()
				l2.Rows = Data.GetDevicePropList()
			}
			ui.Render(l1, l2, l3, p)
		case lg := <-LogChan:
			if isStop {
				goto exit
			}
			if len(hide) != 0 {
				for _, h := range strings.Split(hide, ",") {
					if strings.Contains(lg, "v1."+h) {
						goto exit
					}
				}
			}
			Logger.Add(lg)
			l3.Rows = Logger.dataList
			ui.Render(l1, l2, l3, p)
		exit:
		}
	}

}
```

- utils.go
```go
package main

import (
	"fmt"
	"strings"
	"time"
)

func GetTimeStamp() string {
	t := time.Now()
	return t.Format("2006-01-02-15-04-05")
}

func WrapText(text string, colBreak int) []string {
	if colBreak < 1 {
		return []string{text}
	}
	text = strings.TrimSpace(text)
	var wrapped []string
	var i int
	for i = 0; len(text[i:]) > colBreak; i += colBreak {
		wrapped = append(wrapped, " " + text[i:i+colBreak])

	}
	wrapped = append([]string{"------------------" + GetTimeStamp() + "------------------"}, wrapped...)
	wrapped = append(wrapped, " "+text[i:])
	return wrapped
}

func RemoveSuffixZero(text string) string {
	n, j := len(text), -1
	for i := n - 1; i >= 0; i-- {
		if text[i] == '0' {
			j = i
		} else {
			break
		}
	}
	if j == -1 {
		return text
	} else if j >= 1 && text[j-1] == '.' {
		j--
	}
	return text[:j]
}

func ToStr(value interface{}) (string, error) {
	switch value := value.(type) {
	case string:
		return value, nil
	case int, int32, int64, bool:
		return fmt.Sprintf("%v", value), nil
	case float32, float64:
		return fmt.Sprintf("%f", value), nil
	default:
		return "", fmt.Errorf("cannot transform as a type of %T", value)
	}
}

// translate the cli arguments to websocket command
func NewDeviceSubcribleList(text string) string {
	if len(text) == 0 {
		return `42["v1.devices.subscribe"]`
	}
	var res []string
	for _, d := range strings.Split(text, ",") {
		res = append(res, fmt.Sprintf(`{"id": "%s"}`, d))
	}
	return fmt.Sprintf(`42["v1.devices.subscribe", [%s]]`, strings.Join(res, ", "))
}
```




