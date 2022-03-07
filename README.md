# sitemap


Проект по созданию карты сайтов. Собирает все линки целевого сайта в xml вида:


```xml
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>http://www.example.com/</loc>
  </url>
  <url>
    <loc>http://www.example.com/dogs</loc>
  </url>
</urlset>
```


Что соответствует [стандартному протоколу sitemap](https://www.sitemaps.org/index.html)

Имеется две опции для использования:

--domain - адрес публичнго сайта, для которого будут собираться линки. Если опция не указана по умолчанию будет собрана карта для адреса http://127.0.0.1:8080/

--depth - максимальная глубина поиска. Указывает сколько циклов прохода на новую глубину линков целевого сайта требуется выполнить. По умолчанию равна 3.



Для запуска требуется скопировать проект и из его папки выполнить:

```bash
go run ./main.go --domain=http://exhample.com --depth=10
```

Для получения описания опций можно выполнить:

```bash
go run ./main.go --help
```

Проект использует модуль github.com/raptor72/glink. Вам необходимо будет выполнить go get для его установки. 


Идея для проекта взята из курса Джона Колтона https://courses.calhoun.io/lessons/les_goph_24 ( ссылка на гитхаб: https://github.com/gophercises/sitemap
)
