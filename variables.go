//En Go existen diferentes tipos de variables, por ejemplo:

//int- almacena números enteros (números enteros), como 123 o -123
//float32- almacena números de coma flotante, con decimales, como 19,99 o -19,99
//string- almacena texto, como "Hola mundo". Los valores de cadena están entre comillas dobles.
//bool- almacena valores con dos estados: verdadero o falso

var VARIABLENAME type = value 

variablename := value 

package main
import ("fmt")

func main() {
  var student1 string = "John" //type is string
  var student2 = "Jane" //type is inferred
  x := 2 //type is inferred

  fmt.Println(student1)
  fmt.Println(student2)
  fmt.Println(x)
}


//Declaración de variable sin valor inicial
//En Go, todas las variables se inicializan. 
//Entonces, si declaras una variable sin un valor inicial, su valor se establecerá en el valor predeterminado de su tipo:

package main
import ("fmt")

func main() {
  var a string
  var b int
  var c bool

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}

package main
import ("fmt")

func main() {
  var student1 string
  student1 = "John"
  fmt.Println(student1)
}



package main
import ("fmt")

var a int
var b int = 2
var c = 3

func main() {
  a = 1
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}

package main
import ("fmt")

func main() {
  var a, b, c, d int = 1, 3, 5, 7

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
}

package main
import ("fmt")

func main() {
  var a, b = 6, "Hello"
  c, d := 7, "World!"

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
}