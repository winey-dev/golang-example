
# VIPER SAMPLE
 viper로 configuration 정보를 load 하는 방법에 대한 Sample Code 입니다.
 viper로 configuration 변경 정보를 Notify 받고 후행 적으로 처리하는 방법에 대한 Sample Code 입니다.

## Download
 download 방법은 아래와 같습니다.
 LOCAL GOPATH 의 src 폴더 아래에 해당 파일을 저장 합니다.
 ```
 $ git clone ${url}/viper $GOPATH/src/viper
 ```

## BUILD
  다운로드한 viper의 root 폴더로 이동하여 `go build` 명령어를 수행 합니다.
 ```
  $ cd $GOPAHT/src/viper
  $ go build

 ```

## RUN
 `viper/config/conf.yaml` 에 설정파일 있습니다. 
 ```
 $ cat {SOME DIR}/viper/config/conf.yaml
 server:
     logLevel: info
     address: localhost
     port: 8080
 ```


 ```
 $ ./viper
 config load succ..
 LOG LEVEL      : info
 SERVER ADDRESS : localhost
 SERVER PORT    : 8080
 ```

## TEST
conf.yaml을 수정해보세요
```
config load succ..
LOG LEVEL      : info
SERVER ADDRESS : localhost
SERVER PORT    : 8081
change server.logLevel info->debug
```

## Image Build
`build/docker` 폴더로 이동 합니다.
```
$ cd build/docker
$ ls 
Dockerfile Makefile
```

`Dockerfile`의 `FROM`를 상황에 맞게 수정 합니다.
```
FROM ${fixed}

WORKDIR /

COPY config /config
COPY viper .

CMD ["./viper"]
```
`Makefile` 의 `TARGET`를 수정합니다.
```
REPO_HOST=${your repository host}
TARGET=$(REPO_HOST)/${fixed repository}

all:
    cd ../../; go build
    mv ../../viper ./
    mkdir config; cp ../../config/conf.yaml ./config
    @podman build --network=host -t $(TARGET) .
    @podman login $(REPO_HOST)
    @podman push $(TARGET)
```

### Makefile 수정 시 
`REPO_HOST`와 `REPOSITORY` 를 헷갈리지 마세요.
`REPO_HOST`는 단순히 당신의 Repository 주소 정보가 됩니다. 
`${fixed repository}` 에는 Repository 명과 Image, ImageTag가 들어갑니다.
 
 ex)
 ```
REPO_HOST=localhost:9000
TARGET=$(REPO_HOST)/smlee/viper:0.0.1

all:
    cd ../../; go build
    mv ../../viper ./
    mkdir config; cp ../../config/conf.yaml ./config
    @podman build --network=host -t $(TARGET) .
    @podman login $(REPO_HOST)
    @podman push $(TARGET)
 ```


## Kubernetes Deploy
kubernetes에 Deploy 하기전 반드시 [Image Build](#image-build)를 수행 해주세요. <br>
`build/k8s` 폴더로 이동합니다.

```
$ cd build/k8s
$ ls 
deploy.yaml configmap.yaml
```

deploy.yaml에서 [Image build](#image-build)에서 만든 image source로 수정하세요.
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: viper-sample
  labels:
    app: viper-sample
spec:
  replicas: 1
  selector:
    matchLabels:
      app: viper-sample
  template:
    metadata:
      labels:
        app: viper-sample
    spec:
      containers:
      - name: viper-sample
        image:  ${fixed image source}
        imagePullPolicy: Always
        volumeMounts:
        - name: config-volume
          mountPath: /config
      volumes:
      - name: config-volume
        configMap:
          name: server-config
```

```
$ cd viper/build
$ kubectl apply -f k8s/
configmap/server-config created
deployment.apps/viper-sample created
```

*  `kubectl edit` 를 이용해서 configmap을 수정
* 수정된 configmap 파일을 `kubectl apply -f configmap.yaml`로 적용 

쿠버네티스 [공식 문서](https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/#mounted-configmaps-are-updated-automatically)에 따르면 configmap이 실제 pod에 적용되는데는 꽤 시간이 필요하다고 합니다.
테스트에 참고해주세요.

>When a mounted ConfigMap is updated, the projected content is eventually updated too. This applies in the case where an optionally referenced ConfigMap comes into existence after a pod has started.
>
>The kubelet checks whether the mounted ConfigMap is fresh on every periodic sync. However, it uses its local TTL-based cache for getting the current value of the ConfigMap. As a result, the total delay from the moment when the ConfigMap is updated to the moment when new keys are projected to the pod can be as long as kubelet sync period (1 minute by default) + TTL of ConfigMaps cache (1 minute by default) in kubelet.
