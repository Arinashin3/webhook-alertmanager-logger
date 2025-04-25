package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type ReceivedData struct {
	Alerts []struct {
		Status string `json:"status"`
		Labels struct {
			Alertname string `json:"alertname"`
			Component string `json:"component"`
			HostGroup string `json:"host_group"`
			Hostname  string `json:"hostname"`
			Ipaddr    string `json:"ipaddr"`
			OsType    string `json:"os_type"`
			Severity  string `json:"severity"`
		} `json:"labels"`
		Annotations struct {
			Summary string `json:"summary"`
		} `json:"annotations"`
		StartsAt string `json:"startsAt"`
		EndsAt   string `json:"endsAt"`
	} `json:"alerts"`
}

func Logger(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	fileName := "received.log"
	var t ReceivedData
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	if json.Unmarshal(body, &t) != nil {
		log.Println("Unmarshal")
	}
	r.Body.Close()

	fi, err := os.Stat(fileName)
	if fi == nil {
		f, err := os.Create(fileName)
		log.Println(fi, err)
		f.Close()
	}
	if err != nil {
		log.Println(err)
	}
	for _, alert := range t.Alerts {
		eventAt, err := time.Parse(time.RFC3339, alert.StartsAt)
		if err != nil {
			log.Println(err)

		}

		if eventAt.Unix() > now.AddDate(0, 0, -1).Unix() {
			continue
		}
		str := " alertname=" + alert.Labels.Alertname
		str += " component=" + alert.Labels.Component
		str += " host_group=" + alert.Labels.HostGroup
		str += " hostname=" + alert.Labels.Hostname
		str += " ipaddr=" + alert.Labels.Ipaddr
		str += " os_type=" + alert.Labels.OsType
		str += " severity=" + alert.Labels.Severity
		msg := eventAt.String() + str
		err = os.WriteFile(fileName, []byte(msg), 644)
		io.
			log.Println(err)
	}
	w.Write([]byte("Success!!"))

}
