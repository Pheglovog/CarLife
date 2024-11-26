export PATH=$PATH:./bin
chmod +x *.sh
./buildnetwork.sh

./createchannel.sh

./deploychaincode.sh