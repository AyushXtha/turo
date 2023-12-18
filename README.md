```
  _______   _    _   _____     ____  
 |__   __| | |  | | |  __ \   / __ \ 
    | |    | |  | | | |__) | | |  | |
    | |    | |  | | |  _  /  | |  | |
    | |    | |__| | | | \ \  | |__| |
    |_|     \____/  |_|  \_\  \____/                           
   ```               
  ## Merge parameters together 
   
   ## Usage

   GO Version
  ```
    go build kuro.go
    ./kuro.go urls.txt
  ```

    Python Version
  ``` 
  python3 turopreserve.py urls.txt
  ```
  
### Before
```
https://cm1.pw/a?d=5&k=4
https://cm1.pw?d=5&k=4&z=4
https://cm1.pw?d=5&k=4
https://cm2.pw/4?d=5&k=5
https://cm2.pw/aa?aaa=454&b=5
https://cm2.pw?a=5&k=4&fd=9
https://cm2.pw?d=4&c=5
https://cmc.pw/d
```

### After 
```
https://cm1.pw/a?d=5&k=4
https://cm1.pw?d=5&k=4&z=4
https://cm2.pw/4?d=5&k=5
https://cm2.pw/aa?aaa=454&b=5
https://cm2.pw?a=5&k=4&fd=9&d=4&c=5

```

  ``` 
  python3 turopreserve.py urls.txt fuzzzz
  ```
```
https://cm1.pw/a?d=fuzzzz&k=fuzzzz
https://cm1.pw?d=fuzzzz&k=fuzzzz&z=fuzzzz
https://cm2.pw/4?d=fuzzzz&k=fuzzzz
https://cm2.pw/aa?aaa=fuzzzz&b=fuzzzz
https://cm2.pw?a=fuzzzz&k=fuzzzz&fd=fuzzzz&d=fuzzzz&c=fuzzzz



```
                                 
                                     
                                     
                                     
                                     
                                     
                                     
                                     
                                     
                                     
                                     
                                     
                                     
                                     
                                     
