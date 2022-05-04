#!/bin/sh
echo "正在编译: Aopo—Linux-amd..."
sudo CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=x86_64-linux-musl-gcc CXX=x86_64-linux-musl-g++ go build -ldflags "-s -w" -o ./mark/Aopo_Linux_amd
echo "正在编译: Aopo_Windows_x86..."
sudo CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ go build -ldflags "-s -w" -o ./mark/Aopo_Windows_x86.exe
echo "正在编译: Aopo—Mac-M1..."
sudo CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o ./mark/Aopo_M1
echo "正在编译: Aopo—Mac-Intel..."
sudo CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o ./mark/Aopo_Intel
#echo "正在移除编译信息并混淆: Aopo_Intel..."
#sudo ./mark/go-strip -f ./mark/Aopo_Intel -a -output ./mark/Aopo_Intel_new
#echo "正在移除编译信息并混淆: Aopo_Windows_x86..."
#sudo ./mark/go-strip -f ./mark/Aopo_Windows_x86.exe -a -output ./mark/Aopo_Windows_x86_new.exe
#echo "正在移除编译信息并混淆: Aopo_Linux_amd..."
#sudo ./mark/go-strip -f ./mark/Aopo_Linux_amd -a -output ./mark/Aopo_Linux_amd_new
echo "正在UPX加壳减少体积: Aopo_Intel..."
sudo upx ./mark/Aopo_Intel_new
echo "正在UPX加壳减少体积: Aopo_M1..."
sudo upx ./mark/Aopo_M1
echo "正在UPX加壳减少体积: Aopo_Windows_x86..."
sudo upx ./mark/Aopo_Windows_x86_new.exe
echo "正在UPX加壳减少体积: Aopo_Linux_amd..."
sudo upx ./mark/Aopo_Linux_amd_new
sudo mv ./mark/Aopo_Intel_new ./mark/Aopo_Intel
sudo mv ./mark/Aopo_Windows_x86_new.exe ./mark/Aopo_Windows_x86.exe
sudo mv ./mark/Aopo_Linux_amd_new ./mark/Aopo_Linux_amd
sudo chmod 777 ./mark/Aopo_Intel
sudo chmod 777 ./mark/Aopo_Windows_x86.exe
sudo chmod 777 ./mark/Aopo_Linux_amd
echo "=======================================END结束======================================="