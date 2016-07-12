export HTTP_PROXY=http://proxy.houston.hpecorp.net:8080                    
export HTTPS_PROXY=http://proxy.houston.hpecorp.net:8080

export http_proxy=http://proxy.houston.hpecorp.net:8080                    
export https_proxy=http://proxy.houston.hpecorp.net:8080


#for docker registry to work
REGISTRY_IP=$(docker-machine ip default)
export NO_PROXY=localhost,127.0.0.1,$REGISTRY_IP
export no_proxy=localhost,127.0.0.1,$REGISTRY_IP
