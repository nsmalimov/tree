### Использование:

1. Собрать исполняемый бинарник:
    
    Включить go mod: https://blog.golang.org/using-go-modules

    `cd cmd && go build -o tree`
    
    Также можно не собирать бинарник, а воспользоваться уже скомпилированным, который лежит в папке cmd.

2. * Показать структуру файлов в текущей директории: `./tree`

    * Показать структуру файлов в заданной директории: `./tree 'имя_директории'`

    * Показывать только папки: `./tree -d` или `./tree 'имя_директории' -d`
    
### Запуск тестов:
`go test`

Benchmarks: `go test -bench=.`
### Примечание

Пропускаются скрытые папки, начинающиеся с '.'.

Ещё не совсем понятно, выводить ли в самом начале при принте дерева текущую папку. В примерах работы в doc-файле в 
одном случае выводится, в другом нет. В итоге сделал так, чтобы если путь к папке указан, то вывожу указанный,
если не указан, то вывожу полный до текущей.

### Что можно улучшить

1. Парсинг аргументов командной строки можно вынести в отдельный package, и написать на него тесты.

2. В тестах я заранее руками создал тестовые папки со структурой как в doc. По-хорошему
надо создавать их из кода, а затем удалять. А саму структуру папок и файлов держать в дереве.

3. Сделать рекурсивные функции горутинами. В теории это должно увеличить производительность. Но скорее всего это будет не целесообразно.