from kubernetes import client, config
from kubernetes.client import V1ObjectMeta

from tatope.models import TatOpe, TatOpeSpec, TatOpeStatus

config.load_kube_config()

m = V1ObjectMeta(
    name="hogehoge",
    namespace="default"
)
print(m)

t = TatOpe(
    kind="TatOpe",
    api_version="tatope.tatope.local/v1",
    metadata=m,
    spec=TatOpeSpec(
        foo="fooval1",
        hoge="hogeval1",
        ports=[12,34]
    )
)

print(t)