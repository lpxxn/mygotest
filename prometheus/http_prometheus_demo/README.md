
https://songjiayang.gitbooks.io/prometheus/content/configuration/global.html
```
global 属于全局的默认配置，它主要包含 4 个属性，

scrape_interval: 拉取 targets 的默认时间间隔。
scrape_timeout: 拉取一个 target 的超时时间。
evaluation_interval: 执行 rules 的时间间隔。
external_labels: 额外的属性，会添加到拉取的数据并存到数据库中。

```

host.docker.internal

```
docker run \
    --name prometheus \
    -p 9090:9090 \
    -v $(pwd)/prometheus.yml:/etc/prometheus/prometheus.yml \
    prom/prometheus
```
http://localhost:9090/targets
http://localhost:9090/config

```
docker run -d --name=grafana --link prometheus -p 3000:3000 grafana/grafana
```
admin
admin
grafana setting-> source -> prometheus: 
http://prometheus:9090 
import 10826