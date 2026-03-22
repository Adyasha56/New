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

## Memory management in GO 
```
2 things - make() , new()
1. New()                                    |  Make()
-------------------------------------------------------------------------------------
    (i) allocate memory but no init         | (i) allocate memory & Init
    (ii) u will get a memory address        | (ii) u will get a memory address
    (iii) zerod storage                     | (iii) no zerod storage

2. Garbage collection happens automatically 
3. out of scope or NIL
4.The GOGC variable sets the initial garbage collection target percentage. A collection is triggered when the ratio of freshly allocated data to 
    live data remaining after the previous collection reaches this percentage. The default is GOGC=100. Setting GOGC=off disables the garbage collector entirely. runtime/debug.SetGCPercent allows changing this percentage at run time.
```

## Pointers in golang
```
    1. why the pointer is needed?

    Ans- the prob in all prog lang is - whenever u declare a var, arr, u just a ref of memory. sometimes u pass on these var to a functions , the actual value is never passed on. sometimes a copy of these variables passed on into these multiple funcion. so arises the problem called - Irregularity. so to avoid these we use Pointer. So we pass the memoey address. to pass the actual value . 
```

## Array in Golang
```
1. array is not that much used in golang 
2. there is one another data type present in golang , which has more use cases than array
```

## Slices in Golang
```
very important in golang
1. Dynamic Resizing: Unlike arrays, which have a fixed size that is part of their type, slices can grow or shrink at runtime using the built-in append function.
2. Memory Efficiency: Slices are small, fixed-size structures (the "slice header") that point to an existing array. Copying or passing a slice to a function only copies this small header, not the entire data set.
3.Flexible Interface: Since a slice's length is not part of its type, you can write a single function that accepts slices of any size (e.g., []int), which is impossible with Go arrays (e.g., [5]int is a different type than [10]int).
```
## Maps In Go
```
1. store the data in the format of key val pair
2. key can be anything s as value 
3. not that musch string in case of Slices

languages := make(map[string]string)

	languages["js"] = "used more in web dev"
	languages["py"] = "used more in AI-ML"
	languages["c"] = "used more in embedded systems"

    //all val
    fmt.Println("the value inside the maps are: ", languages)

    //individual val
    fmt.Println("the value inside the maps are: ", languages["js])

```
## Structs in Golang
```
1. consider it as alternative of Class
2. No inheritance in golang
3. No super or parent
eg:  
//inside main func
adyashaNanda := User{"Adyasha", "adyasha12@gmail.com", true, 23}
fmt.Printf("the details of adyasha are: %+v\n", adyashaNanda)
fmt.Printf("Name is : %v and email is: %v  ", adyashaNanda.Name,adyashaNanda.Email)

//after main func
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

```

## switch Case 
```
Simple like other languages
* you need to write fallthrough- to compiling all other cases
```

## methods
```
in other object oriented language whenever a function goes inside a class we called it a method
same in golang - whenever we use a function inside structs we call it a method. 

- so yeah , there is not much about it
```

