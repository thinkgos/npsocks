# 应用名称
name = npsocks
# 型号
model = ${name}
# 固件版本
# git describe --tags `git rev-list --tags --max-count=1`
version = `git describe --tags`
# api版本
apiVersion = v0.0.1
# 设置固件名称
firmwareName = ${name}

execveFile := ${firmwareName}

# 路径相关
ProjectDir=.
BinDir=${CURDIR}/bin

# 编译平台
platform = CGO_ENABLED=0
ifeq (${MAKECMDGOALS},windows)
	platform += GOOS=windows GOARCH=amd64
	execveFile = ${firmwareName}.exe
endif
# 编译选项,如tags,多个采用','分开 noswag
opts = -trimpath
# 编译flags
path = github.com/thinkgos/go-core-package/builder
flags = -ldflags "-X '${path}.BuildTime=`date "+%F %T %z"`' \
	-X '${path}.GitCommit=`git rev-parse --short=8 HEAD`' \
	-X '${path}.GitFullCommit=`git rev-parse HEAD`' \
	-X '${path}.Name=${name}' \
	-X '${path}.Model=${model}' \
	-X '${path}.Version=${version}' \
	-X '${path}.APIVersion=${apiVersion}' -w" # -s 会导致gops识别不了版本

system:
	@echo "----> system executable building..."
	@mkdir -p ${BinDir}
	@${platform} go build ${opts} ${flags} -o ${BinDir}/${execveFile} ${ProjectDir}/main.go
	@#upx --best --lzma ${BinDir}/${execveFile}
	@#bzip2 -c ${BinDir}/${execveFile} > ${BinDir}/${execveFile}.bz2
	@echo "----> system executable build successful"

windows: system

run: system
	@${BinDir}/${execveFile} server

clean:
	@echo "----> cleaning..."
	@go clean
	@rm -r ${BinDir}
	@echo "----> clean successful"

help:
	@echo " ------------- How to build ------------- "
	@echo " make         -- build target for system"
	@echo " run          -- build and run target for system"
	@echo " make clean   -- clean build files"
	@echo " ------------- How to build ------------- "

.PHONY: system run clean help



