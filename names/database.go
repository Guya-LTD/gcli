package names

const LOCAL_CACHE = "/home/admin/.cache/helm/repository/"

// Databases
const MONGODB_VERSION = "9.0.0"//"4.4.1-debian-10-r13"
const MONGODB_NAME = "mongodb"
//const MONGODB = "bitnami/mongodb"
const MONGODB = LOCAL_CACHE + "mongodb-9.2.4.tgz"
const MONGODB_VALUE = "mongodb/values.dev.yaml"

const POSTGRESQL_VERSION = "11.9.0-debian-10-r34"
const POSTGRESQL_NAME = "postgresql"
//const POSTGRESQL = "bitnami/postgresql"
const POSTGRESQL = LOCAL_CACHE + "postgresql-9.8.4.tgz"
const POSTGRESQL_VALUE = "postgresql/values.dev.yaml"

const REDIS_VERSION = "6.0.8-debian-10-r0"
const REDIS_NAME = "redis"
//const REDIS = "bitnami/redis"
const REDIS = LOCAL_CACHE + "redis-11.2.1.tgz"
const REDIS_VALUE = "redis/values.dev.yaml"

const MYSQL_VERSION = "6.0.8-debian-10-r0"
const MYSQL_NAME = "mysql"
//const MYSQL = "bitnami/mysql"
const MYSQL = LOCAL_CACHE + "mysql-6.14.10.tgz"
const MYSQL_VALUE = "mysql/values.dev.yaml"

const DATABASE_BRANCH_DB = MONGODB
const DATABASE_BRANCH_NAME = "branchdb"
const DATABASE_BRANCH_VALUE = DATABASE_BRANCH_NAME + "/values.yaml"

const DATABASE_CART_DB = MONGODB
const DATABASE_CART_NAME = "cartdb"
const DATABASE_CART_VALUE = DATABASE_CART_NAME + "/values.yaml"

const DATABASE_CATALOG_DB = MONGODB
const DATABASE_CATALOG_NAME = "catalogdb"
const DATABASE_CATALOG_VALUE = DATABASE_CATALOG_NAME + "/values.yaml"

const DATABASE_CHAT_DB = POSTGRESQL
const DATABASE_CHAT_NAME = "chatdb"
const DATABASE_CHAT_VALUE = DATABASE_CHAT_NAME + "/values.yaml"

const DATABASE_CHIPMUNK_DB = ""
const DATABASE_CHIPMUNK_NAME = ""
const DATABASE_CHIPMUNK_VALUES = ""

const DATABASE_DYMO_NAME = ""
const DATABASE_DYMO_VALUES = ""

const DATABASE_GATEKEEPER_DB = REDIS
const DATABASE_GATEKEEPER_NAME = "gatekeeperdb"
const DATABASE_GATEKEEPER_VALUE = DATABASE_GATEKEEPER_NAME + "/values.yaml"

const DATABASE_PAYMENT_DB = MONGODB
const DATABASE_PAYMENT_NAME = "paymentdb"
const DATABASE_PAYMENT_VALUES = DATABASE_PAYMENT_NAME + "/values.yaml"

const DATABASE_XPRESS_DB = MYSQL
const DATABASE_XPRESS_NAME = "xpressdb"
const DATABASE_XPRESS_VALUE = DATABASE_XPRESS_NAME + "/values.yaml"

const DATABASE_XTRACK_DB = ""
const DATABASE_XTRACK_NAME = ""
const DATABASE_XTRACK_VALUE = ""