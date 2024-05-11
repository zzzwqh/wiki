```groovy
```groovy
if (ENV.equals("qa")) {
 
    if ("gate|game|friend|dbproxy".contains(GameService)) {
        return ["1","2"];  
    } else if ("gate|game|friend|dbproxy".contains(GameService)) {
        return ["3","4"];  
    }
} else if (ENV.equals("online")) {
 
    if ("gate|game|friend|dbproxy".contains(GameService)) {
        return ["5","6"];  
    } else if ("gate|game|friend|dbproxy".contains(GameService)) {
        return ["7","8"];  
    }
   
}
```
```