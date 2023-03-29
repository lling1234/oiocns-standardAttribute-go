# oiocns-standardAttribute-go
奥集能3.0平台分类标准特性和字典导入工具。

## 上架到应用商店思路
1. 获取用户token：输入账号密码或者通过SDK;
2. 获取speciesId和speciesCode；
3. 解析json文件或excel中的sheet页，sheet页中的内容（拓展：excel文件转json文件）；
4. 通过http.post请求接口，标准的分类特性数据写入平台；
5. 校验数据写入正确性。

## 特别注意
发送http POST请求的json格式只用arguments数组里面的对象就可以，不用带arguments、invocationId、target、type，否则会返回错误信息400和请求方法没找到。
```json
{
            "module": "thing",
            "action": "QuerySpeciesTree",
            "params": {
                "id": "380663455457349633",
                "page": {
                    "offset": 0,
                    "limit": 0,
                    "filter": ""
                }
            }
        }
```

## 查询实现
### 1.查询标准树

```json
{
    "arguments": [
        {
            "module": "thing",
            "action": "QuerySpeciesTree",
            "params": {
                "id": "380663455457349633",
                "page": {
                    "offset": 0,
                    "limit": 0,
                    "filter": ""
                }
            }
        }
    ],
    "invocationId": "289",
    "target": "Request",
    "type": 1
}
```

### 2.查询标准分类特性

```json
{
    "arguments": [
        {
            "module": "thing",
            "action": "QuerySpeciesDict",
            "params": {
                "id": "423526170806587392",
                "spaceId": "380663455457349633",
                "recursionOrg": true,
                "recursionSpecies": true,
                "page": {
                    "offset": 0,
                    "limit": 10000,
                    "filter": ""
                }
            }
        }
    ],
    "invocationId": "290",
    "target": "Request",
    "type": 1
}
```
### 3.查询标准字典定义

```json
{
    "arguments": [
        {
            "module": "thing",
            "action": "QuerySpeciesAttrs",
            "params": {
                "id": "423526170806587392",
                "spaceId": "380663455457349633",
                "recursionOrg": true,
                "recursionSpecies": true,
                "page": {
                    "offset": 0,
                    "limit": 10,
                    "filter": ""
                }
            }
        }
    ],
    "invocationId": "291",
    "target": "Request",
    "type": 1
}
```

### 4.查询标准字典定义项

```json
{
    "arguments": [
        {
            "module": "thing",
            "action": "QueryDictItems",
            "params": {
                "id": "424533157321248768",
                "spaceId": "380663455457349633",
                "page": {
                    "offset": 0,
                    "limit": 10,
                    "filter": ""
                }
            }
        }
    ],
    "invocationId": "296",
    "target": "Request",
    "type": 1
}
```



## 新增实现

接口：http://anyinone.com:800/orginone/kernel/rest/request
请求方式：post
请求头部信息：Authorization：xxx

### 1.新增标准分类

josn请求数据

```json
{
    "arguments": [
        {
            "module": "thing",
            "action": "CreateSpecies",
            "params": {
                "parentId": "366950230895235075",
                "name": "aaa",
                "code": "aaa",
                "belongId": "380663455457349633",
                "authId": "361356410044420096",
                "public": true,
                "remark": "aaa"
            }
        }
    ],
    "invocationId": "285",
    "target": "Request",
    "type": 1
}
```

### 2.新增分类特性


```json
{   
    "module": "thing",
    "action": "CreateAttribute",
    "params": {
        "speciesId": "423420264014024704",
        "speciesCode": "TransferManagement1",
        "public": true,
        "valueType": "描述型",
        "name": "444111",
        "code": "444111",
        "belongId": "362253609301315584",
        "authId": "361356410044420096",
        "remark": "444111 2023年3月14日13:48:37"
    }
}
```

json响应数据

```json
{
    "type": 3,
    "invocationId": "35",
    "result": {
        "code": 200,
        "data": {
            "id": "423478734755074048",
            "name": "111",
            "code": "111",
            "valueType": "描述型",
            "public": true,
            "remark": "111 2023年3月14日13:53:15",
            "speciesId": "423420393341194240",
            "belongId": "362253609301315584",
            "authId": "361356410044420096",
            "status": 1,
            "createUser": "358626578617470976",
            "updateUser": "358626578617470976",
            "version": "1",
            "createTime": "2023-03-14 13:53:18.220",
            "updateTime": "2023-03-14 13:53:18.220"
        },
        "msg": "success",
        "success": true
    }
}
```

### 3.新增字典分类

```json
{
    "module": "thing",
    "action": "CreateDict",
    "params": {
        "name": "55522",
        "code": "55522",
        "belongId": "380663455457349633",
        "public": true,
        "remark": "55522",
        "speciesId": "425658752432214016"
    }
}
```

### 4.新增字典项

```json
{
    "arguments": [
        {
            "module": "thing",
            "action": "CreateDictItem",
            "params": {
                "name": "aaa",
                "value": "aaa",
                "belongId": "380663455457349633",
                "remark": "aaa",
                "dictId": "424533157321248768"
            }
        }
    ],
    "invocationId": "38",
    "target": "Request",
    "type": 1
}
```
