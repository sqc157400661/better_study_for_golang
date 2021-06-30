## redis压测

### 环境：

- redis_version:6.2.4
-  redis benchmark



#### （1）字节大小：10 byte

```
# Memory
used_memory:2081880
used_memory_human:1.99M
used_memory_rss:6631424
used_memory_rss_human:6.32M
```

```
redis-benchmark -n 1000000 -r 30  -d 10
```

结果：

```
Summary:
  throughput summary: 48508.37 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.595     0.128     0.495     1.223     1.671    21.087
```

```
info Memory
# Memory
used_memory:14164664
used_memory_human:13.51M
used_memory_rss:19148800
used_memory_rss_human:18.26M
```

每个key大约占用：4.16byte

#### （2）字节大小：20 byte

```
redis-benchmark -n 1000000 -r 30  -d 20
```

结果：

```
Summary:
  throughput summary: 46257.75 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.643     0.144     0.503     1.311     1.767    22.063
```

```
info Memory
# Memory
used_memory:36337288
used_memory_human:34.65M
used_memory_rss:41984000
used_memory_rss_human:40.04M
```

每个key大约占用：2.17byte

#### （3）字节大小：50 byte

```
redis-benchmark -n 1000000 -r 30  -d 50
```

结果：

```
Summary:
  throughput summary: 47598.64 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.613     0.152     0.503     1.255     1.719    25.183
```

```
info Memory
# Memory
used_memory:88724888
used_memory_human:84.61M
used_memory_rss:95887360
used_memory_rss_human:91.45M
```



#### （4）字节大小：100 byte

```
redis-benchmark -n 1000000 -r 30  -d 100
```

结果：

```
1
```



#### （5）字节大小：1K

```
redis-benchmark -n 1000000 -r 30  -d 1024
```

结果：

```
1
```



#### （6）字节大小：5K

```
redis-benchmark -n 1000000 -r 30  -d 5120
```

结果：

```
1
```



10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能

