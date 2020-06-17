package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

var jsonstr string = `[
    {
        "Node": {
            "ID": "6a561072-eab7-cb63-ab56-17932a62cef8",
            "Node": "sunreal112",
            "Address": "192.168.3.112",
            "Datacenter": "dc1",
            "TaggedAddresses": {
                "lan": "192.168.3.112",
                "lan_ipv4": "192.168.3.112",
                "wan": "192.168.3.112",
                "wan_ipv4": "192.168.3.112"
            },
            "Meta": {
                "consul-network-segment": ""
            },
            "CreateIndex": 163854,
            "ModifyIndex": 163854
        },
        "Service": {
            "ID": "anzhi-server-api-646d2977e399f84f3ff5c4480a7b509a",
            "Service": "anzhi-server-api",
            "Tags": [
                "weight=1",
                "secure=false",
                "contextPath=/anzhi-app-api"
            ],
            "Meta": null,
            "Port": 9002,
            "Address": "192.168.3.112",
            "TaggedAddresses": {
                "lan_ipv4": {
                    "Address": "192.168.3.112",
                    "Port": 9002
                },
                "wan_ipv4": {
                    "Address": "192.168.3.112",
                    "Port": 9002
                }
            },
            "Weights": {
                "Passing": 1,
                "Warning": 1
            },
            "EnableTagOverride": false,
            "CreateIndex": 172029,
            "ModifyIndex": 172029,
            "Proxy": {
                "MeshGateway": {},
                "Expose": {}
            },
            "Connect": {}
        },
        "Checks": [
            {
                "Node": "sunreal112",
                "CheckID": "serfHealth",
                "Name": "Serf Health Status",
                "Status": "passing",
                "Notes": "",
                "Output": "Agent alive and reachable",
                "ServiceID": "",
                "ServiceName": "",
                "ServiceTags": [],
                "Type": "",
                "Definition": {
                    "Interval": "0s",
                    "Timeout": "0s",
                    "DeregisterCriticalServiceAfter": "0s",
                    "HTTP": "",
                    "Header": null,
                    "Method": "",
                    "Body": "",
                    "TLSSkipVerify": false,
                    "TCP": ""
                },
                "CreateIndex": 163854,
                "ModifyIndex": 163854
            },
            {
                "Node": "sunreal112",
                "CheckID": "service:anzhi-server-api-646d2977e399f84f3ff5c4480a7b509a",
                "Name": "Service 'anzhi-server-api' check",
                "Status": "passing",
                "Notes": "",
                "Output": "",
                "ServiceID": "anzhi-server-api-646d2977e399f84f3ff5c4480a7b509a",
                "ServiceName": "anzhi-server-api",
                "ServiceTags": [
                    "weight=1",
                    "secure=false",
                    "contextPath=/anzhi-app-api"
                ],
                "Type": "ttl",
                "Definition": {
                    "Interval": "0s",
                    "Timeout": "0s",
                    "DeregisterCriticalServiceAfter": "0s",
                    "HTTP": "",
                    "Header": null,
                    "Method": "",
                    "Body": "",
                    "TLSSkipVerify": false,
                    "TCP": ""
                },
                "CreateIndex": 172029,
                "ModifyIndex": 172030
            }
        ]
    },
    {
        "Node": {
            "ID": "6a561072-eab7-cb63-ab56-17932a62cef8",
            "Node": "sunreal112",
            "Address": "192.168.3.112",
            "Datacenter": "dc1",
            "TaggedAddresses": {
                "lan": "192.168.3.112",
                "lan_ipv4": "192.168.3.112",
                "wan": "192.168.3.112",
                "wan_ipv4": "192.168.3.112"
            },
            "Meta": {
                "consul-network-segment": ""
            },
            "CreateIndex": 163854,
            "ModifyIndex": 163854
        },
        "Service": {
            "ID": "anzhi-server-api-e4edffdb1a35b151b1af33bb12266fbe",
            "Service": "anzhi-server-api",
            "Tags": [
                "weight=1",
                "secure=false",
                "contextPath=/anzhi-app-api"
            ],
            "Meta": null,
            "Port": 8080,
            "Address": "192.168.100.93",
            "TaggedAddresses": {
                "lan_ipv4": {
                    "Address": "192.168.100.93",
                    "Port": 8080
                },
                "wan_ipv4": {
                    "Address": "192.168.100.93",
                    "Port": 8080
                }
            },
            "Weights": {
                "Passing": 1,
                "Warning": 1
            },
            "EnableTagOverride": false,
            "CreateIndex": 172036,
            "ModifyIndex": 172036,
            "Proxy": {
                "MeshGateway": {},
                "Expose": {}
            },
            "Connect": {}
        },
        "Checks": [
            {
                "Node": "sunreal112",
                "CheckID": "serfHealth",
                "Name": "Serf Health Status",
                "Status": "passing",
                "Notes": "",
                "Output": "Agent alive and reachable",
                "ServiceID": "",
                "ServiceName": "",
                "ServiceTags": [],
                "Type": "",
                "Definition": {
                    "Interval": "0s",
                    "Timeout": "0s",
                    "DeregisterCriticalServiceAfter": "0s",
                    "HTTP": "",
                    "Header": null,
                    "Method": "",
                    "Body": "",
                    "TLSSkipVerify": false,
                    "TCP": ""
                },
                "CreateIndex": 163854,
                "ModifyIndex": 163854
            },
            {
                "Node": "sunreal112",
                "CheckID": "service:anzhi-server-api-e4edffdb1a35b151b1af33bb12266fbe",
                "Name": "Service 'anzhi-server-api' check",
                "Status": "passing",
                "Notes": "",
                "Output": "",
                "ServiceID": "anzhi-server-api-e4edffdb1a35b151b1af33bb12266fbe",
                "ServiceName": "anzhi-server-api",
                "ServiceTags": [
                    "weight=1",
                    "secure=false",
                    "contextPath=/anzhi-app-api"
                ],
                "Type": "ttl",
                "Definition": {
                    "Interval": "0s",
                    "Timeout": "0s",
                    "DeregisterCriticalServiceAfter": "0s",
                    "HTTP": "",
                    "Header": null,
                    "Method": "",
                    "Body": "",
                    "TLSSkipVerify": false,
                    "TCP": ""
                },
                "CreateIndex": 172036,
                "ModifyIndex": 172037
            }
        ]
    }
]`

func main()  {
	n := gjson.Get(jsonstr,`#`)
	fmt.Println(n.Num)

	//for i:= 0;i<=n ;i++ {
	//	r := gjson.Get(jsonstr,i+"Service.Port")
	//}
	//b := gjson.GetMany(jsonstr,`#.Service`)
	//fmt.Println(b)
	//
	//a := gjson.GetMany(jsonstr,`#.Service.Port`)
	//fmt.Println(a)
}
