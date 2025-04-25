from kubernetes import client, config
from kubernetes.client import V1ObjectMeta

from tatope.models import TatOpe, TatOpeSpec, TatOpeStatus

t = TatOpe(
    kind="TatOpe"
    api_version="tatope.tatope.local/v1"
    metadata=V1ObjectMeta(
        name="hohoho"
        namespace="default"
    ),
    spec=TatOpeSpec(
        foo="fooval1",
        hoge="hogeval1",
        ports=[12,34]

    )
)