通过go实现webapi接口，数据的增删该查功能执行goWeb.go文件后
（go run goWeb.go)

使用postman模拟调用接口

>http:localhost:9090/user

增 : 
>post

body: 

```
	account:***
	
	password: ***
	
	name: ***
	
	sex: ***
	
	phone: ***
	
	level: ***
```

删 : 
>delete
		
body:

```
	account:***
	
	password:***
	
	name:***
	
	sex:***
	
	phone:***
	
	level:***
```

改 : 
>put

body: 

```
	account:***
	 
	password:***
```
查:
>get
