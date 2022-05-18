pkill -9 "servercode"
cd clientcode
go build clientcode.go
cd ..
cd servercode
go build servercode.go
cd ..
clientcode/clientcode  -create -id 1020
clientcode/clientcode  -create -id 5060
clientcode/clientcode  -create -id 43
clientcode/clientcode  -create -id 2343
clientcode/clientcode  -create -id 535
clientcode/clientcode  -write -id 1020 -name token1 -low 0 -mid 10 -high 100
clientcode/clientcode  -write -id 2343 -name token2 -low 36 -mid 335 -high 503
clientcode/clientcode  -write -id 43 -name token3 -low 123 -mid 434 -high 1000
clientcode/clientcode  -write -id 5060 -name token4 -low 43 -mid 56 -high 123
clientcode/clientcode  -write -id 535 -name token5 -low 6743 -mid 7543 -high 8522
clientcode/clientcode  -read -id 1020
clientcode/clientcode  -read -id 2343
clientcode/clientcode  -read -id 5060
clientcode/clientcode  -read -id 43
clientcode/clientcode  -read -id 535
echo "Program sleeping for 5 seconds to allow for fail-silent emulation for token 1020 on server running on port 65000"
sleep 5
echo "Refer log/log_65000.log for more details"
clientcode/clientcode  -read -id 1020
clientcode/clientcode  -read -id 1020
clientcode/clientcode  -read -id 1020
clientcode/clientcode  -read -id 1020
clientcode/clientcode  -read -id 1020
clientcode/clientcode  -read -id 1020
clientcode/clientcode  -read -id 1020
clientcode/clientcode  -drop -id 2343
clientcode/clientcode  -drop -id 1020
clientcode/clientcode  -drop -id 535
clientcode/clientcode  -drop -id 5060
clientcode/clientcode  -drop -id 43
pkill -9 "servercode"
