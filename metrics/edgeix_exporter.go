package metrics

import (
	"github.com/czerwonk/bird_exporter/client"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"github.com/czerwonk/bird_exporter/protocol"
	//"strconv"
)

type edgeIXDesc struct {
	rpkiInvalidDesc		*prometheus.Desc
}

type edgeIXMetricExporter struct {
	descriptions *edgeIXDesc
	client	client.Client
}

func EdgeIXExporter(prefix string, client client.Client) MetricExporter {
	//d := make(map[string]*edgeIXDesc)
	d := getDescE(prefix + "edgeix")

	return &edgeIXMetricExporter{descriptions: d, client: client}
}

func getDescE(prefix string) *edgeIXDesc {
	labels := []string{"name"}
	d := &edgeIXDesc{}

	d.rpkiInvalidDesc = prometheus.NewDesc(prefix+"_rpki_invalids", "RPKI Invalid Per Table", labels, nil)
	return d
}

func (m *edgeIXMetricExporter) Describe(ch chan<- *prometheus.Desc) {
	d := m.descriptions
	ch <- d.rpkiInvalidDesc
}

func (m *edgeIXMetricExporter) Export(p *protocol.Protocol, ch chan<- prometheus.Metric, newFormat bool) {
	if p == nil {
		return
	}
	d := m.descriptions
	data, err := m.client.GetEdgeIX()
	if err != nil {
		log.Errorln(err)
		return
	}

	for _, table := range data {
		l := []string{table.TableName}
		//l.name = table.TableName
		ch <- prometheus.MustNewConstMetric(d.rpkiInvalidDesc, prometheus.GaugeValue, float64(table.MatchingRoutes), l ...)
		//l := []string{strconv.Itoa(table.TotalRoutes), table.TableName}
		//ch <- prometheus.MustNewConstMetric(d.rpkiInvalidDesc, prometheus.GaugeValue, float64(table.MatchingRoutes), l ...)
	}

}
