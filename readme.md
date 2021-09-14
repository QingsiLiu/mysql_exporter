监控：
1.可用性
2.延迟
    请求消耗时间
    操作使用时间
3.错误次数
4.容量
    当前请求多少/总请求多少
    当前连接数量/总的连接数量

mysql => exporter =>
    监控对象api => 获取指标信息（计算）
    sql查询 => show global status

mysql可用性
    操作失败：
        select 1;
        ping
慢查询次数：
    show global status where variable_name = 'low_queries'
容量：
    qps:
        show global status where variable_name = 'Queries'
    tps:
        insert, delete, update
        com_insert
        com_update
        com_delete
        com_select
        com_replace
    连接:
        show global status where variable_name = 'Threads_running'
        show global variables where variable_name = 'max_connections'
    流量：
        show global status where variable_name = 'Bytes_received'
        show global status where variable_name = 'Bytes_send'

// mysql 连接信息 => mysql host, port, user
