## Random facts library

This library written in Golang and provides API to get random fact by set of keywords. Facts storage is configurable, but library contains some sample facts

### Dependencies

Project does not contains any dependencies

### Usage

Install library via command

    go get "github.com/sheremetat/randfacts-lib" 
    
Copy files from `facts/` folder to `your/folder/*`. Also, you may create your own files with facts (see file structure bellow). Then import to your project
```go
package main

import (
   facts "github.com/sheremetat/randfacts-lib"
}
```
Init the library

```go
func main(){
    factLib, err := facts.Init("your/folder/")
    if err != nil {
        panic("Cannot init facts library")
    }
}
```

And use it to retrieve random fact as string value and error:

```go
...
word := "keyword"
fact, err := factLib.GetFact(word)
if err == nil {
    // do something with fact
}
```
Method `GetFact(...)` will analyse text and retrieve random fact by if parameter is a one of keys. If 
Or you can ask library to return first random fact if text contains any of available keywords

```go
...
text := "long text with keyword"
fact, err := factLib.FindFact(text)
if err == nil {
    // do something with fact
}
``` 
### Facts file example

File with facts is a simple text file. Each line from file contains th fact.
First line contains keywords separated by comma. Search by keywords is case insensitive. File without keywords will be skipped and ignored. Also, file with sixe more than 100000 bytes will be skipped.

```
keywords:key1,key2,key3
Awesome fact 1
Awesome fact 2
Awesome fact 3
```

### Contribution

Contribution is welcome! Send your pull requests with awesome facts or new functionality

### Licence

You can use this library "as is" for any projects