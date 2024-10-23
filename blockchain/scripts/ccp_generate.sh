#!/bin/bash

function one_line_pem {
    echo "`awk 'NF {sub(/\\n/, ""); printf "%s\\\\\\\n",$0;}' $1`"
}

function yaml_ccp {
    local TP=$(one_line_pem $2)
    local OP=$(one_line_pem $3)
    sed -e "s#\${TLSCAPEM}#$TP#" \
        -e "s#\${ORGCAPEM}#$OP#" \
        $1 | sed -e $'s/\\\\n/\\\n          /g'
}

TLSCAPEM=organizations/component_supplier.carlife.com/tlsca/tlsca.component_supplier.carlife.com-cert.pem
ORGCAPEM=organizations/component_supplier.carlife.com/ca/ca.component_supplier.carlife.com-cert.pem
TEMPLATE=scripts/CCP/connect-component_supplier.yaml
echo "$(yaml_ccp $TEMPLATE $TLSCAPEM $ORGCAPEM)" > organizations/component_supplier.carlife.com/connect-component_supplier.yaml

TLSCAPEM=organizations/manufacturer.carlife.com/tlsca/tlsca.manufacturer.carlife.com-cert.pem
ORGCAPEM=organizations/manufacturer.carlife.com/ca/ca.manufacturer.carlife.com-cert.pem
TEMPLATE=scripts/CCP/connect-manufacturer.yaml
echo "$(yaml_ccp $TEMPLATE $TLSCAPEM $ORGCAPEM)" > organizations/manufacturer.carlife.com/connect-manufacturer.yaml

TLSCAPEM=organizations/store.carlife.com/tlsca/tlsca.store.carlife.com-cert.pem
ORGCAPEM=organizations/store.carlife.com/ca/ca.store.carlife.com-cert.pem
TEMPLATE=scripts/CCP/connect-store.yaml
echo "$(yaml_ccp $TEMPLATE $TLSCAPEM $ORGCAPEM)" > organizations/store.carlife.com/connect-store.yaml

TLSCAPEM=organizations/insurer.carlife.com/tlsca/tlsca.insurer.carlife.com-cert.pem
ORGCAPEM=organizations/insurer.carlife.com/ca/ca.insurer.carlife.com-cert.pem
TEMPLATE=scripts/CCP/connect-insurer.yaml
echo "$(yaml_ccp $TEMPLATE $TLSCAPEM $ORGCAPEM)" > organizations/insurer.carlife.com/connect-insurer.yaml

TLSCAPEM=organizations/maintenancer.carlife.com/tlsca/tlsca.maintenancer.carlife.com-cert.pem
ORGCAPEM=organizations/maintenancer.carlife.com/ca/ca.maintenancer.carlife.com-cert.pem
TEMPLATE=scripts/CCP/connect-maintenancer.yaml
echo "$(yaml_ccp $TEMPLATE $TLSCAPEM $ORGCAPEM)" > organizations/maintenancer.carlife.com/connect-maintenancer.yaml

TLSCAPEM=organizations/consumer.carlife.com/tlsca/tlsca.consumer.carlife.com-cert.pem
ORGCAPEM=organizations/consumer.carlife.com/ca/ca.consumer.carlife.com-cert.pem
TEMPLATE=scripts/CCP/connect-consumer.yaml
echo "$(yaml_ccp $TEMPLATE $TLSCAPEM $ORGCAPEM)" > organizations/consumer.carlife.com/connect-consumer.yaml