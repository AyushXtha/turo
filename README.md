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
  ``` 
  python3 turo.py urls.txt
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
https://cm1.pw/a?k=FUZZ&d=FUZZ
https://cm1.pw?k=FUZZ&d=FUZZ&z=FUZZ
https://cm2.pw/4?k=FUZZ&d=FUZZ
https://cm2.pw/aa?b=FUZZ&aaa=FUZZ
https://cm2.pw?k=FUZZ&d=FUZZ&c=FUZZ&fd=FUZZ&a=FUZZ
```

   
   
                                 
                                     
                                     
                                     
                                     
                                     
                                     
                                     
                                     
                                     
                                     
                                     
                                     
                                     
                                     
