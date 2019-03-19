# Road to GoLang
**Interface**

Type assertions

    // string
    var i interface{} = "hello"  
    s, ok := i.(string)
    fmt.Println(s, ok)
    
    
    // function pointer
    func myfunc(str string) string {
        return fmt.printf("your str is %s", str)
    }
    var i interface() = myfunc
    if v, ok := i.(func(string) string); ok {
        v("mysmchin")
    } 
    
    // another interface

