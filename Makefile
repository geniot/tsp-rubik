PROJECT_NAME := tsp-rubik
PROGRAM_NAME := rubik
DEPLOY_PATH := /mnt/SDCARD/Apps/Rubik

IP := 192.168.0.101
USN := root
PWD := tina

all: clean docker deploy

dist: clean docker zip

clean:
	rm dist -rf

docker:
	#docker run -d --name trimui-sdk -c 1024 -it --volume=/opt/TrimuiProjects/:/work/ --workdir=/work/ trimui-sdk
	docker exec trimui-sdk /bin/bash -c 'cd ${PROJECT_NAME} && make build'

build:
	go build -tags="sdl es2" -o dist/${PROGRAM_NAME} ${PROJECT_NAME}/src/

deploy:
	sshpass -p ${PWD} ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null ${USN}@${IP} "rm ${DEPLOY_PATH}/${PROGRAM_NAME} -f"
	sshpass -p ${PWD} scp -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null dist/${PROGRAM_NAME} ${USN}@${IP}:${DEPLOY_PATH}
	sshpass -p ${PWD} ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null ${USN}@${IP} "chmod 777 ${DEPLOY_PATH}/${PROGRAM_NAME}"
	sshpass -p ${PWD} ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null ${USN}@${IP} "if pgrep ${PROGRAM_NAME}; then pkill -f ${PROGRAM_NAME}; fi"
	sshpass -p ${PWD} ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null ${USN}@${IP} "sh -c 'cd /tmp; ${DEPLOY_PATH}/${PROGRAM_NAME}'" &

zip:
	cp res/* dist && cd dist && zip ${PROGRAM_NAME}.zip config.json icon.png launch.sh ${PROGRAM_NAME}