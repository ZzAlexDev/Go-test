package main

/*
Hello, 世界
Welcome to a tour of the Go programming language.
The tour is divided into a list of modules that you can access by clicking on A Tour of Go on the top left of the page.
You can also view the table of contents at any time by clicking on the menu on the top right of the page.
Throughout the tour you will find a series of slides and exercises for you to complete.
You can navigate through them using
"previous" or PageUp to go to the previous page,
"next" or PageDown to go to the next page.
The tour is interactive. Click the Run button now (or press Shift + Enter) to compile and run the program on a remote server. The result is displayed below the code.
These example programs demonstrate different aspects of Go. The programs in the tour are meant to be starting points for your own experimentation.
Edit the program and run it again.
When you click on Format (shortcut: Ctrl + Enter), the text in the editor is formatted using the gofmt tool. You can switch syntax highlighting on and off by clicking on the syntax button.
When you're ready to move on, click the right arrow below or type the PageDown key.
*/

/*
Привет, Друг
Добро пожаловать в экскурсию по языку программирования Go .
Тур разделен на список модулей, к которым вы можете получить доступ, щелкнув A Tour of Go в левом верхнем углу страницы.
Вы также можете просмотреть оглавление в любое время, щелкнув меню в правом верхнем углу страницы.
На протяжении всего тура вы найдете серию слайдов и упражнений для выполнения.
Вы можете перемещаться по ним, используя
«предыдущая» или PageUpперейти на предыдущую страницу,
«Далее» или PageDownдля перехода на следующую страницу.
Экскурсия интерактивная. Нажмите кнопку « Выполнить сейчас» (или нажмите Shift+ Enter), чтобы скомпилировать и запустить программу на удаленном сервере. Результат отображается под кодом.
Эти примеры программ демонстрируют различные аспекты Go. Программы тура должны стать отправной точкой для ваших собственных экспериментов.
Отредактируйте программу и запустите ее снова.
Когда вы нажимаете « Формат» (сочетание клавиш: Ctrl+ Enter), текст в редакторе форматируется с помощью инструмента gofmt . Вы можете включать и выключать подсветку синтаксиса, нажимая кнопку синтаксиса .
Когда вы будете готовы двигаться дальше, нажмите стрелку вправо ниже или введите PageDownклавишу .
*/

/*
func main() {
	fmt.Println("Hello, Friends")
}
*/

