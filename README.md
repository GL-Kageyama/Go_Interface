# Go Interface Sample

## Interface
Interfaces are considered separately from implementation and use.  
And the Go interface is also used to resolve "import cycle".  

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
