MysqlUser=$1
MysqlPassword=$2
HAProxyIp=$3

for id in `mysql -u $MysqlUser -p$MysqlPassword -e "show processlist;" | grep $HAProxyIp | awk '{print $1}'`;
do mysql -u $MysqlUser -p$MysqlPassword -e "kill $id";
done
