simple Rest API example with native http package

### Tried two examples using DefaultServeMux and NewServeMux

1. DefaultServeMux
     - The `DefaultServeMux` is a global variable that is created when the Go runtime starts. It is used by the `http.ListenAndServe()` function if you do not pass a custom ServeMux to it.
      - DefaultServeMux is shared by all go applications, any package that imports the net/http package can register handlers with the DefaultServeMux
      - as it allows third-party packages to register handlers that you may not be aware of,can be security risk. 
2. NewServeMux
    - The `NewServeMux()` function creates a new ServeMux object. You can use this to create your own custom routing for your application 
    - newly created ServeMux is not shared by other go apps/packages, so secured and you have complete control over routing.