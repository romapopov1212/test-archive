# test-archive

# 1) собрать бинарник - go build -o pm ./cmd

# 2) пример packet.json - 
{
   "name": "packet-1",
   "ver": "1.0.0",
   "targets": [
       "./testdata/*.txt"
   ],
   "packets": []
}
# 3) запуск - ./pm create packet.json

на данный момент настройки сервера захардкожены, находятся в main.go