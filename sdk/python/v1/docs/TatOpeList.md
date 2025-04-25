# TatOpeList

TatOpeList contains a list of TatOpe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**api_version** | **str** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**items** | [**List[TatOpe]**](TatOpe.md) |  | 
**kind** | **str** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**metadata** | [**V1ListMeta**](V1ListMeta.md) |  | [optional] 

## Example

```python
from tatope.models.tat_ope_list import TatOpeList

# TODO update the JSON string below
json = "{}"
# create an instance of TatOpeList from a JSON string
tat_ope_list_instance = TatOpeList.from_json(json)
# print the JSON string representation of the object
print(TatOpeList.to_json())

# convert the object into a dict
tat_ope_list_dict = tat_ope_list_instance.to_dict()
# create an instance of TatOpeList from a dict
tat_ope_list_from_dict = TatOpeList.from_dict(tat_ope_list_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


