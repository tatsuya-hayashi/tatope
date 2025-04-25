from kubernetes import client, config

from tatope.models import TatOpe, TatOpeSpec, TatOpeStatus

from tatope.models.v1_object_meta import V1ObjectMeta

config.load_kube_config()

m = V1ObjectMeta(
    name="hogehoge",
    namespace="default"
)
#print(m)

t = TatOpe(
    kind="TatOpe",
    api_version="tatope.tatope.local/v1",
    metadata=m,
    spec=TatOpeSpec(
        hoge="hogehogefooval",
        ports=[12,34]
    )
).to_dict()

#print(t)

crd_api = client.CustomObjectsApi()

try:
    result = crd_api.create_namespaced_custom_object(
       group="tatope.tatope.local",
       version="v1",
       namespace="default",
       plural="tatopes",
       body=t
    )
except Exception as e:
    print(e)




crd_api.patch_namespaced_custom_object_status(
    group="tatope.tatope.local",
    version="v1",
    namespace="default",
    plural="tatopes",
    name="hogehoge",
    body={'status':{'state':'active'}}
)
crd_api.get_namespaced_custom_object(
    group="tatope.tatope.local",
    version="v1",
    namespace="default",
    plural="tatopes",
    name="hogehoge"
)