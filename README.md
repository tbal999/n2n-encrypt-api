# n2n-encrypt-api
Encodes and decodes strings!
This is best used on small strings - like one line sentences. 
For example, on a 30kb string the key will be 166kb and the encrypted message will be over 1GB in size!

Commands:
1) /encode/{string} - encodes a string and generates a key that's the length of the string, and the encrypted output in bytes - example: {"String":"TEST","Key":[1,2,3,4],"Result":"encryptedoutputinbytes"}

2) /decode/JSON - decodes JSON formatted like above

available here https://n2n-encrypt.herokuapp.com
