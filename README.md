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
func (d *Dog) GetDog() string {  
&ensp;&ensp;&ensp;&ensp;cat := &cat.Cat{ CatObj: NewDogRoar() }  
&ensp;&ensp;&ensp;&ensp;return cat.GetCat()   
}  
note : In the GetDog function, the GetCat function is called, but the result is "Bow Wow !". 

## Output Sample
$ go build  
$ ./Go_Interface   
Bow Wow !  
