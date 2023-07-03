# Golang-training
Basics of golang programming
So far I have completed the concepts of variables, arrays, slices, for loops, if-else statements, modules like 'fmt','time','math','unicode/utf8' etc,functions, pointers, recursion, struct, methods, interfaces, struct embedding. 
Slices are better to use than arrays.
make statement is used in arrays,slices,maps etc. Empty map can be created using make(map[key]value). 'Delete' used in map by delete(map_name,"key").
range are used in arrays,slices,maps,keys(maps),string using unicode code points. rune values are there for characters. 
Functions declarartion contains return data type. subset of returned values can be obtained by blank identifier.
Variadic functions can be used to pass any number of arguments. Anonymous functions are inline functions declared without names to form closures. 
Closures can also be recursive, but this requires the closure to be declared with a typed var explicitly before it’s defined.
Go string literals are encoded with UTF8 text. Rune count can be calculated using UTF8 package(RuneCountInString) and decoded using DecodeRuneInString.
Values enclosed within '' are rune literals which can be compared with rune value.
Structure are mutable. Pointers are derefrenced automatically when accessed with dot. 
If a struct type is only used for a single value, we don’t have to give it a name. The value can have an anonymous struct type. This technique is commonly used for table-driven tests.
Methods can be defined for either pointer or value receiver types(*rect,rect).
