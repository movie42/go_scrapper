# Go를 이용해서 scrapper 만들기

> update 2021.04.07 | ver.0.1
> [Nomad Coder의 쉽고 빠른 Go 시작하기](https://nomadcoders.co/go-for-beginners/lobby)를 들으면서 따라 만들어본 scrapper다.

## 왜 고?

go lang은 유튜브를 보다가 처음 알게 되었다. 구글에서 멀티 스레드를 위해 만든 언어라고 한다. 파이썬 스크래퍼도 만들어봤지만 파이썬 스크래퍼와 차원이 다른 빠르기를 보여준다. 아직 이 언어를 잘 이해하고 있는 건 아니지만 일단 go라는 언어를 체험해보는 수준에서 만들어보았다.

## 사용방법

git pull을 하여 go run main.go를 실행하면 로컬 서버가 실행된다.
로컬 서버에서 job을 입력하고 (예:javascript) 앤터를 치면 바로 csv파일이 다운로드 된다. 속도가 미친듯이 빠르다.
