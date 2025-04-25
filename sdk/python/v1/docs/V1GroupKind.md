# V1GroupKind

GroupKind specifies a Group and a Kind, but does not force a version.  This is useful for identifying concepts during lookup stages without having partially valid types

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**group** | **str** |  | [default to '']
**kind** | **str** |  | [default to '']

## Example

```python
from tatope.models.v1_group_kind import V1GroupKind

# TODO update the JSON string below
json = "{}"
# create an instance of V1GroupKind from a JSON string
v1_group_kind_instance = V1GroupKind.from_json(json)
# print the JSON string representation of the object
print(V1GroupKind.to_json())

# convert the object into a dict
v1_group_kind_dict = v1_group_kind_instance.to_dict()
# create an instance of V1GroupKind from a dict
v1_group_kind_from_dict = V1GroupKind.from_dict(v1_group_kind_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


