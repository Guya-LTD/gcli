package names

const ELASTICSEARCH = LOCAL_CACHE +  "elastic/elasticsearch"
//const ELASTICSEARCH = "elastic/elasticsearch"
const ELASTICSEARCH_VERSION = "7.9.1"
const ELASTICSEARCH_DEPLOYMENT_NAME = "elasticsearch"
const ELASTICSEARCH_DEPLOYMENT_DIR = "elasticsearch"
const ELASTICSEARCH_DEPLOYMENT_VALUE = ELASTICSEARCH_DEPLOYMENT_DIR + "/values.dev.yaml"
const ELASTICSEARCH_DEPLOYMENT_LOCAL_STORAGE_VALUE = ELASTICSEARCH_DEPLOYMENT_DIR + "/local-path-storage.yaml"

const LOGSTASH = LOCAL_CACHE + "logstash-7.9.1.tgz"
//const LOGSTASH = "elastic/logstash"
const LOGSTASH_DEPLOYMENT_VERSION = "7.9.1"
const LOGSTASH_DEPLOYMENT_NAME = "logstash"
const LOGSTASH_DEPLOYMENT_DIR = "logstash"
const LOGSTASH_DEPLOYMENT_VALUE = LOGSTASH_DEPLOYMENT_DIR + "/values.dev.yaml"

const KIBANA = LOCAL_CACHE + "kibana-7.9.1.tgz"
//const KIBANA = "elastic/kibana"
const KIBANA_DEPLOYMENT_VERSION = "7.9.1"
const KIBANA_DEPLOYMENT_NAME = "kibana"
const KIBANA_DEPLOYMENT_DIR = "kibana"
const KIBANA_DEPLOYMENT_VALUE = KIBANA_DEPLOYMENT_DIR + "/values.dev.yaml"

const RABBITMQ = LOCAL_CACHE + "rabbitmq-7.6.8.tgz"
//const RABBITMQ = "bitnami/rabbitmq"
const RABBITMQ_NAME = "rabbitmq"