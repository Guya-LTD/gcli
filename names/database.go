package names

// Databases
const MONGODB_VERSION = "4.4.1-debian-10-r13"
const MONGODB_NAME = "mongodb"
const MONGODB = "bitnami/mongodb"
const MONGODB_VALUE = "mongodb/values.yaml"

const POSTGRESQL_VERSION = "11.9.0-debian-10-r34"
const POSTGRESQL_NAME = "postgresql"
const POSTGRESQL = "bitnami/postgresql"
const POSTGRESQL_VALUE = "postgresql/values.yaml"

const REDIS_VERSION = "6.0.8-debian-10-r0"
const REDIS_NAME = "redis"
const REDIS = "bitnami/redis"
const REDIS_VALUE = "redis/values.yaml"

const MYSQL_VERSION = "6.0.8-debian-10-r0"
const MYSQL_NAME = "redis"
const MYSQL = "bitnami/mysql"
const MYSQL_VALUE = "mysql/values.yaml"

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