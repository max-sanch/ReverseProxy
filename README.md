# Тестовое задание для Effective Technologies


* [Инструкции](#guides)
    * [Запуск приложения](#launch-app)
        * [На Windows](#Windows)
        * [На Linux](#Linux)

## <a name="guides"></a> Инструкции

### <a name="launch-app"></a>Запуск приложения

Для запуска приложения вам необходимы `Docker`, `docker-compose` и `git`.
<br>
При изменении порта в config файле, необходимо указать тот же порт в файле docker-compose.


#### <a name="Windows"></a>На Windows

1) Установите git и [Docker Desktop](https://www.docker.com/products/docker-desktop)

2) Скачайте проект с GitHub:

        git clone https://github.com/max-sanch/ReverseProxy.git

3) Перейдите в деректорию проекта и введите:

       docker-compose up --build

#### <a name="Linux"></a>На Linux

1) Установите Docker

2) Установите docker-compose:

        sudo apt install docker-compose

3) При необходимости установите git

4) Скачайте проект с GitHub:

        sudo git clone https://github.com/max-sanch/ReverseProxy.git

5) Перейдите в деректорию проекта и введите:

       sudo docker-compose up --build