https://golang.org/ref/spec#defer_statements

A "defer" statement invokes a function whose execution is deferred to the moment the surrounding
 function **returns**, either because the surrounding function executed a return statement, reached the end of its function body, or because the corresponding goroutine is **panicking**.

Each time a "defer" statement executes, the function value and parameters to the call are evaluated as usual and 
saved anew but the actual function is not invoked. 
Instead, **deferred functions are invoked immediately before the surrounding function returns**, 
in the reverse order they were deferred. If a deferred function value evaluates to nil, 
execution panics when the function is invoked, not when the "defer" statement is executed.

For instance, if the deferred function is a **"function literal"** and the surrounding function has named **"result 
parameters"** that are in scope within the literal, **the deferred function may access and modify the "result parameters"** 
before they are returned. If the deferred function has any return values, they are **discarded** when the function completes. (See also the section on handling panics.)

    lock(l)
    defer unlock(l)  // unlocking happens before surrounding function returns
    
    // prints 3 2 1 0 before surrounding function returns
    for i := 0; i <= 3; i++ {
    	defer fmt.Print(i)
    }
    
    // f returns 1
    func f() (result int) {
    	defer func() {
    		result++
    	}()
    	return 0
    }
    
    
https://tiancaiamao.gitbooks.io/go-internals/content/zh/03.4.html