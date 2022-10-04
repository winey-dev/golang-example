#Include MAKE CONFIG

REPO=$(REPO)
TARGET=$(REPO)/smlee/viper:0.0.1

all:
	cd ../../; go build
	mv ../../viper ./
	mkdir config; cp ../../config/conf.yaml ./config
	@podman build --network=host -t $(TARGET) .
	@podman login $(REPO) 
	@podman push $(TARGET)
