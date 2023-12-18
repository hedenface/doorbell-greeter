.DEFAULT_GOAL := doorbell
FORCE: ;

doorbell: FORCE
	sudo cp passwd /etc/doorbell.passwd
	_html-templates/convert-to-functions.sh
	go build -o doorbell .
