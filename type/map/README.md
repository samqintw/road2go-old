# Road to GoLang
**Map**

init

    var x map[string]int = make(map[string]int)
    
    // or 
    x := make(map[string]int)
    
    // or with vaule
    x := map[string]int {
        "third": 3,
        "forth": 4,
    }
    
    // with array value
    x := make(map[string][]int)
    map["first"] = append(map["first"], 5)
    
    // with interface{}
    x := make(map[string]interface{})
    

access

    x["first"] = 1
    x["second"] = 2
    
delete
    
    delete(x, "third")



