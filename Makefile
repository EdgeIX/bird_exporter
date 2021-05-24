deploy:
	go build .

	mkdir -p /opt/bird_exporter/

	mv bird_exporter /opt/bird_exporter/

	cp exporter.service /etc/systemd/exporter.service

	systemctl enable exporter

	systemctl start exporter
