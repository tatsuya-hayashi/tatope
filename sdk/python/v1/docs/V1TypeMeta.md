# V1TypeMeta

TypeMeta describes an individual object in an API response or request with strings representing the type of the object and its API schema version. Structures that are versioned or persisted should inline TypeMeta.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**api_version** | **str** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**kind** | **str** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 

## Example

```python
from tatope.models.v1_type_meta import V1TypeMeta

# TODO update the JSON string below
json = "{}"
# create an instance of V1TypeMeta from a JSON string
v1_type_meta_instance = V1TypeMeta.from_json(json)
# print the JSON string representation of the object
print(V1TypeMeta.to_json())

# convert the object into a dict
v1_type_meta_dict = v1_type_meta_instance.to_dict()
# create an instance of V1TypeMeta from a dict
v1_type_meta_from_dict = V1TypeMeta.from_dict(v1_type_meta_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


