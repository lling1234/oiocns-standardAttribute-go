# oiocns-standardAttribute-go
标准分类特性导入小工具

## 上架到应用商店思路
1. 获取用户token：输入账号密码或者通过SDK;
2. 获取speciesId和speciesCode；
3. 解析json文件或excel中的sheet页，sheet页中的内容（拓展：excel文件转json文件）；
4. 通过http.post请求接口，标准的分类特性数据写入平台；
5. 校验数据写入正确性。

## 实现
接口：http://anyinone.com:800/orginone/kernel/rest/request
请求方式：post
请求头部信息：Authorization：xxx

josn请求数据

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

```json
{
    "type": 3,
    "invocationId": "36",
    "result": {
        "code": 200,
        "data": {
            "limit": 20,
            "total": 6,
            "result": [
                {
                    "id": "416237430006484992",
                    "name": "唯一标识",
                    "code": "Id",
                    "valueType": "描述型",
                    "public": true,
                    "remark": "证明存在的唯一标识，由系统自动生成（雪花ID）。",
                    "speciesId": "366950230870069248",
                    "authId": "361356410044420096",
                    "status": 1,
                    "createUser": "358221262448889856",
                    "updateUser": "358221262448889856",
                    "version": "1",
                    "createTime": "2023-02-22 14:18:56.631",
                    "updateTime": "2023-03-11 13:42:41.000"
                },
                {
                    "id": "422398957721882624",
                    "name": "创建人",
                    "code": "Creater",
                    "valueType": "组织型",
                    "public": true,
                    "remark": "首次创建对象的人员。",
                    "speciesId": "366950230870069248",
                    "authId": "361356410044420096",
                    "status": 1,
                    "createUser": "358221262448889856",
                    "updateUser": "358221262448889856",
                    "version": "2",
                    "createTime": "2023-03-11 14:22:39.315",
                    "updateTime": "2023-03-11 14:23:27.151"
                },
                {
                    "id": "422399137753993216",
                    "name": "创建时间",
                    "code": "CreateTime",
                    "valueType": "时间型",
                    "public": true,
                    "remark": "首次创建的时间。",
                    "speciesId": "366950230870069248",
                    "authId": "361356410044420096",
                    "status": 1,
                    "createUser": "358221262448889856",
                    "updateUser": "358221262448889856",
                    "version": "1",
                    "createTime": "2023-03-11 14:23:22.238",
                    "updateTime": "2023-03-11 14:23:22.238"
                },
                {
                    "id": "422406362727845888",
                    "name": "状态",
                    "code": "Status",
                    "valueType": "选择型",
                    "public": true,
                    "dictId": "422405346380877824",
                    "remark": "生命状态。",
                    "speciesId": "366950230870069248",
                    "authId": "361356410044420096",
                    "status": 1,
                    "createUser": "358221262448889856",
                    "updateUser": "358221262448889856",
                    "version": "3",
                    "createTime": "2023-03-11 14:52:04.806",
                    "updateTime": "2023-03-11 14:52:20.041",
                    "dict": {
                        "id": "422405346380877824",
                        "name": "状态字典",
                        "code": "状态字典",
                        "remark": "标识对象的当前状态。",
                        "public": false,
                        "speciesId": "366950230870069248",
                        "belongId": "358223410553294848",
                        "status": 1,
                        "createUser": "358221262448889856",
                        "updateUser": "358221262448889856",
                        "version": "1",
                        "createTime": "2023-03-11 14:48:02.490",
                        "updateTime": "2023-03-11 14:48:02.490",
                        "dictItems": [
                            {
                                "id": "422405910112112640",
                                "name": "正常",
                                "value": "正常",
                                "public": false,
                                "belongId": "358223410553294848",
                                "dictId": "422405346380877824",
                                "status": 1,
                                "createUser": "358221262448889856",
                                "updateUser": "358221262448889856",
                                "version": "1",
                                "createTime": "2023-03-11 14:50:16.893",
                                "updateTime": "2023-03-11 14:50:16.893"
                            },
                            {
                                "id": "422406098562191360",
                                "name": "已销毁",
                                "value": "已销毁",
                                "public": false,
                                "belongId": "358223410553294848",
                                "dictId": "422405346380877824",
                                "status": 1,
                                "createUser": "358221262448889856",
                                "updateUser": "358221262448889856",
                                "version": "1",
                                "createTime": "2023-03-11 14:51:01.824",
                                "updateTime": "2023-03-11 14:51:01.824"
                            }
                        ]
                    }
                },
                {
                    "id": "422436687172472832",
                    "name": "修改时间",
                    "code": "ModifiedTime",
                    "valueType": "时间型",
                    "public": true,
                    "remark": "最后一次修改时间。",
                    "speciesId": "366950230870069248",
                    "authId": "361356410044420096",
                    "status": 1,
                    "createUser": "358221262448889856",
                    "updateUser": "358221262448889856",
                    "version": "1",
                    "createTime": "2023-03-11 16:52:34.716",
                    "updateTime": "2023-03-11 16:52:34.716"
                },
                {
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
                    "createTime": "2023-03-14 13:53:18.221",
                    "updateTime": "2023-03-14 13:53:18.221"
                }
            ]
        },
        "msg": "success",
        "success": true
    }
}
```