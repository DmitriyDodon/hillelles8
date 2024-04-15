# hillelles8

в корне проекта выпольнить команду <pre>cp .env.example .env</pre>

в корне проекта создать файл data.db с помощью <pre>touch data.db</pre> - это файлик для базы данных

роуты в достыпные в приложении 

route | method | description
_____________________________________
/car       |  POST    | создание машины
           |          | {
           |          |      "color": "#8f80de",
           |          |      "price_in_cents": 500000,
           |          |      "max_speed_mph": 140,
           |          |      "max_speed_kmp": 280,
           |          |      "vendor_name": "Honda",
           |          |      "model_name": "S2000"
           |          |  }
____________________________________
 /car/list | GET      | вывод списка всех созданых машин
 ___________________________________
/car/:carId| GET      | получения конкретной машины
____________________________________
/car/:carId| PUT      | Обновление машины
           |          | {
           |          |      "color": "#8f80de",
           |          |      "price_in_cents": 500000,
           |          |      "max_speed_mph": 140,
           |          |      "max_speed_kmp": 280,
           |          |      "vendor_name": "Honda",
           |          |      "model_name": "S2000"
           |          |  }
____________________________________
/car/:carId|DELETE    | удаления машины

