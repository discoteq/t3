# t3: the templating tool

`t3` is a tool to populate a template from a set of predefined data.  Today, often the best tools available for this job is `sed(1)` or `awk(1)`. While extremely powerful, these tools are terrible at templating.

`t3` is designed for a world with `docker`, `etcd`, `consul` & `vault`. Or really, a world with `curl`, `jq` and URLs like `http://169.254.169.254/latest/meta-data/`. A world with `cloud-init` scripts on auto-scaled instances. A world to impatient to type all the letters in `kubernetes`. `t3` is the hero we deserve.

## The simplest thing that could possibly work

```
cat >services.json <<EOF
{
  "statsd": [
    {
      "hostname": "10.0.0.216",
      "port": 8126
    }
  ]
}
EOF

cat >config.mustache <<EOF
{#statsd}
statsd.host={{hostname}}:{{port}}
{/statsd}
EOF

t3 -d services.json -t config.mustache > services.cfg

cat services.cfg
statsd.host=10.0.0.216:8126
```


## That was delightfully simple, but what could this possibly be good for?

Imagine you've got a service called `myface` and you want to reload its config any time the values in your service config JSON file are updated.

```
#!/bin/sh
# myface-cfg-update - watch for service updates
#   and safely load them into myface

BASE=myface
TEMP=`mktemp -t $BASE.XXXXXXXXXX` || exit 1

SERVICE_MAP=/var/lib/discoteq/services.json
TEMPLATE=/opt/myface/config.mustache
CONFIG=/etc/myface.cfg
DAEMON=myface

# watch for service info changes in registry
# re-eval config template
# verify generated config is valid
# swap if valid and different
# reload if swapped
fswatch -1 $SERVICE_MAP &&
  t3 -d $SERVICE_MAP -t $TEMPLATE > $TEMP &&
  myface --check-cfg $TEMP &&
  flock $CONFIG diffswp $TEMP $CONFIG &&
  service $DAEMON reload
```
