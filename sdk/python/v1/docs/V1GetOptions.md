# V1GetOptions

GetOptions is the standard query options to the standard REST get call.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**api_version** | **str** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**kind** | **str** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**resource_version** | **str** | resourceVersion sets a constraint on what resource versions a request may be served from. See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.  Defaults to unset | [optional] 

## Example

```python
from tatope.models.v1_get_options import V1GetOptions

# TODO update the JSON string below
json = "{}"
# create an instance of V1GetOptions from a JSON string
v1_get_options_instance = V1GetOptions.from_json(json)
# print the JSON string representation of the object
print(V1GetOptions.to_json())

# convert the object into a dict
v1_get_options_dict = v1_get_options_instance.to_dict()
# create an instance of V1GetOptions from a dict
v1_get_options_from_dict = V1GetOptions.from_dict(v1_get_options_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


