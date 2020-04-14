### api server监听地址

    0.0.0.0：8000
 
 ### url
 
 getuserinfo
 
    方法 GET
    http://127.0.0.1:8000/api/GETUsrInfo?name=asdffasd
    用户存在返回json
    {"name":"asdffasd","age":25,"address":"sadfasdfasdf"}   
    用不存在返回
    user is not exsit.
    
 setuserinfo（三种方式）
 
    1. url 传值  方法GET
        url：http://127.0.0.1:8000/api/SETUserInfo_urlargs?name=asdffasd&age=25&address=sadfasdfasdf
    
    2. body 传值  方法POST
        值的结构是form-data结构
        url：http://127.0.0.1:8000/api/SETUserInfo_bodykv
        request：
            user set is ok
            
    3. body 传值  方法POST
        值的结构是json  raw格式
        url：http://127.0.0.1:8000/api/SETUserInfo_bodyjson
        body内容：
            {"name":"asdffasd","age":25,"address":"sadfasdfasdf"}
         request：
            user set is ok
 ### POST form-data 说明
     
     Content-Type: multipart/form-data; boundary=--------------------------380384369246806537812958
     Content-Length: 387
     
     ----------------------------380384369246806537812958
     Content-Disposition: form-data; name="name"
     
     dsfadsfa
     ----------------------------380384369246806537812958
     Content-Disposition: form-data; name="age"
     
     25
     ----------------------------380384369246806537812958
     Content-Disposition: form-data; name="address"
     
     sadfasdfas
     ----------------------------380384369246806537812958--
