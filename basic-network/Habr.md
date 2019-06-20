# Описание программной части проекта

Добрый день, дорогие хабровчане! Меня зовут Роман Комаров, я работаю в IBM техническим архитектором. Если вы еще не слышали о применении технологии Blockchain в сфере бизнеса, то настоятельно рекомендую ознакомиться со статьей моего коллеги, который повествует об [архитектуре и теоритической части Hyperledger Fabric](#https://habr.com/ru/company/ibm/blog/444874/), а также рекомендую ознакомиться с тем, как можно [использовать kubernetes для разработки блокчейн проектов](#https://habr.com/ru/company/ibm/blog/351808/).

Помимо основной работы в нашей компании существуют возможности развиваться в других различных направлениях. Таким образом появился IBM Garage. Это то место, где опытные коллеги делятся знаниями с людьми, которые хотят расширить свои познания в той или иной области. Мы с коллегами в гараже изучаем трендовые технологии с разных сторон, чтобы лучше понять их и разобраться. Так я открыл для себя технологию Blockchain с новой, неизвестной мне стороны.
Я и мои коллеги готовим для вас целый цикл статей на тему практического применения технлогии Blockchain. Хотелось бы отметить, что наши публикации имеют несколько целей:
* познакомить с Hyperledger Fabric;
* познакомить с React;
* подготовить читателя к восприятию программного кода, который используется в проекте;

Если вы еще не знакомы и очень хочется попробовать, то можно перейти на [github нашего проекта](https://github.com/komaroman/StudReg), и не забудьте [убедиться в наличии необходимого ПО](#app_a) рекомендуемой версии.

## Chaincode

В Hyperledger Fabric для определения смарт-контрактов используется термин chaincode. Chaincode (Smart Contract) - инструмент контролируемого доступа к данным в реестре, а также реализация бизнес-логики транзакций, выполняемых в сети. Чейнкод устанавливается на пиры и запускается в отдельных Docker контейнерах. Для написания чейнкода используются языки высокого уровня. В текущей версии Fabric есть API для Go, node.js и Java. Chaincode инициализирует состояние ledger'а и управляет им с помощью транзакций, отправляемых приложениями. Для того, чтобы чейнкод был воспринят пиром, необходимо реализовать интерфейс чейнкода, описанный в соответствующем API. Например в Go это интефейс [Chaincode](https://github.com/hyperledger/fabric/blob/release-1.2/core/chaincode/shim/interfaces.go).
В нашем примере реализация интерфейса Chaincode выглядит следующим образом:
```go
type MainChaincode struct {}

func (cc *MainChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	return shim.Success(nil)
}

func (cc *MainChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	function, args := stub.GetFunctionAndParameters()

	switch function {
        case "createStudent":
            return CreateStudent(stub, args)
        case "queryStudent":
            return QueryStudent(stub, args)
        case "updateStudent":
            return UpdateStudent(stub, args)
        default:
            return shim.Error("Unknown function")
	}
}
```
Метод `Init` вызывается, когда chaincode инициализируется в канале. Метод `Invoke` вызывается в начале выполнения транзакций в сети. Получает транзакцию создания экземпляра или обновления, чтобы chaincode мог выполнить любую необходимую инициализацию, включая инициализацию состояния приложения. Метод `Invoke` вызывается в ответ на получение транзакции `invoke` для обработки транзакций предложений.
```go
package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	err := shim.Start(new(MainChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

```

### Package, import

Go разработан как язык, который поощряет хорошие инженерные практики. Одной из таких практик является возможность многократно использовать фрагменты кода, используя package. Package `main` сообщает компилятору Go, что пакет должен компилироваться как исполняемая программа, а не как `sharedlibrary`. Основной функцией в package `main` будет точка входа нашей исполняемой программы. При построении общих библиотек в пакете не будет `main package` и функции `main`.

Директива `import` используется для того, чтобы подключать функционал из других пакетов. В коде мы импортировали пакет `fmt` для использования функции `Printf`. Пакет `fmt` поставляется из стандартной библиотеки Go. При импорте пакетов компилятор Go будет искать расположения, заданные переменными среды `GOROOT` и `GOPATH`. Пакеты из стандартной библиотеки доступны в папке `GOROOT`. Созданные вами пакеты и импортированные пакеты сторонних разработчиков доступны в расположении `GOPATH`.

Другой интерфейс в chaincode `shim` API - ChaincodeStubInterface:

* [Go](https://godoc.org/github.com/hyperledger/fabric/core/chaincode/shim#ChaincodeStubInterface)
* [node.js](https://fabric-shim.github.io/fabric-shim.ChaincodeStub.html)
* [Java](https://fabric-chaincode-java.github.io/org/hyperledger/fabric/shim/ChaincodeStub.html)

`Shim package` предоставляет API для chaincode для доступа к переменным состояния, контекста транзакции и взаимодействия с другим чейнкодом.

## NodeJS scripts

Основными компонентами для взаимодействия с Hyperledger Composer посредством NodeJS являются:
* `fabric-ca-client` - этот пакет инкапсулирует API для взаимодействия с Certificate Authority Fabric для управления жизненным циклом сертификатов пользователей, регистрацией, обновлением и т.д.
* `fabric-client` - этот пакет инкапсулирует API для взаимодействия с Peer'ами и Orderer'ами сети Fabric для установки и создания экземпляров чейнкода, отправки транзакций и выполнения запросов.

Для установки NodeJS модулей мы будем использовать npm.
Из корневой папки запускаем установку.
```sh
npm install
```
Далее необходимо установить Create React App.
```sh
cd client/
npm install
```

## Запуск HLF

Это готовый пример из `fabric samples`.
Ниже описан процесс разворачивания сети на локалькой машине.

```sh
export MSYS\_NO\_PATHCONV=1

starttime=$(date +%s)

LANGUAGE=${1:-"golang"}

CC\_NAME=main

CC\_SRC\_PATH=github.com/main

if ["$LANGUAGE" = "node" -o "$LANGUAGE" = "NODE"]; then

CC\_SRC\_PATH=/opt/gopath/src/github.com/fabcar/node

Fi

rm -rf ./hfc-key-store
# launch network; create channel and join peer to channel

cd ../basic-network

./start.sh
# Now launch the CLI container in order to install, instantiate chaincode
# and prime the ledger with our 10 cars

docker-compose -f ./docker-compose.yml up -d cli

docker exec -e "CORE\_PEER\_LOCALMSPID=Org1MSP" -e "CORE\_PEER\_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode install -n $CC\_NAME -v 1.0 -p "$CC\_SRC\_PATH" -l "$LANGUAGE"

docker exec -e "CORE\_PEER\_LOCALMSPID=Org1MSP" -e "CORE\_PEER\_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode instantiate -o orderer.example.com:7050 -C mychannel -n $CC\_NAME -l "$LANGUAGE" -v 1.0 -c &#39;{"Args":[""]}&#39; -P "OR (&#39;Org1MSP.member&#39;,&#39;Org2MSP.member&#39;)"

sleep 10
```
---
## Взаимодействие с веб сервером
В нашем примере мы будем использовать Create React App boilerplate + Express Server. Для их одновременного запуска в режиме разработки мы будем использовать `concurrently`. Также необходимо учесть, что express сервер и react app должны работать на разных портах, в нашем случае сервер на `3000` и приложение на `5000`.
```js
"dev": "concurrently \"npm run server\" \"npm run client\""
```
Express Server будет содержать всего 2 пути - сохранение в ledger (`invoke`) и возврат из ledger'a (`query`). В первом случае это будет обработка POST запроса вместе с параметрами из формы, а во втором обработка GET запроса с поиском в ledger'е по ключу (`ID`). Скрипт запуска express сервера выглядит следующим образом:
```js
const express = require('express');
const bodyParser = require('body-parser'); //понадобиться для быстрого и удобного извлечения тела запроса

const app = express();

var urlencodedParser = bodyParser.urlencoded({ extended: false });

app.post('/api/invoke', urlencodedParser, (req, res) => {
    //req processing => invoke
});

app.get('/api/query', (req, res) => {
    //req processing => query
});

const port = 5000;

app.listen(port, () => console.log(`Server started on port ${port}`));
```
### React App
Относительно Create React App наше приложение притерпело совсем небольшие изменения. Для быстрой верстки веб-страниц мы подключили библиотеку Bootstrap. Отдельные компоненты страницы были вынесены в `/src/components` для удобства. Хотелось бы отметить концептуальные вещи относительно наших основных компонентов - форм, данные с которых мы будем отправлять в Express server для обработки. В state компонента Invoke Form и Query Form мы буде хранить данные с формы. Также, у этого React компонента будет 2 обработчика: `handleSubmit` - для обработки заполненной формы и `handleChange` - для обработки изменений значений в input полях формы для соответственного изменения state.
```js
this.state = {
  studId: '',
  studFirstName: '',
  studLastName: '',
  studMiddleName: '',
  studPlaceOfBirth: '',
  studDateOfBirth: '',
  studPassNum: '',
  studGender: '',
  studMaritalStatus: '',
};
```
Для отправки данных на endpoint Express сервера мы использовали `fetch`.
```js
const INVOKE_API_URI = 'http://localhost:5000/api/invoke';

const fetchOpts = {
  method: 'post',
  mode: 'cors',
  headers: {
    "Content-Type": "application/json",
  },
  body,
};

fetch(INVOKE_API_URI, fetchOpts)
  .then((data) => console.log('Request succeeded with JSON response', data))
  .catch((error) => console.error('Request failed', error));
```
Запуск нашего приложения для взаимодействия с Hyperledger Fabric мы будем осуществлять с помощью `npm`.
```sh
npm run dev
```
Можно тестировать!

PS: Задавайте свои вопросы. Охватить все нюансы работы с Hyperledger Fabric в небольшой статье довольно таки сложно, поэтому тут рассматривается базовый сценарий, но понимание, как это работает должно открывать для вас много разных направлений для развития своих навыков, а также своих проектов.
Удачи!
#

## Полезные ресурсы
[GoDoc](https://godoc.org/)
[HyperLedger Documentation](https://hyperledger-fabric.readthedocs.io/en/release-1.3/getting_started.html)

## <a id="app_a"></a>Appendix A: Установка необходимого ПО

Для успешного запуска Blockchain сети нам потребуются:
- NodeJS и npm или yarn
- Docker
- Docker Compose
- Go language

### NodeJS
> Node.js version 9.x is not supported at this time.
> Node.js - version 8.9.x is required

[Install NodeJS](https://nodejs.org/en/download/releases/)

### Go Lang
>Go version 1.10.x is required.

[Install Go](https://golang.org/dl/)

### Docker and Docker Compose
- MacOSX, *nix, or Windows 10: [Docker Install](https://www.docker.com/get-started) (Docker version 17.06.2-ce or greater is required).
- Older versions of Windows: [Docker Toolbox](https://docs.docker.com/toolbox/toolbox_install_windows/) - again, Docker version Docker 17.06.2-ce or greater is required.