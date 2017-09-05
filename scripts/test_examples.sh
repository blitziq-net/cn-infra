#!/usr/bin/env bash

TMP_FILE="/tmp/out"
exitCode=0
PREV_IFS="$IFS"

# test whether output of the command contains expected lines
# arguments
# 1-st command to run
# 2-nd array of expected strings in the

function testOutput {
IFS="${PREV_IFS}"

    #run the command
    $1 > ${TMP_FILE} 2>&1

IFS="
"
    echo "Testing $1"
    rv=0
    # loop through expected lines
    for i in $2
    do
        if grep "${i}" /tmp/out > /dev/null ; then
            echo "OK - '$i'"
        else
            echo "Not found - '$i'"
            rv=1
        fi
    done

    # if an error occurred print the output
    if [[ ! $rv -eq 0 ]] ; then
        cat ${TMP_FILE}
        exitCode=1
    fi

    echo "================================================================"
    rm ${TMP_FILE}
    return ${rv}
}

function startEtcd {
    docker run -p 2379:2379 --name etcd -d -e ETCDCTL_API=3 \
        quay.io/coreos/etcd:v3.1.0 /usr/local/bin/etcd \
             -advertise-client-urls http://0.0.0.0:2379 \
                 -listen-client-urls http://0.0.0.0:2379 > /dev/null
    # dump etcd content to make sure that etcd is ready
    docker exec etcd etcdctl get --prefix ""
}

function stopEtcd {
    docker stop etcd > /dev/null
    docker rm etcd > /dev/null
}

function startKafka {
    docker run -p 2181:2181 -p 9092:9092 --name kafka -d \
 --env ADVERTISED_HOST=0.0.0.0 --env ADVERTISED_PORT=9092 spotify/kafka > /dev/null
    # list kafka topics to ensure that kafka is ready
    docker exec kafka  /opt/kafka_2.11-0.10.1.0/bin/kafka-topics.sh --list --zookeeper localhost:2181 > /dev/null 2> /dev/null

}

function stopKafka {
    docker stop kafka > /dev/null
    docker rm kafka > /dev/null
}

#### Logging #############################################################
expected=("Debug log example
Info log example
Warn log example
Error log example
Stopping agent...
")

testOutput examples/logs_plugin/logs_plugin "${expected}"

#### Etcd #############################################################

startEtcd
expected=("Saving  /phonebook/Peter
")
cmd="examples/etcdv3_lib/editor/editor  --cfg examples/etcdv3_lib/etcd.conf  put  Peter Company 0907"
testOutput "${cmd}" "${expected}"

stopEtcd

#### Kafka #############################################################
startKafka

expected=("Kafka connecting
Consuming started
Sync published
Message is stored in topic(test)/partition(0)/offset(1)
")

testOutput examples/kafka_lib/mux/mux "${expected}"

stopKafka

########################################################################

exit ${exitCode}
