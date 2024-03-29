# StudReg
Перед запуском blockchain сети необходимо удостовериться в наличии необходимого ПО:
1) Проверить версию Docker: docker --version (Поддерживаемая версия - 17.06.2-ce и выше) [Install Docker](https://www.docker.com/get-docker)
2) Docker Compose устанавливается вместе с Docker. Если у вас уже был установлен Docker устаревшей версии, то нужно проверить наличие и версию Docker Compose: docker-compose --version (Поддерживаемая версия - 1.14.0 и выше) [Install Docker Compose](https://www.docker.com/get-docker)
3) Чейнкод этого проекта написан на Go, поэтому необходимо установить Go версии 1.11.х и выше [Install Go](https://golang.org/dl/) | [Конфигурация Go](https://hyperledger-fabric.readthedocs.io/en/release-1.4/prereqs.html#go-programming-language)
4) Node.js Runtime и NPM - Поддерживается только 8-я мажорная версия. [Node.js v8.x](https://nodejs.org/en/download/)

[Более детальная инструкция по установке ПО.](https://hyperledger-fabric.readthedocs.io/en/release-1.4/prereqs.html)

## Запуск Hyperledger Fabric

Переходим в папку `scripts` и запускаем сеть.
```sh
cd scripts
sh ./startFabric.sh
```
Начнется процесс загрузки необходимых образов из Docker Hub. В конце должна получиться примерно такая картина. Проверьте на наличие ошибок, их быть не должно. Предупреждения могут быть, это не критично.
![](/img/img1.png)
В конце выполним `docker ps -a`, чтобы посмотреть запущенные контейнеры. У вас должно получиться следующее:
![](/img/img2.png)
Внимательно проверьте, чтобы не было упавших контейнеров, в противном случае необходимо выполнить `startFabric.sh` еще раз или проверить логи упавшего контейнера `docker logs $CONTAINER_ID`.

Теперь в папке `scripts` выполняем `node enrollAdmin.js` и после него `node registerUser.js`.

Если все отлично, то идем по инструкции в терминале:
Нам нужно выполнить установку `node_modules` для клиента и для сервера. Сначала в корне проекта выполним `npm install`. После завершения установки переходим в директорию `client` и выполняем тот же `npm install`.

Запуск нашего приложения для взаимодействия с Hyperledger Fabric мы будем осуществлять с помощью `npm`. Из корня проекта 
```sh
npm run dev
```
Вас должно перенаправить на страницу с React приложением. Адрес по умолчанию - http://localhost:3000/. С левой части страницы можно отправлять транзакции. С правой части страницы можно по ID делать поиск в ledger'е. В терминале, где был запущен dev сервер, будет выводится log со статусом транзакции.

Можно тестировать!

Текущее состояние ledger'а, а также информацию по всем блокам можно найти в базе данных канала в CouchDB. Адрес сервиса по умолчанию:
http://localhost:5984/_utils/

---

## Описание программной части

В Hyperledger Fabric для определения смарт-контрактов используется термин chaincode. Chaincode (Smart Contract) - инструмент контролируемого доступа к данным в реестре, а также реализация бизнес-логики транзакций, выполняемых в сети. Chaincode устанавливается на пиры и запускается в отдельных Docker контейнерах. Для написания chaincode используются языки высокого уровня. В текущей версии Fabric есть API для Go, node.js и Java. Chaincode инициализирует состояние ledger'а и управляет им с помощью транзакций, отправляемых приложениями. Для того, чтобы chaincode был воспринят пиром, необходимо реализовать интерфейс chaincod'а, описанный в соответствующем API. Например в Go это интефейс [Chaincode](https://github.com/hyperledger/fabric/blob/release-1.4/core/chaincode/shim/interfaces.go).
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

`Shim package` предоставляет API для chaincode для доступа к переменным состояния, доступа к контексту транзакции и взаимодействия с другим chaincod'ом.

## Node.js scripts

Основными компонентами для взаимодействия с Hyperledger Fabric посредством NodeJS являются:
* `fabric-ca-client` - этот пакет инкапсулирует API для взаимодействия с Certificate Authority Fabric для управления жизненным циклом сертификатов пользователей, регистрацией, обновлением и т.д.
* `fabric-client` - этот пакет инкапсулирует API для взаимодействия с Peer'ами и Orderer'ами сети Fabric для установки и создания экземпляров chaincod'а, отправки транзакций и выполнения запросов.

## Команды для взаимодействия peer'ов и chaincod'а в канале.

Это готовый пример из `fabric samples`.
Ниже описаны команды, необходимые для создания канала, добавления пира в канал а также `install` и `instantiate` у chaincode.

```sh
# Создание канала, создание Genesis block.
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@org1.example.com/msp" peer0.org1.example.com peer channel create -o orderer.example.com:7050 -c mychannel -f /etc/hyperledger/configtx/channel.tx

# Добавление в канал пира.
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@org1.example.com/msp" peer0.org1.example.com peer channel join -b mychannel.block

# При подключении нескольких организаций в один канал, необходимо выполнить fetch с указанием Genesis блока.
docker exec -e "CORE_PEER_LOCALMSPID=Org2MSP" -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@org2.example.com/msp" peer0.org2.example.com peer channel fetch config mychannel.block -o orderer.example.com:7050 -c mychannel

# Установка chaincode. Необходимо установить chaincode на каждом peer'е, который будет выполнять этот chaincode.
docker exec -e "CORE\_PEER\_LOCALMSPID=Org1MSP" -e "CORE\_PEER\_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode install -n $CC\_NAME -v 1.0 -p "$CC\_SRC\_PATH" -l "$LANGUAGE"

# После установки chaincode, необходимо создать экземпляр chaincode на канале, чтобы узлы могли взаимодействовать с ledger'ом через контейнер chaincode.
docker exec -e "CORE\_PEER\_LOCALMSPID=Org1MSP" -e "CORE\_PEER\_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode instantiate -o orderer.example.com:7050 -C mychannel -n $CC\_NAME -l "$LANGUAGE" -v 1.0 -c '{"Args":[""]}' -P "OR ('Org1MSP.member','Org2MSP.member')"
```

## Взаимодействие с веб сервером
В нашем примере мы будем использовать Create React App boilerplate + Express Server. Для их одновременного запуска в режиме разработки мы будем использовать `concurrently`. Также необходимо учесть, что express сервер и react app должны работать на разных портах, в нашем случае сервер на `3000` и приложение на `5000`.
```js
"dev": "concurrently \"npm run server\" \"npm run client\""
```
Express Server будет содержать всего 2 пути - сохранение в ledger (`invoke`) и возврат из ledger'a (`query`). В первом случае это будет обработка POST запроса вместе с параметрами из формы, а во втором обработка GET запроса с поиском в ledger'е по ключу (`ID`). Скрипт запуска express сервера выглядит следующим образом:
```js
const express = require('express');
const bodyParser = require('body-parser'); //понадобится для быстрого и удобного извлечения тела запроса

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
Относительно бойлерплэйта Create React App наше приложение притерпело совсем небольшие изменения. Для быстрой верстки веб-страниц мы подключили библиотеку Bootstrap. Отдельные компоненты страницы были вынесены в `/src/components` для удобства. Хотелось бы отметить концептуальные вещи относительно наших основных компонентов - форм, данные с которых мы будем отправлять в Express server для обработки. В state компонента Invoke Form и Query Form мы будем хранить данные с формы. Также, у этого React компонента будет 2 обработчика: `handleSubmit` - для обработки заполненной формы и `handleChange` - для обработки изменений значений в input полях формы для соответственного изменения state.
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

Больше информации об этом boilerplate'е можно найти тут, [Create React App documentation](https://facebook.github.io/create-react-app/docs/getting-started).

А также документация на библиотеку [React documentation](https://reactjs.org/).
