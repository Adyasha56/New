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
## User Input 
```
1. read documentation about bufio, Os 
2. we use reader := bufio..NewReader(os.Stdin) = this line helpful for taking inputs, 
3.after this comes comma ok || err ok syntax
4.input, _ := reader.ReadString('\n') 
5. and then u print the input

    thers a lot to know, as of now just got the idea how to write the syntax and how to get te user input 

```

## Connversions in Golang
```
1. lets say - u take an input , but go by default made this to string 
2. now ifu want to add 1 to a the input so it will show u the err "can not convert"
3. Solution - we will use a package known as "strconv" 
4.numRating,err = strconv.ParseFloat(input,64) - in this we used strconv package,it has multiple function so we are using ParseFloat as of now, which takes 2 things as its parameter  
            (i) - input ,and size which is 64 bits 
            (ii) - as we pass 2 values so we need to obey comma ok syntax by adding "err" after the variable,
5. Post to this we need to handle the err cautiosly 
6. numRating,err := strconv.ParseFloat(input,64)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("added 1 to your rating:: ", numRating+1)
	} 

    // in this code u will face an err like this 
    err msz = he type of input is stringstrconv.ParseFloat: parsing "2\r\n": invalid syntax
7. Solution - 
        instead of this => `numRating,err := strconv.ParseFloat(input,64)`
        we should use another package named "strings" 
        `numRating,err := strconv.ParseFloat(strings.TrimSpace(input),64)`

    // what we have done ? 
    // we used strings package and inside that TrimSpace function to avoid that \n trailing , cause its the only thing which caused the err during conversion
```

## Handling Time
```
1. for getting the current date - we do use time package
2. presentTime := time.Now()
	fmt.Println(presentTime) 
    o/p = 2026-03-09 12:09:17.2904399 +0530 IST m=+0.000000001
3. but the output looks messy, so we use 'Format()' func.
4. FOR DATE - fmt.Println(presentTime.Format("01-02-2006")) - have to use this syntax
5. FOR TIME & Date - fmt.Println(presentTime.Format("01-02-2006  15:04:05")) - have to  use this time
6. FOR TIME DATE & DAY - fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))
```

## Build GO for diff OS
```
1. if i want to run a go code as well as it able to run in the mac & linux env as well 
2. for that we will do few things
        (i) - in the terminal run the command - go env
        (ii) -it will show some parameters , among all GOOS is what , we deal with 
        (iii) - now write the command $env:GOOS = "darwin"; go build if u are in powershell
                powershell core - $env:GOOS="darwin"; go build
                Windows sub system - GOOS="darwin" go build
                for multiple platforms - $env:GOOS = "darwin"; $env:GOARCH = "amd64"; go build

```