#!/bin/bash
argsfileName='./privacy/args.txt'
read args < $argsfileName

projectName=`pwd | rev | cut -d"/" -f1 | rev`
cd ./src

if go build -o ${projectName}; then
    cd ..
    mv ./src/${projectName} ./bin/${projectName}
    cd ./bin/
    (./${projectName} ${args}) > out.txt
fi
