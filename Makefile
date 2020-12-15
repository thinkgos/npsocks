# 应用名称
name = onlys
# 型号
model = ${name}
# 固件版本
# git describe --tags `git rev-list --tags --max-count=1`
version = v0.0.1 #`git describe --tags`
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
# 编译选项,如tags,多个采用','分开 noswag
opts = -trimpath
# 编译flags
path = github.com/thinkgos/only-socks5/pkg/builder
flags = -ldflags "-X '${path}.BuildTime=`date "+%F %T %z"`' \
	-X '${path}.GitCommit=`git rev-parse --short=8 HEAD`' \
	-X '${path}.GitFullCommit=`git rev-parse HEAD`' \
	-X '${path}.Name=${name}' \
	-X '${path}.Model=${model}' \
	-X '${path}.Version=${version}' \
	-X '${path}.APIVersion=${apiVersion}' -s -w"

system:
	@echo "----> system executable building..."
	@mkdir -p ${BinDir}
	@${platform} go build ${opts} ${flags} -o ${BinDir}/${execveFile} ${ProjectDir}/main.go
	@#upx --best --lzma ${BinDir}/${execveFile}
	@#bzip2 -c ${BinDir}/${execveFile} > ${BinDir}/${execveFile}.bz2
	@echo "----> system executable build successful"

run: system
	@${BinDir}/${execveFile} server

prepare:
	@go generate ${ProjectDir}/...

clean:
	@echo "----> cleaning..."
	@go clean
	@rm -r ${BinDir}
	@echo "----> clean successful"

help:
	@echo " ------------- How to build ------------- "
	@echo " make         -- build target for system"
	@echo " run          -- build and run target for system"
	@echo " make prepare -- go generate"
	@echo " make clean   -- clean build files"
	@echo " ------------- How to build ------------- "

.PHONY: system run prepare clean help



