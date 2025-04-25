# V1GroupVersion

GroupVersion contains the \"group\" and the \"version\", which uniquely identifies the API.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**group** | **str** |  | [default to '']
**version** | **str** |  | [default to '']

## Example

```python
from tatope.models.v1_group_version import V1GroupVersion

# TODO update the JSON string below
json = "{}"
# create an instance of V1GroupVersion from a JSON string
v1_group_version_instance = V1GroupVersion.from_json(json)
# print the JSON string representation of the object
print(V1GroupVersion.to_json())

# convert the object into a dict
v1_group_version_dict = v1_group_version_instance.to_dict()
# create an instance of V1GroupVersion from a dict
v1_group_version_from_dict = V1GroupVersion.from_dict(v1_group_version_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


