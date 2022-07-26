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
Notice that .air.toml is the configuration file for air. 
```
>> air
```

Instead of using *Go Template Support* extension, we use *gotemplate-syntax* extension for VS code.  
Originally, the template files under the templates folder is named with the suffix **tmpl**.  Now after changing the extension, we use the suffix **gohtml**.  


Where should we put the \<script\> tag in HTML?  
The antiquated approach: we can put it at the bottom of \<body\>  
The modern way: use async or defer
```
<script src="path/to/script1.js" async></script>
<script src="path/to/script2.js" async></script>

or

<script src="path/to/script1.js" defer></script>
<script src="path/to/script2.js" defer></script>
```

## 2022-01-20
The Access-Control-Allow-Credentials response header tells browsers whether to expose the response to the frontend JavaScript code when the request's credentials mode (Request.credentials) is include.  

When the Request.credentials is “include” mode browsers will expose the response to front-end JavaScript code if the Access-Control-Allow-Credentials is set true.  

Credentials are actually cookies, authorization headers or TLS (Transport Layer Security) client certificates.  

The credentials read-only property of the Request interface indicates whether the user agent should send cookies from the other domain in the case of cross-origin requests.  

```
// xhr
var xhr = new XMLHttpRequest();
xhr.open('GET', 'https://www.geeksforgeeks.org/', true); 
xhr.withCredentials = true; 
xhr.send(null);

// fetch
fetch(url, {
    credentials: 'include'  
})
```

## 2022-02-18
package encoding/json  

```
type point struct {
    X int `json:"x"`
    Y int `json:"y"`
}

p := &point {
    X: 1,
    Y: 2,
}

m, _ := json.MarshalIndent(p, "", "  ")

fmt.Println(string(m))
```
The output of the above code snippet:  
```
{  
  "x": 1,  
  "y": 2  
}  
```

## 2022-05-30
1. Create a simple project:  
```
>> mkdir prj && cd prj
>> go mod init prj
```

2. Use single code case to produce two different binary, one for the front end, one for the back end.  
Question:  
How to create multiple binary from the same repo?

## 2022-06-07
1. Cast an interface into a concrete type
```
concreteExample = interfaceExample.(T)
```
2. Project details  
cmd/api: back end  
cmd/web: front end

3. Package encoding/json
```
type stripePayload struct {
    Currency string `json:"currency"`
    Amount   string `json:"amount"`
}
```

## 2022-06-08
1. How to create a basic HTTP server
```
srv := &http.Server{
    Addr:              fmt.Sprintf(":%d", app.config.port),
    Handler:           app.routes(),
    IdleTimeout:       30 * time.Second,
    ReadTimeout:       10 * time.Second,
    ReadHeaderTimeout: 5 * time.Second,
    WriteTimeout:      5 * time.Second,
}

func (app *application) routes() http.Handler {
    mux := chi.NewRouter()
    mux.Use(cors.Handler(cors.Options{
        AllowedOrigins:   []string{"https://*", "http://*"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-TOKEN"},
        AllowCredentials: false,
        MaxAge:           300, // 300 seconds
    }))
    mux.Get("/path/to/endpoint", app.GetEndpoint)
    return mux
}

import (
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/cors"
)

func (app *application) GetEndpoint(w http.ResponseWriter, r *http.Request) {
    res := jsonResponse{
        OK: true,
    }

    out, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
        app.errorLog.Println(err)
    }

    w.Header().Set("Content-Type", "json/application")
    w.Write(out)
}
```
2. Makefile should be named with "Makefile" or "makefile"

## 2023-04-20
```
>> go mod tidy
```

The go mod tidy command is used to clean up and optimize the Go module dependencies in a project.

When you use external packages in a Go project, you typically specify them in your project's go.mod file, along with their version numbers. However, as you add and remove dependencies, your go.mod file can become cluttered with unused packages, and versions of packages that are no longer needed.

go mod tidy cleans up your go.mod file by removing any unused or unnecessary dependencies, and updating the version numbers of the packages that you are using to the latest compatible version. It does this by analyzing your code and looking for the packages that are actually being used in your project, and removing any that are not needed.