.DEFAULT_GOAL := doorbell
FORCE: ;

doorbell: FORCE
	cd src/ && make doorbell
