#!/bin/bash

######################################################################
# VARIABLES
######################################################################

NETWORK_NAME=generate_certs_network
CONTAINER_NAME=make_certs
CRYPTO_PATH=$PWD/crypto
OUTPUT_DIR=$PWD/output

######################################################################
# START CA CONTAINER
######################################################################

echo "Create docker network..."
NETWORK_IDS=$(docker network ls -f="name=$NETWORK_NAME" -q)

if [ -z "$NETWORK_IDS" -o "$NETWORK_IDS" == " " ]; then
    echo "Docker network:"
    docker network create --attachable --driver bridge $NETWORK_NAME
else
    echo "Network $NETWORK_NAME already exists"
fi

echo "Create and start CA container..."
docker run -it -d --network="$NETWORK_NAME" --name $CONTAINER_NAME --restart always \
    -p 7054:7054 \
    -e FABRIC_CA_HOME=/etc/hyperledger/fabric-ca-server \
    -e FABRIC_CA_SERVER_CA_NAME=platform-ca \
    -e FABRIC_CA_SERVER_TLS_ENABLED=false \
    -e FABRIC_CA_SERVER_CA_CERTFILE=/etc/hyperledger/fabric-ca-server-config/ca.settlementplatform.io-cert.pem \
    -e FABRIC_CA_SERVER_CA_KEYFILE=/etc/hyperledger/fabric-ca-server-config/0d3dde5e796c61219aa6c8e5b2625e4f21f06769a132797f32d232c499bdb284_sk \
    -v $CRYPTO_PATH:/etc/hyperledger/fabric-ca-server-config \
    "hyperledger/fabric-ca:1.4" \
    sh -c "fabric-ca-server start -b admin:adminpw -d"

echo "Waiting while it is starting..."
sleep 3
echo

######################################################################
# ENROLL CA ADMIN CERTFICATE
######################################################################

echo "Enroll admin"
docker exec $CONTAINER_NAME fabric-ca-client enroll -u http://admin:adminpw@localhost:7054

######################################################################
# ADD AFFILIATIONS
######################################################################

echo
echo "Add participant affiliation"
docker exec $CONTAINER_NAME fabric-ca-client affiliation add platform_participants
docker exec $CONTAINER_NAME fabric-ca-client affiliation add platform_admins
docker exec $CONTAINER_NAME fabric-ca-client affiliation add platform_kyc_officers
docker exec $CONTAINER_NAME fabric-ca-client affiliation add kyc_agents
echo
echo "Available affiliations:"
docker exec $CONTAINER_NAME fabric-ca-client affiliation list

######################################################################
# REGISTER USERS
######################################################################

echo
echo "REGISTER USERS"
echo "=============="

echo "Register platform participant"
REGISTER_RESULT="$(docker exec $CONTAINER_NAME fabric-ca-client register --id.name platform_participant --id.affiliation platform_participants)"
PARTICIPANT_PASSWORD=$(echo $REGISTER_RESULT| cut -d' ' -f 2)

echo "Register platform admin"
REGISTER_RESULT="$(docker exec $CONTAINER_NAME fabric-ca-client register --id.name platform_admin --id.affiliation platform_admins)"
ADMIN_PASSWORD=$(echo $REGISTER_RESULT| cut -d' ' -f 2)

echo "Register KYC officer"
REGISTER_RESULT="$(docker exec $CONTAINER_NAME fabric-ca-client register --id.name platform_kyc_officer --id.affiliation platform_kyc_officers)"
KYC_OFFICER_PASSWORD=$(echo $REGISTER_RESULT| cut -d' ' -f 2)

echo "Register KYC agent"
REGISTER_RESULT="$(docker exec $CONTAINER_NAME fabric-ca-client register --id.name kyc_agent --id.affiliation kyc_agents)"
KYC_AGENT_PASSWORD=$(echo $REGISTER_RESULT| cut -d' ' -f 2)

######################################################################
# ENROLL USER CERTIFICATES
######################################################################

[ ! -d $OUTPUT_DIR ] && mkdir $OUTPUT_DIR

echo
echo "ENROLL USER CERTIFICATES"
echo "========================"

echo "Enroll participant's certificate with password " + $PARTICIPANT_PASSWORD
docker exec $CONTAINER_NAME fabric-ca-client enroll -u http://platform_participant:$PARTICIPANT_PASSWORD@localhost:7054 --csr.names C=RU,O=platform,OU=platform_participants,ST=Moscow
echo "Save certificate to file platform_participant.pem"
docker exec $CONTAINER_NAME cat /etc/hyperledger/fabric-ca-server/msp/signcerts/cert.pem > $OUTPUT_DIR/platform_participant.pem

echo "Enroll admin's certificate with password " + $ADMIN_PASSWORD
docker exec $CONTAINER_NAME fabric-ca-client enroll -u http://platform_admin:$ADMIN_PASSWORD@localhost:7054 --csr.names C=RU,O=platform,OU=platform_admins,ST=Moscow
echo "Save certificate to file platform_admin.pem"
docker exec $CONTAINER_NAME cat /etc/hyperledger/fabric-ca-server/msp/signcerts/cert.pem > $OUTPUT_DIR/platform_admin.pem

echo "Enroll KYC officer's certificate with password " + $KYC_OFFICER_PASSWORD
docker exec $CONTAINER_NAME fabric-ca-client enroll -u http://platform_kyc_officer:$KYC_OFFICER_PASSWORD@localhost:7054 --csr.names C=RU,O=platform,OU=platform_kyc_officers,ST=Moscow
echo "Save certificate to file platform_kyc_officer.pem"
docker exec $CONTAINER_NAME cat /etc/hyperledger/fabric-ca-server/msp/signcerts/cert.pem > $OUTPUT_DIR/platform_kyc_officer.pem

echo "Enroll KYC agent's certificate with password " + $KYC_AGENT_PASSWORD
docker exec $CONTAINER_NAME fabric-ca-client enroll -u http://kyc_agent:$KYC_AGENT_PASSWORD@localhost:7054 --csr.names C=RU,O=platform,OU=kyc_agents,ST=Moscow
echo "Save certificate to file kyc_agent.pem"
docker exec $CONTAINER_NAME cat /etc/hyperledger/fabric-ca-server/msp/signcerts/cert.pem > $OUTPUT_DIR/kyc_agent.pem

######################################################################
# CLEAR CERTIFICATES
######################################################################

echo
echo "Stop working and clear resources"
docker stop $CONTAINER_NAME && docker rm $CONTAINER_NAME
docker network rm $NETWORK_NAME
