# V1RootPaths

RootPaths lists the paths available at root. For example: \"/healthz\", \"/apis\".

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**paths** | **List[str]** | paths are the paths available at root. | 

## Example

```python
from tatope.models.v1_root_paths import V1RootPaths

# TODO update the JSON string below
json = "{}"
# create an instance of V1RootPaths from a JSON string
v1_root_paths_instance = V1RootPaths.from_json(json)
# print the JSON string representation of the object
print(V1RootPaths.to_json())

# convert the object into a dict
v1_root_paths_dict = v1_root_paths_instance.to_dict()
# create an instance of V1RootPaths from a dict
v1_root_paths_from_dict = V1RootPaths.from_dict(v1_root_paths_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


