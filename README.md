# Wisdom-Httpd - 至理名言

wisdom-httpd是一个提供人生至理名言、冥想开悟的小工具，激励我们在人生苦短的时代，无论遇到什么挫折，都应该积极、努力的追求人生的价值和意义！

Website: https://wisdom.archstat.com

## How To

### 服务编译

```shell
cd /path
git clone https://github.com/lupguo/wisdom-httpd.git
go build
./wisdom-httpd -c ./config.yaml
```

### 服务启动

root用户配置`wisdom-httpd` systemctl配置

```shell
vim /etc/systemd/system/wisdom-httpd.service
...
# 开机启动服务
systemctl enable wisdom-httpd.service
# 服务启动
systemctl start wisdom-httpd.service
```

`/etc/systemd/system/wisdom-httpd.service` 内容如下:

```shell
[Unit]
Description=wisdom-httpd
After=network.target

[Service]
ExecStart=/data/go/bin/wisdom-httpd -c /data/projects/github.com/lupguo/wisdom-httpd/config.yaml
ExecStop=/bin/kill -SIGINT $MAINPID
ExecReload=/bin/kill -SIGHUP $MAINPID
Restart=always

[Install]
WantedBy=multi-user.target
```

### 测试

```shell
# html 格式，可以用于Notion内嵌
curl localhost:1666

# json 格式
curl localhost:1666/wisdom?type=json
{"sentence":"We dream and we build. We never give up, we never quit. 我们梦想，我们努力，决不放弃，决不退缩"}✔ /data/proje 
```

### Web配置

#### Caddy

`wisdom.archstat.com.caddyfile`配置文件：

```sh
$ cat /etc/caddy/Caddyfile.d/wisdom.archstat.com.caddyfile
wisdom.archstat.com {
    # 静态资源配置
    root * /data/projects/github.com/lupguo/wisdom-httpd/dist/prod
    file_server {
        index index.html
    }

    # API 代理配置
    reverse_proxy /api/* 127.0.0.1:1666 {
        header_up Host {host}
        header_up X-Real-IP {remote}
    }

}
```

**Systemd服务配置**: 

```
[Unit]
Description=Caddy web server
Documentation=https://caddyserver.com/docs/
After=network.target

[Service]
Type=notify
User=caddy
Group=caddy
ExecStartPre=/usr/bin/caddy validate --config /etc/caddy/Caddyfile
ExecStart=/usr/bin/caddy run --environ --config /etc/caddy/Caddyfile
ExecReload=/usr/bin/caddy reload --config /etc/caddy/Caddyfile
TimeoutStopSec=5s
LimitNOFILE=1048576
LimitNPROC=512
PrivateTmp=true
ProtectHome=true
ProtectSystem=full
AmbientCapabilities=CAP_NET_BIND_SERVICE

[Install]
WantedBy=multi-user.target
```

#### Nginx(可选）

```
server {
    listen 80;
    server_name you-domain.com;
    return 301 https://$host$request_uri;
}

server {
    listen 443 ssl;
    server_name you-domain.com;

    ssl_certificate /path/to/ssl_certificate.crt;
    ssl_certificate_key /path/to/ssl_certificate.key;

    location / {
        proxy_pass http://localhost:1666;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

## 灵感

人生苦短，每个人都应该积极探寻生命的价值和意义所在。

经验的探寻，在信息过载、信息茧房、信息碎片化的当代，想系统的获取高信息密度、有价值的信息，最快速、最高效的就是读优秀的好书。

因为一本好书是凝聚了作者的心血而成。从书籍中阅读和了解到作者的思想，和作者产生思维空间上的共鸣，这是多么Nice的一件事情。

人生路漫漫，总会充满了迷茫、彷徨、看不清，如果此时有一个高人能够为你拨云见日，那该是多么幸运的事情。

现实世界通常是残酷的，每个人都忙于生计，如同骡子一般每天辛勤劳作，获取那些微的报酬，给予生活的基本经济保障来源，还不时担心失业、教育、医疗、养老各种负担，人生真难。

我们又无法像曹操那么豁达“何以解忧，唯有杜康”，毕竟通过酒精麻痹后，没有任何改变，第二天还是要继续生活，或许这就是现实！

我们想做改变，但更多时候是苦于看不清方向，然后每天兜兜转转、忙忙碌碌、稀里糊涂的就过完了这一生。

当回首往事的时候，都无法为自己过去的任何事情而感到骄傲和欣慰，哪怕是一些许的小成功，因为我们之前没有规划过我们的人生，时间就这样在我们生命中无尽的流逝。

## 自己

想一想自己今年2024年，都已经36岁了，常说三十而立、四十不惑，自己从毕业就一直工作，到现在都已经有近14年了；

想想除了自己的技术有一些成长、思维变得成熟些，在事业上好像也没有做出什么特别的事情，难免陷入经常反思自己是不是做得不够多、不够好才导致了现在的状况，时常会感慨“命运不济、命途多舛”。

最近几年，在腾讯期间看到了很多优秀的人，一些人算是白领中的精英了，我从他们身上看到的很多高效能人士的优秀特质，比如积极主动、责任心、要事第一、以终为始、统合综效等等。

关于《高效能人士7个习惯》，我的那些同事（比我小4、5岁）可能很早就听过这本书了，因为在那之前他们已经通过读书或读书分享会，亦或是父母朋友，已经了解到书中的内容了，并积极贯彻到工作中。

所以回到读书这件事，不管是埃隆马斯克、吴军、罗翔、鲁迅、曾国藩、王阳明、王勃、李白、荀子、孔子，这些人信息掌握程度、思维高度、视野格局都已经远超过大部分人了。

然而这些人也都在强调读好书的重要性，所以寻找好的书，寻求经典，即便你无人可交流，好在有书本相伴！！

## 知行合一

在知道和做到之间还有一个鸿沟，即王阳明说到的“知行合一”，清华校训中也有“言大于行”一说。

尽信书不如无书，我们不应该为了读书而读书，应该是采用“我思故我在”方式去读取能解惑的书，同时主动的寻求经典书籍阅读。

## 至理名言

为何为会有至理名言，到了一定年龄，经历过一些事情才会有顿悟感觉，才会和当事人有共情。

比如之前听到很多的"时间管理"一词，会浅显认为只要做好计划，按部就班执行就可以达成目标，须不知人的精力是有限的，更重要的是对自己的精力做管理，而非时间管理。
人的一生也是有限的，有所为有所不为，这样才能更好的过好这一生！

我们在日常工作生活中，通常会因为忙于工作，时间悄然而逝去，时常会感慨年龄大了，一年年越来越快的感觉！

所以，这就是我希望通过一个工具，能够在自己日常工作过程中，时常不断警醒自己，好好关注自己，好好过完这一生！

## 最后

最终一句话，知易行难，道阻且长，在人生这条漫漫路上，我们还要不断修行和精进，不断放空自己，寻找人生的意义。

回到现实，我们每个人都拥有追求美好生活的权利，不论现实世界是怎样，都应该活出自己改有的样子！

保持习惯，多读好书，祝好！