/*

	func main() {
		fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	}

	func main() {
		fmt.Println(math.Pi)
	}

	func add(x int, y int) int {
		return x + y
	}

	func main() {
		rand.Seed(time.Now().UnixNano())
		x := rand.Intn(100)
		y := rand.Intn(100)
		fmt.Println("My number is", x)
		fmt.Println("My number is", y)
		fmt.Println(add(x, y))
	}


func swap(x, y string) (string, string) {
	return x, y
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}


func split(sum int) (x, y int) {
	x = sum * 1 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}


//Variables
//The var statement declares a list of variables; as in function argument lists, the type is last.
//A var statement can be at package or function level. We see both in this example.

func main() {
	var c, python, java bool
	var i int
	fmt.Println(i, c, python, java)
}


//Variables with initializers
//A var declaration can include initializers, one per variable.
//If an initializer is present, the type can be omitted; the variable will take the type of the initializer.



var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

//Short variable declarations
//Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
//Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

//Краткие объявления переменных
//Внутри функции := короткий оператор присваивания может использоваться вместо var объявления с неявным типом.
//Вне функции каждый оператор начинается с ключевого слова (var, func, и т. д.), поэтому := конструкция недоступна.

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}


//Basic types
//Go's basic types are
//bool
//string
//int  int8  int16  int32  int64
//uint uint8 uint16 uint32 uint64 uintptr
//byte	// alias for uint8
//rune	// alias for int32
//		// represents a Unicode code point
//float32 float64
//complex64 complex128
//The example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.
//The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.

//Основные типы
//Основные типы Go:
//bool
//string
//int int8 int16 int32 int64
//uint uint8 uint16 uint32 uint64 uintptr
//byte // псевдоним для uint8
//rune // псевдоним для int32 // представляет кодовую точку Unicode
//float32 float64 complex64 complex128
//В примере показаны переменные нескольких типов, а также то, что объявления переменных могут быть «разнесены» на блоки, как и в случае операторов импорта.
//Типы int, uint, и uintptrобычно имеют ширину 32 бита в 32-битных системах и 64 бита в 64-битных системах. Когда вам нужно целочисленное значение, вы должны использовать intего, если только у вас нет особой причины использовать размерный или беззнаковый целочисленный тип.

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 11i)
)

func main() {
	fmt.Printf("Type: %T, Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T, Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T, Value: %v\n", z, z)
}


//	Zero values
//	Variables declared without an explicit initial value are given their zero value.
//	The zero value is:

//	0 for numeric types,
//	false for the boolean type, and
//	"" (the empty string) for strings.

//	Нулевые значения
//	Переменным, объявленным без явного начального значения, присваивается нулевое значение .
//	Нулевое значение:
//	0	для числовых типов,
//	false	для логического типа и
//	""	(пустая строка) для строк.

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

}


// Type conversions
//	The expression T(v) converts the value v to the type T.
//	Some numeric conversions:
//		var i int = 42
//		var f float64 = float64(i)
//		var u uint = uint(f)
//	Or, put more simply:
//		i := 42
//		f := float64(i)
//		u := uint(f)
//	Unlike in C, in Go assignment between items of different type requires an explicit conversion. Try removing the float64 or uint conversions in the example and see what happens.

//	Преобразования типов
//	Выражение T(v)преобразует значение vв тип T.
//	Некоторые числовые преобразования:
//		var i int = 42
//		var f float64 = float64(i)
//		var u uint = uint(f)
//	Или, говоря проще:
//		i := 42
//		f := float64(i)
//		u := uint(f)
//	В отличие от C, в Go присваивание между элементами разного типа требует явного преобразования. Попробуйте удалить преобразования float64или в примере и посмотрите, что произойдет.uint

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}


//Type inference
//When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.
//When the right hand side of the declaration is typed, the new variable is of that same type:
//		var i int
//		j := i // j is an int

//	But when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depending on the precision of the constant:
//		i := 42           // int
//		f := 3.142        // float64
//		g := 0.867 + 0.5i // complex128

//	Try changing the initial value of v in the example code and observe how its type is affected.

//	Вывод типа
//	При объявлении переменной без указания явного типа (с использованием :=синтаксиса или синтаксиса var =выражения) тип переменной выводится из значения в правой части.
//	Когда вводится правая часть объявления, новая переменная имеет тот же тип:
//		var i int
//		j := i // j является целым числом
//	Но когда правая часть содержит нетипизированную числовую константу, новая переменная может быть int, float64или complex128в зависимости от точности константы:
//		i := 42 // int
//		f := 3,142 // float64
//		g := 0,867 + 0,5i // комплекс128
//	Попробуйте изменить начальное значение vв примере кода и посмотрите, как это повлияет на его тип.


func main() {
	v := 42 // change me! //измени меня
	fmt.Printf("v is of type %T\n", v)
}


//	Constants
//	Constants are declared like variables, but with the const keyword.
//	Constants can be character, string, boolean, or numeric values.
//	Constants cannot be declared using the := syntax.

//	Константы
//	Константы объявляются как переменные, но с constключевым словом.
//	Константы могут быть символьными, строковыми, логическими или числовыми значениями.
//	Константы не могут быть объявлены с использованием :=синтаксиса.


const Pi = 3.14

func main() {
	const World = "FRIENDS"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	const Truth = true
	fmt.Println("Go rules?", Truth)
}


//	Numeric Constants
//	Numeric constants are high-precision values.
//	An untyped constant takes the type needed by its context.
//	Try printing needInt(Big) too.
// (An int can store at maximum a 64-bit integer, and sometimes less.)

//	Числовые константы
//	Числовые константы являются высокоточными значениями.
//	Нетипизированная константа принимает тип, необходимый для ее контекста.
//	Попробуйте needInt(Big)тоже распечатать.
//	(int может хранить максимум 64-битное целое число, а иногда и меньше.)

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.

	// Создать огромное число, сдвинув 1 бит влево на 100 позиций.
	// Другими словами, двоичное число, состоящее из 1, за которым следуют 100 нулей.
	Big = 1 << 64

	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	// Снова сдвинем вправо на 99 позиций, так что получим 1<<1 или 2.
	Small = Big >> 62
)

func needInt(x uint64) uint64 { return x*10 - 1 }

func needFloat(x float64) float64 { return x * 0.1 }

func main() {
	fmt.Println(needInt(Big))
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}


//========================== FOR =================================

// For
// Go has only one looping construct, the for loop.
// The basic for loop has three components separated by semicolons:
// the init statement: executed before the first iteration
// the condition expression: evaluated before every iteration
// the post statement: executed at the end of every iteration
// The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement.
// The loop will stop iterating once the boolean condition evaluates to false.
// Note: Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.

//	Цик For
//	В Go есть только одна циклическая конструкция — for цикл.
//	Базовый for цикл состоит из трех компонентов, разделенных точкой с запятой:
//	оператор инициализации: выполняется до первой итерации
//	выражение условия: оценивается перед каждой итерацией
//	оператор post: выполняется в конце каждой итерации
//	Оператор init часто представляет собой короткое объявление переменной, а объявленные в нем переменные видны только в области действия for оператора.
//	Цикл прекратит повторение, как только логическое условие будет оценено как false.
//	Примечание. В отличие от других языков, таких как C, Java или JavaScript, здесь нет круглых скобок, окружающих три компонента for оператора, и фигурные скобки { } всегда необходимы.

func main() {
	z := 0
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
		fmt.Println(i, sum)
		z = i
	}

	fmt.Println(sum / z)

}


*/
