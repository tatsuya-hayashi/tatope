# K8sIoApimachineryPkgVersionInfo

Info contains versioning information. how we'll want to distribute that information.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**build_date** | **str** |  | [default to '']
**compiler** | **str** |  | [default to '']
**git_commit** | **str** |  | [default to '']
**git_tree_state** | **str** |  | [default to '']
**git_version** | **str** |  | [default to '']
**go_version** | **str** |  | [default to '']
**major** | **str** |  | [default to '']
**minor** | **str** |  | [default to '']
**platform** | **str** |  | [default to '']

## Example

```python
from tatope.models.k8s_io_apimachinery_pkg_version_info import K8sIoApimachineryPkgVersionInfo

# TODO update the JSON string below
json = "{}"
# create an instance of K8sIoApimachineryPkgVersionInfo from a JSON string
k8s_io_apimachinery_pkg_version_info_instance = K8sIoApimachineryPkgVersionInfo.from_json(json)
# print the JSON string representation of the object
print(K8sIoApimachineryPkgVersionInfo.to_json())

# convert the object into a dict
k8s_io_apimachinery_pkg_version_info_dict = k8s_io_apimachinery_pkg_version_info_instance.to_dict()
# create an instance of K8sIoApimachineryPkgVersionInfo from a dict
k8s_io_apimachinery_pkg_version_info_from_dict = K8sIoApimachineryPkgVersionInfo.from_dict(k8s_io_apimachinery_pkg_version_info_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


