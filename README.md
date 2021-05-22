# EdgeIX Fork of Bird Exporter
![EdgeIX](https://www.edgeix.net/img/logo.png)

Adding the ability to handle Adhoc metrics based on Large BGP Communities for the EdgeIX Metric Platform. All credit remains with the original creator.

## bird_exporter 
[![Go Report Card](https://goreportcard.com/badge/github.com/czerwonk/bird_exporter)](https://goreportcard.com/report/github.com/czerwonk/bird_exporter)

Metric exporter for bird routing daemon to use with Prometheus

## Remarks
Since bird_exporter uses the bird unix sockets, bird has to be installed on the same maschine as bird_exporter. Also the user executing bird_exporter must have permission to access the bird socket files. 

## Third Party Components
This software uses components of the following projects
* Prometheus Go client library (https://github.com/prometheus/client_golang)

## License
(c) Daniel Czerwonk, 2016. Licensed under [MIT](LICENSE) license.

## Prometheus
see https://prometheus.io/

## Bird routing daemon
see http://bird.network.cz/
