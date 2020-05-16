# Создание динамической библиотеки на языке Golang

### Задание:
<details>
<summary>Вариант 30</summary>
Составить  функцию,  проверяющую,  являются  ли  данное  число  простым.  Вторая функция должна вернуть ближайшее большее данного простое число. Например, для числа 11 первая функция возвращает true, вторая –13. Для числа 14 первая функция должна вернуть false, а вторая –17
</details>
___

#### Описание:
Файл **functions.go** содержит функции, которые необходимо создать по заданию
Для компиляции в динамическую библиотеку надо выполнить команду:
```console
[dima@localhost SUAI]$ go build -o chudinov.so -buildmode=c-shared functions.go
```
Файл **functions_test.go** содержит набор unit-тестов для проверки корректности работы функций 
Для  запуска тестов надо выполнить команду:
```console
[dima@localhost SUAI]$ go test -v
```
<details>
<summary>Пример вывода:</summary>
```console
[dima@localhost SUAI]$ go test -v
=== RUN   TestIsSimple
--- PASS: TestIsSimple (0.00s)
=== RUN   TestNextSimple
--- PASS: TestNextSimple (0.00s)
=== RUN   TestByTK
--- PASS: TestByTK (0.00s)
    functions_test.go:46: Test IsSimple function on simple digit in TK
    functions_test.go:51: Test IsSimple function on not simple digit in TK
    functions_test.go:56: Test NextSimple function on data in TK
PASS
ok      dll     0.003s
```
</details>