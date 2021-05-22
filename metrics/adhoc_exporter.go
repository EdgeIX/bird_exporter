package metrics

import (
	"fmt"
	"os"
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
	labels := []string{"table", "use", "community", "host"}
	//community := fmt.Sprintf("%v_%v_%v", com.ASN, com.First, com.Last)
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	filteredDesc := prometheus.NewDesc(m.prefix+"stats", "Adhoc Exported Data", labels, nil)
	//filteredDesc := prometheus.NewDesc(m.prefix+fmt.Sprintf("%s", community), fmt.Sprintf("%s using community: %s", com.Name, community), labels, nil)

	l := []string{data.Name, com.Name, fmt.Sprintf("%v:%v:%v", com.ASN, com.First, com.Last), hostname}
	ch <- prometheus.MustNewConstMetric(filteredDesc, prometheus.GaugeValue, float64(data.Matched), l...)
}
