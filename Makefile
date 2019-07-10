
build:
	./gen.sh

	# go
	cd go && go build && cd ../

	# php
	cd php && composer install && cd ../

	# java
	cd java
	mvn protobuf:compile
	mvn protobuf:compile-custom
	mvn clean package -Dmaven.test.skip=true
	cd ../

.PHONY: clean

clean:
	rm -rf go/pbf/* php/pbf/*
