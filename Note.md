## documentation little bit

```
1. gopath - that sets the env variable so that we can access the go codeanywhere in the system
    The Go path is used to resolve import statements.It is implemented by and documented in the go/build package.

     go env GOPATH - run this to check actual path 
```
## types 

```
1. case sensitive almost
2. variable type should be known in advance
3. everything is a type
4. String, Bool, Integer, Floating, Complex, Arrays, Maps, Slices,Structs, Pointers,     Function
```
## Variables & Constraints

```
1. create variable we use "var" then variable name + data type 
    eg = var username string = "adyasha"
2. %T = placeholder 
3. var value uint8 = 255 -> thsi is fine 
    BUT 
   value uint8 = 256 (this will give an err) -> cannot use 256 (untyped int constant) as uint8 value in variable declaration (overflows)
```

## Default values & Alias

```
1. if u create a variable of type int and dont assign a value to it so the default value which is get associated is 0
2. implicit way to assign a value
	var name = "adyasha"
	fmt.Println(name)
	fmt.Printf("the name is of type : %T \n", name) // string

3. declaration without var - just use this symbol ":=" || This is known as "Valorous Operator"
    Note - cant be declared without a method
	myName := "haadsaa"
	fmt.Println(myName);
    it will work fine , even if u change the value, the type will auto assign

4. const LoginToken string = "fdsksjk"    = HERE capital L has certain significance like its a public accesiible token, 
```
