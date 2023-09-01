# Music-player service
### Тестовое задание для поступления в GoCloudCamp.

Cервис позволяющий управлять музыкальным плейлистом. Также можно выполнять CRUD операции с песнями в плейлисте, а также воспроизводить, приостанавливать, переходить к следующему и предыдущему трекам.

Модуль обладает следующими возможностями:
 - Play - начинает воспроизведение
 - Pause - приостанавливает воспроизведение
 - AddSong - добавляет в конец плейлиста песню
 - Next воспроизвести след песню
 - Prev воспроизвести предыдущую песню
 
### Для локального запуска:
 
Клонировать сам проект

```bash
  git clone https://github.com/zhas-off/test-music-player.git
```

Перейти в директорию проекта

```bash
  cd test-music-player
```

Запустить приложение

```bash
  make build
  make run
```

После запуска сервера, в другом терминале выполнить команду

```bash
  make evans
```

## Примеры использования
В клиенте evans можно выполнить следующие команды. Само приложение работает в стандартном порту http://localhost:50051
```bash
  call AddSong
  name (TYPE_STRING) => some_music
  time (TYPE_STRING) => 00:01:15
```
```bash
  call DeleteSong
  name (TYPE_STRING) => some_music
```
```bash
  call PlaySong
```
```bash
  call Next
```
```bash
  call PauseSong
```
```bash
  call Prev
```
