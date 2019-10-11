**Задание:**

Есть 3 yaml файла следующего содержания

a) содержимое файла project_global.yaml
```depends:
project:
	name: MyFirstConfigProject
	debug: false
web:
	port: 80
service:
	google:
		host: http://google.com
		enabled: true
```

b) содержимое файла project_qa.yaml
```depends: global
project:
	name: QaProject
	debug: true
service:
	google:
		host: http://qa.google.com
		enabled: false
```

c) содержимое файла project_local.yaml
```depends: qa
web:
	port: 8080
service:
	google:
		enabled: true
```

Необходимо написать пакет, который сканирует все вышеуказанные файлы и наполняет структуру данных значениями из файлов. Необходимо учитывать, что ключ depends означает, что значения наследуются от содержимого файла, которое формируется на основании значения из поля depends. Например depends: qa означает что структура данных заполнится сначала значениями из project_qa.yaml, а затем значения из текущего файла перепишут значения из project_qa.yaml. Наполнение данных должно происходить в зависимости от передаваемого ключа (local/qa/global)
Пакет не должен зависит от заполняемой структуры данных.




Тип структуры, которая должна быть заполнена и подается на вход пакету вместе с ключем (local/qa/global)
```
type Config struct {
	Project struct {
		Name string
		Debug bool
	}
	Web struct {
		Port int
	}
	Service struct {
		Google struct {
			Host string
			Enabled bool
		}
	}
}
```

**Решение:**
Для изменения ключа необходимо в файле Dockerfile изменить 22 строчку, задав необходимый ключ.
Правильно было бы ключ брать из environment'а, но я следовал заданию.

Далее в терминале в корне проекта выполнить команду `docker-compose up --build`
В ответе будет содержаться конфиг, соответствующий ключу.