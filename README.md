# Разноцветные контейнеры

## Команды
- Запуск: 
``
make run
``
- Тестирование:
``
make test
``
- Сборка:
  ``
  make build
  ``
## Реализация

Представим входные данные в виде матрицы, где, допустим, по горизонтали будут цвета,
а по вертикали будут номера контейнеров. Тогда каждая строчка это будет содержимое контейнера.
Чтобы добиться желаемого результата, то есть чтобы в каждом контейнере был 1 цвет, надо чтобы
матрица приводилась к виду, когда у нас в матрице лежат шары только на диагонали, потому что это означает
что `i` цвет лежит в `i` -ом контейнере:
Разберем первый пример из задания, выйдет примерно так:

| №\Цвет | К | З |
|--------|---|---|
| 1      | 1 | 2 | 
| 2      | 2 | 1 | 

Если возможно разложить шары перекладыванием, то у нас должна получиться матрица вида

| №\Цвет | К | З |
|--------|---|---|
| 1      | 0 | 3 |
| 2      | 3 | 0 |

Мы взяли 1 красный шар в 1 контейнере и 1 зелеёный щар в 2 контейнере и переложили.
В моём алгоритме диагональ будет главная, это не играет роль, главное это отражает, что у нас в 1-ом контейнере 1 цвет.

По условию задания мы можем перекладывать только по 1 шару, но в алгоритме принято решение перекладывать по минимуму и
двух выбранных контейнеров. То есть в алгоритме мы переложим, допустим, не 5 раз по 1 шару из одного контейна в другой, 
а сразу 1 раз по 5 шаров, при этом ничего условию задания не противоречит, но мы значительно ускорим алгоритм. Условие
что 1 шар из одной кучи и 1 шар из другой сохарняется.

Суть алгоритма заключается в том, что мы выбираем контейнер и считаем, что в этом контейнере должны лежать только 
определенного цвета шары, остальных цветов шары в этом контейнере мы начинаем менять на красные в других, то есть 
постепенно в на контейнер забирать красные щары, а в другие класть шары цвета, в котором должны быть шары 
данного цвета. В итоге у нас постепенно в стобце и строке row1 не должно быть цветов никаких, то есть все элементы равны
0, кроме того, который стоит на главной диагонали.

После стоит выполнить проверки матрицы, что кроме диагональных элементов никаких больше нет в матрице. 

### Асимтотика:
Алгоритм перекладывания: O(n^2)\
Алгоритм проверки: O(n^2)\
В итоге: O(2*n^2) = O(n^2)\

> При n <= 100 асимтотика удовлетворительная


## Задание
Есть n контейнеров и какое-то количество шаров разного цвета, количество разных цветов тоже n. Шары в случайном порядке распределены по контейнерам. Допустима лишь одна операция с шарами: два шара в разных контейнерах можно поменять местами. Нужно ответить на вопрос: можно ли используя эту единственную операцию отсортировать шары так, чтобы выполнялись такие условия:
- в каждом контейнере лежат лишь шары одинакового цвета
- каждый цвет лежит лишь в одном контейнере

### Ограничения: 
- 1<= n <= 100 - количество контейнеров и цветов 
- 0 <= количество шаров одного цвета в одном контейнере <= 1000000000

### Формат входа: 
Первая строчка содержит n. Остальные n строчек содержат n чисел, разделённых пробелами. Каждая строчка описывает содержимое контейнера. Каждое из чисел обозначает количество шаров i-го цвета в этом контейнере, где i - порядковый номер числа в строчке
### Формат выхода: 
Написать "yes", если сортировку можно сделать, или "no", если нельзя.

### Примеры
#### Пример 1: 
##### Вход:
2\
1 2\
2 1

##### Выход: 
yes

Пояснение: у нас две коробки с (условно) красными и зелёными шарами. В первой коробке один красный и два зелёных, во второй - два красных и один зелёный. Если из первой коробки взять красный шар, из второй - зелёный и поменять их местами (допустимая операция), то мы получаем, что в первой коробке три зелёных шара и других нет, а во второй - три красных шара и других тоже нет. То есть получилось выполнить нужную сортировку.

#### Пример 2: 
##### Вход:
3\
10 20 30\
1 1 1\
0 0 1
##### Выход: 
no

Пояснение: у нас три коробки и три цвета, условно красный, зелёный и синий. В первой коробке 10 красных, 20 зелёных и 30 синих. Во второй - по одному шару каждого цвета. В третьей - лишь один синий шар. Как шары местами не меняй, но в первой коробке всегда будут все три цвета, то есть сделать сортировку не получается.
