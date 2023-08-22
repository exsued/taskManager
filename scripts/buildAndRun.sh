#!/bin/bash
argsfileName='./privacy/args.txt'
read args < $argsfileName

projectName=`pwd | rev | cut -d"/" -f1 | rev`
cd ./src

if go build -o ${projectName}; then
    workDir=`pwd`
    mate-terminal --working-directory="${workDir}" --command "bash -c '(./${projectName} ${args})'"
    cd ..
    mv ./src/${projectName} ./bin/${projectName}
fi