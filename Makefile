PROJECT_NAME := tsp-rubik
PROGRAM_NAME := rubik
DEPLOY_PATH := /mnt/SDCARD/Apps/Rubik

IP := 192.168.0.101
USN := root
PWD := tina

all: clean docker deploy

clean:
	rm bin/${PROGRAM_NAME} -f

docker:
	#docker run -d --name go-aarch64 -c 1024 -it --volume=/home/vitaly/GolandProjects/:/work/ --workdir=/work/ rust-aarch64
	docker exec go-aarch64 /bin/bash -c 'cd ${PROJECT_NAME} && make build'

build:
	rm bin/${PROGRAM_NAME} -f #docker container caches the binary
	go build -o bin/${PROGRAM_NAME} ${PROJECT_NAME}/src/

deploy:
	sshpass -p ${PWD} ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null ${USN}@${IP} "rm ${DEPLOY_PATH}/${PROGRAM_NAME} -f"
	sshpass -p ${PWD} scp -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null bin/${PROGRAM_NAME} ${USN}@${IP}:${DEPLOY_PATH}
	sshpass -p ${PWD} ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null ${USN}@${IP} "chmod 777 ${DEPLOY_PATH}/${PROGRAM_NAME}"
	sshpass -p ${PWD} ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null ${USN}@${IP} "if pgrep ${PROGRAM_NAME}; then pkill -f ${PROGRAM_NAME}; fi"
	sshpass -p ${PWD} ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null ${USN}@${IP} "sh -c 'cd /tmp; ${DEPLOY_PATH}/${PROGRAM_NAME}'" &
