package metrics

import (
	"fmt"
	"os"
	"regexp"
	"github.com/czerwonk/bird_exporter/protocol"
	"github.com/prometheus/client_golang/prometheus"
)

type AdhocExporter struct {
	prefix		string
}

func NewAdhocExporter(prefix string) *AdhocExporter {
	m := &AdhocExporter{prefix: prefix}
	return m
}

func (m *AdhocExporter) Describe(ch chan<- *prometheus.Desc) {
}

func (m *AdhocExporter) Export(p *protocol.Protocol, ch chan<- prometheus.Metric, data *protocol.Adhoc, com protocol.LargeCommunity) {
	labels := []string{"table", "use", "community", "host", "ip_version", "asn"}

	table := data.Name
	asn := "24224"
	switch data.Name {
		case "master4", "master6":
			table = data.Name
		default:
			if com.MasterOnly {
				return
			} else {
				asn_table := regexp.MustCompile("^t_\\d+_as(\\d+)$").FindStringSubmatch(data.Name)
				if len(asn_table) > 0 {
					asn = asn_table[1]
				}
			}
	}

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	filteredDesc := prometheus.NewDesc(m.prefix+"stats", "Adhoc Exported Data", labels, nil)
	l := []string{table, com.Name, fmt.Sprintf("%v:%v:%v", com.ASN, com.First, com.Last), hostname, p.IPVersion, asn}
	ch <- prometheus.MustNewConstMetric(filteredDesc, prometheus.GaugeValue, float64(data.Matched), l...)
}
