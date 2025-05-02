# V1GroupVersionResource

GroupVersionResource unambiguously identifies a resource.  It doesn't anonymously include GroupVersion to avoid automatic coercion.  It doesn't use a GroupVersion to avoid custom marshalling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**group** | **str** |  | [default to '']
**resource** | **str** |  | [default to '']
**version** | **str** |  | [default to '']

## Example

```python
from tatope.models.v1_group_version_resource import V1GroupVersionResource

# TODO update the JSON string below
json = "{}"
# create an instance of V1GroupVersionResource from a JSON string
v1_group_version_resource_instance = V1GroupVersionResource.from_json(json)
# print the JSON string representation of the object
print(V1GroupVersionResource.to_json())

# convert the object into a dict
v1_group_version_resource_dict = v1_group_version_resource_instance.to_dict()
# create an instance of V1GroupVersionResource from a dict
v1_group_version_resource_from_dict = V1GroupVersionResource.from_dict(v1_group_version_resource_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


