# TatOpeSpec

TatOpeSpec defines the desired state of TatOpe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**foo** | **str** | Foo is an example field of TatOpe. Edit tatope_types.go to remove/update | [optional] 
**hoge** | **str** |  | [default to '']
**network** | [**Network**](Network.md) |  | [optional] 
**ports** | **List[int]** |  | [optional] 

## Example

```python
from tatope.models.tat_ope_spec import TatOpeSpec

# TODO update the JSON string below
json = "{}"
# create an instance of TatOpeSpec from a JSON string
tat_ope_spec_instance = TatOpeSpec.from_json(json)
# print the JSON string representation of the object
print(TatOpeSpec.to_json())

# convert the object into a dict
tat_ope_spec_dict = tat_ope_spec_instance.to_dict()
# create an instance of TatOpeSpec from a dict
tat_ope_spec_from_dict = TatOpeSpec.from_dict(tat_ope_spec_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


