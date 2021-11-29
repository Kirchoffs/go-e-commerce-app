# Notes

## 2022-1-17
In serve() function, a server object is returned.  
For a server object, **handler** field is required.  
```
srv := &http.Server{
    ...
    Handler:           app.routes(),
    ...
}
```

In app.routes() (in cmd/web/routes.go) function, a http.Handler object is returned.  
For a handler, we can specify the endpoint and the corresponding handling function.  
This handling funciton is just like controller in Java web, request and response are the parameters of it.  

## 2022-1-18
In development, we don't want to take the data from the cache, since we want to see the changes.  

Using air to do the development: https://github.com/cosmtrek/air  
.air.toml is the configuration file for air 
```
>> air
```

Instead of using *Go Template Support* extension, we use *gotemplate-syntax* extension for VS code.  
Originally, the template files under the templates folder is named with the suffix **tmpl**.  Now after changing the extension, we use the suffix **gohtml**.
