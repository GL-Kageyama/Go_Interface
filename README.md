# Go Interface Sample

## Interface
Interfaces are considered separately from implementation and use.  
And the Go interface is also used to resolve "circular reference".  

### Circular reference
The Go language prohibits circular references to packages.  
A circular reference means that packages A and B refer to each other.  
This specification contributes to a layered architecture.  

## Structure
dog/  
&ensp;&ensp;Dog.go  
&ensp;&ensp;DogLoar.go  
cat/  
&ensp;&ensp;Cat.go  
&ensp;&ensp;DogInterface.go  
main.go  

## Point
Dog.go  
 
```Go
func (d *Dog) GetDog() string {  
  cat := &cat.Cat{ CatObj: NewDogRoar() }  
  // GetCat -> DogInterface -> GetDogRoar
  return cat.GetCat()   
} 
```
note : In the GetDog function, the GetCat function is called, but the result is "Bow Wow !". 

## Output Sample
$ go build  
$ ./Go_Interface   
Bow Wow !  
