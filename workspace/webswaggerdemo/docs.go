package main

//This file is generated automatically. Do not try to edit it manually.

var resourceListingJson = `{
    "apiVersion": "0.0.0",
    "swaggerVersion": "1.2",
    "apis": [
        {
            "path": "/a",
            "description": "Test API"
        }
    ],
    "info": {
        "title": "title",
        "description": "description",
        "contact": "user@domain.com",
        "termsOfServiceUrl": "http://...",
        "license": "MIT",
        "licenseUrl": "http://osensource.org/licenses/MIT"
    }
}`
var apiDescriptionsJson = map[string]string{"a": `{
    "apiVersion": "0.0.0",
    "swaggerVersion": "1.2",
    "basePath": "{{.}}",
    "resourcePath": "/a",
    "produces": [
        "application/json"
    ],
    "apis": [
        {
            "path": "/a/b",
            "description": "get struct3",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "GetStruct3",
                    "type": "string",
                    "items": {},
                    "summary": "get struct3",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "string"
                        }
                    ],
                    "produces": [
                        "application/json"
                    ]
                }
            ]
        }
    ]
}`}
