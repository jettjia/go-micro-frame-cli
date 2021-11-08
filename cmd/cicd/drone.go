package cicd

import (
	"github.com/jettjia/go-micro-frame-cli/util"
	"log"
)

var (
	oldStr = `kind: pipeline
type: docker
name: default

clone:
  disable: true

steps:
  - name: clone
    pull: if-not-exists
    image: appleboy/drone-ssh
    settings:
      host: 10.4.7.100
      username: root
      password: 123456
      port: 22
      script:
        - cd /drone/src
        - git clone ssh://git@10.4.7.100:10022/jettjia/srv-example-grpc-test.git
        - cd srv-example-grpc-test
        - git checkout master
        - chmod -R 777 ./*
  
  - name: build-image
    pull: if-not-exists
    image: appleboy/drone-ssh
    settings:
      host: 10.4.7.100
      username: root
      password: 123456
      port: 22
      script:
        - cd /drone/src/srv-example-grpc-test
        - docker build -t my-go-micro .
        - docker tag my-go-micro 10.4.7.100:85/test/my-go-micro:latest
        - docker login -u admin -p Harbor12345 10.4.7.100:85
        - docker push 10.4.7.100:85/test/my-go-micro:latest

        - docker rmi my-go-micro
        - docker rmi 10.4.7.100:85/test/my-go-micro:latest

  # 发布 dev环境
  - name: deploy-dev
    pull: if-not-exists
    image: appleboy/drone-ssh
    settings:
      host: 10.4.7.71
      username: root
      password: 123456
      port: 22
      script:
        - docker login -u admin -p Harbor12345 10.4.7.100:85
        - docker pull 10.4.7.100:85/test/my-go-micro:latest
        - list=$(docker ps -a | grep srv-example | awk '{print $1}')
        - test "$list" = "" && echo "none my-go-micro containers running" || docker stop $list
        - docker rm $list
        - docker run -p 50051:50051 -d -v /data/my-go-micro:/apps/tmp 10.4.7.100:85/test/my-go-micro:latest
    when:
      event: [ push, pull_request ]
      branch: [ dev ]

  # 发布 master环境
  - name: deploy-prod
    pull: if-not-exists
    image: appleboy/drone-ssh
    settings:
      host: 10.4.7.102
      username: root
      password: 123456
      port: 22
      script:
        - docker login -u admin -p Harbor12345 10.4.7.100:85
        - docker pull 10.4.7.100:85/test/my-go-micro:latest
        - list=$(docker ps -a | grep srv-example | awk '{print $1}')
        - test "$list" = "" && echo "none my-go-micro containers running" || docker stop $list
        - docker rm $list
        - docker run -p 50051:50051 -d -v /data/my-go-micro:/apps/tmp 10.4.7.100:85/test/my-go-micro:latest
    when:
      event: [ push, pull_request ]
      branch: [ master ]


volumes:
  - name: sshkeys
    host:
      path: /root/.ssh

  - name: dockerdaemon
    host:
      path: /etc/docker/daemon.json
`
)

// Run 当前项目生成 dockerfile 文件
func Run(args []string) {
	//imgName := args[0]

	log.Println("create .drone start ...")

	util.WriteStringToFileMethod(".drone.yml", oldStr)
	log.Println("Done")
}
