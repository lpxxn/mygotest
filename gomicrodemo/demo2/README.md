
cd /Users/li/go/src/github.com/mygotest/gomicrodemo/demo2/
protoc --proto_path=$GOPATH/src/github.com/mygotest/gomicrodemo/demo2/proto:. --micro_out=. --go_out=. greeter.proto
protoc --proto_path=$GOPATH/src/github.com/mygotest/gomicrodemo/demo2/proto:. --micro_out=. --go_out=. ./proto/*.proto



或者进入
cd /Users/li/go/src/github.com/mygotest/gomicrodemo/demo2/proto

protoc --proto_path=.:. --micro_out=. --go_out=. greeter.proto

protoc --proto_path=.:. --micro_out=. --go_out=. ./*.proto



Seems like etcd is configured to run in localhost. Please make sure ports and data directory do not conflict!


# after transferring certs to remote machines
mkdir -p ${HOME}/certs
cp /tmp/certs/* ${HOME}/certs


# make sure etcd process has write access to this directory
# remove this directory if the cluster is new; keep if restarting etcd
# rm -rf /tmp/etcd/s1


/tmp/test-etcd/etcd --name s1 \
  --data-dir /tmp/etcd/s1 \
  --listen-client-urls https://localhost:2379 \
  --advertise-client-urls https://localhost:2379 \
  --listen-peer-urls https://localhost:2380 \
  --initial-advertise-peer-urls https://localhost:2380 \
  --initial-cluster s1=https://localhost:2380,s2=https://localhost:22380,s3=https://localhost:32380 \
  --initial-cluster-token tkn \
  --initial-cluster-state new \
  --client-cert-auth \
  --trusted-ca-file ${HOME}/certs/etcd-root-ca.pem \
  --cert-file ${HOME}/certs/s1.pem \
  --key-file ${HOME}/certs/s1-key.pem \
  --peer-client-cert-auth \
  --peer-trusted-ca-file ${HOME}/certs/etcd-root-ca.pem \
  --peer-cert-file ${HOME}/certs/s1.pem \
  --peer-key-file ${HOME}/certs/s1-key.pem




# after transferring certs to remote machines
mkdir -p ${HOME}/certs
cp /tmp/certs/* ${HOME}/certs


# make sure etcd process has write access to this directory
# remove this directory if the cluster is new; keep if restarting etcd
# rm -rf /tmp/etcd/s2


/tmp/test-etcd/etcd --name s2 \
  --data-dir /tmp/etcd/s2 \
  --listen-client-urls https://localhost:22379 \
  --advertise-client-urls https://localhost:22379 \
  --listen-peer-urls https://localhost:22380 \
  --initial-advertise-peer-urls https://localhost:22380 \
  --initial-cluster s1=https://localhost:2380,s2=https://localhost:22380,s3=https://localhost:32380 \
  --initial-cluster-token tkn \
  --initial-cluster-state new \
  --client-cert-auth \
  --trusted-ca-file ${HOME}/certs/etcd-root-ca.pem \
  --cert-file ${HOME}/certs/s2.pem \
  --key-file ${HOME}/certs/s2-key.pem \
  --peer-client-cert-auth \
  --peer-trusted-ca-file ${HOME}/certs/etcd-root-ca.pem \
  --peer-cert-file ${HOME}/certs/s2.pem \
  --peer-key-file ${HOME}/certs/s2-key.pem




# after transferring certs to remote machines
mkdir -p ${HOME}/certs
cp /tmp/certs/* ${HOME}/certs


# make sure etcd process has write access to this directory
# remove this directory if the cluster is new; keep if restarting etcd
# rm -rf /tmp/etcd/s3


/tmp/test-etcd/etcd --name s3 \
  --data-dir /tmp/etcd/s3 \
  --listen-client-urls https://localhost:32379 \
  --advertise-client-urls https://localhost:32379 \
  --listen-peer-urls https://localhost:32380 \
  --initial-advertise-peer-urls https://localhost:32380 \
  --initial-cluster s1=https://localhost:2380,s2=https://localhost:22380,s3=https://localhost:32380 \
  --initial-cluster-token tkn \
  --initial-cluster-state new \
  --client-cert-auth \
  --trusted-ca-file ${HOME}/certs/etcd-root-ca.pem \
  --cert-file ${HOME}/certs/s3.pem \
  --key-file ${HOME}/certs/s3-key.pem \
  --peer-client-cert-auth \
  --peer-trusted-ca-file ${HOME}/certs/etcd-root-ca.pem \
  --peer-cert-file ${HOME}/certs/s3.pem \
  --peer-key-file ${HOME}/certs/s3-key.pem




ETCDCTL_API=3 /tmp/test-etcd/etcdctl \
  --endpoints localhost:2379,localhost:22379,localhost:32379 \
  --cacert ${HOME}/certs/etcd-root-ca.pem \
  --cert ${HOME}/certs/s1.pem \
  --key ${HOME}/certs/s1-key.pem \
  endpoint health


ETCDCTL_API=3 /tmp/test-etcd/etcdctl \
  --endpoints localhost:2379,localhost:22379,localhost:32379 \
  --cacert ${HOME}/certs/etcd-root-ca.pem \
  --cert ${HOME}/certs/s1.pem \
  --key ${HOME}/certs/s1-key.pem \
 endpoint status

ETCDCTL_API=3 /tmp/test-etcd/etcdctl \
 -w table \
  --endpoints localhost:2379,localhost:22379,localhost:32379 \
  --cacert ${HOME}/certs/etcd-root-ca.pem \
  --cert ${HOME}/certs/s1.pem \
  --key ${HOME}/certs/s1-key.pem \
 endpoint status


/tmp/test-etcd/etcdctl \
  --endpoints https://localhost:2379,https://localhost:22379,https://localhost:32379 \
  --ca-file ${HOME}/certs/etcd-root-ca.pem \
  --cert-file ${HOME}/certs/s1.pem \
  --key-file ${HOME}/certs/s1-key.pem \
 cluster-health


/tmp/test-etcd/etcdctl \
  --endpoint https://localhost:2379 \
  --ca-file ${HOME}/certs/etcd-root-ca.pem \
  --cert-file ${HOME}/certs/s1.pem \
  --key-file ${HOME}/certs/s1-key.pem \
 cluster-health

ETCDCTL_API=3 /tmp/test-etcd/etcdctl \
  --endpoints localhost:2379,localhost:22379,localhost:32379 \
  --cacert ${HOME}/certs/etcd-root-ca.pem \
  --cert ${HOME}/certs/s1.pem \
  --key ${HOME}/certs/s1-key.pem \
put mykey "this is awesome! lp"

ETCDCTL_API=3 /tmp/test-etcd/etcdctl \
  --endpoints localhost:2379,localhost:22379,localhost:32379 \
  --cacert ${HOME}/certs/etcd-root-ca.pem \
  --cert ${HOME}/certs/s1.pem \
  --key ${HOME}/certs/s1-key.pem \
get mykey

ETCDCTL_API=3 /tmp/test-etcd/etcdctl \
  --endpoints localhost:2379,localhost:22379,localhost:32379 \
  --cacert ${HOME}/certs/etcd-root-ca.pem \
  --cert ${HOME}/certs/s1.pem \
  --key ${HOME}/certs/s1-key.pem \
 ls


ETCDCTL_API=3 /tmp/test-etcd/etcdctl \
  --endpoints localhost:2379,localhost:22379,localhost:32379 \
  --cacert ${HOME}/certs/etcd-root-ca.pem \
  --cert ${HOME}/certs/s1.pem \
  --key ${HOME}/certs/s1-key.pem \
member list


得到所有的keys

ETCDCTL_API=3 /tmp/test-etcd/etcdctl \
  --endpoints localhost:2379,localhost:22379,localhost:32379 \
  --cacert ${HOME}/certs/etcd-root-ca.pem \
  --cert ${HOME}/certs/s1.pem \
  --key ${HOME}/certs/s1-key.pem \
get  --from-key ''

ETCDCTL_API=3 /tmp/test-etcd/etcdctl \
  --endpoints localhost:2379,localhost:22379,localhost:32379 \
  --cacert ${HOME}/certs/etcd-root-ca.pem \
  --cert ${HOME}/certs/s1.pem \
  --key ${HOME}/certs/s1-key.pem \
get "" --prefix=true

监视watch
ETCDCTL_API=3 /tmp/test-etcd/etcdctl \
  --endpoints localhost:2379,localhost:22379,localhost:32379 \
  --cacert ${HOME}/certs/etcd-root-ca.pem \
  --cert ${HOME}/certs/s1.pem \
  --key ${HOME}/certs/s1-key.pem \
 watch --prefix /micro




ETCDCTL_API=3 /tmp/test-etcd/etcdctl \
  --endpoints localhost:2379,localhost:22379,localhost:32379 \
  --cacert ${HOME}/certs/etcd-root-ca.pem \
  --cert ${HOME}/certs/s1.pem \
  --key ${HOME}/certs/s1-key.pem \
get "" --from-key
