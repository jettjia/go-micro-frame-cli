package gogs

import (
	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/jettjia/go-micro-frame-cli/util"
)

// RunGogs 安装 gogs
func RunGogs() {
	installMysql()
	installGogs()
}

func installMysql() {
	gproc.ShellExec("sudo mkdir -p /data/mysql-01/conf")
	gproc.ShellExec("sudo mkdir -p /data/mysql-01/data")

	mysqlConf := `[client]
port= 3306
socket  = /tmp/mysql.sock
#default-character-set = utf8mb4

## The MySQL server
[mysqld]
port = 3306
socket  = /tmp/mysql.sock
user = mysql
skip-external-locking
skip-name-resolve
#skip-grant-tables
#skip-networking
###################################### dir
#basedir=/usr/local/mysql
datadir=/var/lib/mysql
tmpdir=/var/lib/mysql
secure_file_priv=/var/lib/mysql
###################################### some app
log-error=mysql.err
pid-file=/var/lib/mysql/mysql.pid
local-infile=1
event_scheduler=0
federated
default-storage-engine=InnoDB
#default-time-zone= '+8:00'
log_timestamps=SYSTEM
character-set-client-handshake = FALSE
character-set-server = utf8mb4
collation-server = utf8mb4_unicode_ci
init_connect='SET NAMES utf8mb4'
#5.6
explicit_defaults_for_timestamp=true

#fulltext
innodb_optimize_fulltext_only
ft_min_word_len=1
#ft_max_word_len
innodb_ft_min_token_size=1

###################################### memory allocate and myisam configure
max_connections=3000
#back_log=200
max_connect_errors=10000

key_buffer_size = 16M
max_allowed_packet = 16M
table_open_cache = 10240
sort_buffer_size = 2M
read_buffer_size = 2M
read_rnd_buffer_size = 2M
join_buffer_size=2M
myisam_sort_buffer_size = 4M
#net_buffer_length = 2M
thread_cache_size = 24

query_cache_type=1
query_cache_size=256M
query_cache_limit=32M

tmp_table_size=1G
max_heap_table_size=1G

#thread_concurrency =48
###################################### replication
server-id = 10951
log-bin=mysql-bin
binlog_format=mixed
max_binlog_size=1G
#binlog_cache_size=512M
log_slave_updates=true
log_bin_trust_function_creators=true
expire_logs_days=15
replicate-ignore-db=mysql
replicate-ignore-db=test
replicate-ignore-db=information_schema
replicate-ignore-db=performance_schema
replicate-wild-ignore-table=mysql.%
replicate-wild-ignore-table=test.%
replicate-wild-ignore-table=information_schema.%
replicate-wild-ignore-table=performance_schema.%

lower_case_table_names = 1
#read_only=1
master_info_repository=TABLE
relay_log_info_repository=TABLE

###################################### slow-query
long_query_time=1
slow_query_log=1
slow_query_log_file=/var/lib/mysql/slow-query.log
interactive_timeout=600
wait_timeout=600
#log_queries_not_using_indexes=1

###################################### innodb configure
innodb_file_per_table
#innodb_file_format=Barracuda
#innodb_io_capacity=200

innodb_data_home_dir = /var/lib/mysql
#innodb_data_file_path = ibdata1:2000M;ibdata2:10M:autoextend
innodb_log_group_home_dir = /var/lib/mysql
innodb_buffer_pool_size =4G
# Set .._log_file_size to 25 % of buffer pool size
innodb_log_file_size = 1G
innodb_log_files_in_group = 3
innodb_log_buffer_size = 32M
#innodb_lock_wait_timeout = 50
innodb_flush_log_at_trx_commit = 1
sync_binlog=0
sql-mode="STRICT_TRANS_TABLES,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION"
##########################################
[mysqldump]
quick
max_allowed_packet = 16M

[mysql]
no-auto-rehash
default-character-set = utf8mb4
prompt=\\U \\h \\R:\\m:\\s \\d>

[myisamchk]
key_buffer_size = 20M
sort_buffer_size = 20M
read_buffer = 2M
write_buffer = 2M

[mysqlhotcopy]
interactive-timeout`

	util.WriteStringToFileMethod("/data/mysql-01/conf/my.cnf", mysqlConf)

	// install mysql
	_, err := gproc.ShellExec(`docker run -d \
-p 3306:3306 \
-v /etc/localtime:/etc/localtime:ro \
-v /data/mysql-01/data:/var/lib/mysql \
-v /data/mysql-01/conf:/etc/mysql/conf.d \
-v /data/mysql-01/conf/my.cnf:/etc/mysql/my.cnf \
--env MYSQL_ROOT_PASSWORD=123456 \
--name mysql5.7-01 mysql:5.7`)
	if err != nil {
		mlog.Fatal("docker run mysql: ", err)
		return
	}

	mlog.Print("mysql done!")
}

func installGogs() {
	_, err := gproc.ShellExec(`docker run -d \
--name=gogs \
-p 10022:22 -p 3000:3000 \
-v /data/gogs:/data \
gogs/gogs`)

	if err != nil {
		mlog.Fatal("docker run gogos: ", err)
		return
	}

	mlog.Print("http://ip:3000")

	mlog.Print("gogs done!")
}
