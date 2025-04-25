# V1GroupVersionKind

GroupVersionKind unambiguously identifies a kind.  It doesn't anonymously include GroupVersion to avoid automatic coercion.  It doesn't use a GroupVersion to avoid custom marshalling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**group** | **str** |  | [default to '']
**kind** | **str** |  | [default to '']
**version** | **str** |  | [default to '']

## Example

```python
from tatope.models.v1_group_version_kind import V1GroupVersionKind

# TODO update the JSON string below
json = "{}"
# create an instance of V1GroupVersionKind from a JSON string
v1_group_version_kind_instance = V1GroupVersionKind.from_json(json)
# print the JSON string representation of the object
print(V1GroupVersionKind.to_json())

# convert the object into a dict
v1_group_version_kind_dict = v1_group_version_kind_instance.to_dict()
# create an instance of V1GroupVersionKind from a dict
v1_group_version_kind_from_dict = V1GroupVersionKind.from_dict(v1_group_version_kind_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


