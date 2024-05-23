执行 asprof 生成火焰图需要宿主机配置如下
```bash
echo 1028 > /proc/sys/kernel/perf_event_mlock_kb
sysctl kernel.perf_event_paranoid=1
sysctl kernel.kptr_restrict=0
```
