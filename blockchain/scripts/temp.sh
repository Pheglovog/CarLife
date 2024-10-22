TLSCAPEM=organizations/consumer.carlife.com/tlsca/tlsca.consumer.carlife.com-cert.pem
ORGCAPEM=organizations/consumer.carlife.com/ca/ca.consumer.carlife.com-cert.pem
TEMPLATE=scripts/CCP/connect-consumer.yaml
echo "$(yaml_ccp $TEMPLATE $TLSCAPEM $ORGCAPEM)" > organizations/consumer.carlife.com/connect-consumer.yaml