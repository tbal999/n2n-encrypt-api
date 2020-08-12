# n2n-encrypt-api
Encodes and decodes strings!
This is best used on small strings - like two or three words - or usernames/passwords.
For example, on a 30kb string the key will be 166kb and the encrypted message will be over 1GB in size!

How?
- It shifts every single character by a random amount and creates a key that's the length of the string. 
- Then it adds duplication to make it harder to read.
- The key generatad will be as long as the string and is random every time.

Commands:
1) /encode/{string} - encodes a string and generates a key that's the length of the string, and the encrypted output in bytes - example: {"String":"TEST","Key":[1,2,3,4],"Result":"encryptedoutputinbytes"}

2) /decode/JSON - decodes JSON formatted like above

available here https://n2n-encrypt.herokuapp.com
