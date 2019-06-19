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
Картинка 1  терминал с транзакцией
В конце выполним `docker ps -a`, чтобы посмотреть запущенные контейнеры. У вас должно получиться следующее:
Картинка 2 терминал с контейнерами
Внимательно проверьте, чтобы не было упавших контейнеров, в противном случае необходимо выполнить `startFabric.sh` еще раз.

Если все отлично, то идем по инструкции в терминале:
Нам нужно выполнить установку `node_modules` для клиента и для сервера. Сначала в корне проекта выполним `npm install`. После завершения установки переходим в директорию `client` и выполняем тот же `npm install`.
Теперь в папке `scripts` выполняем `node enrollAdmin.js` и после него `node registerUser.js`.

This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).

## Available Scripts

In the project directory, you can run:

### `npm start`

Runs the app in the development mode.<br>
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

The page will reload if you make edits.<br>
You will also see any lint errors in the console.

### `npm test`

Launches the test runner in the interactive watch mode.<br>
See the section about [running tests](https://facebook.github.io/create-react-app/docs/running-tests) for more information.

### `npm run build`

Builds the app for production to the `build` folder.<br>
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.<br>
Your app is ready to be deployed!

See the section about [deployment](https://facebook.github.io/create-react-app/docs/deployment) for more information.

### `npm run eject`

**Note: this is a one-way operation. Once you `eject`, you can’t go back!**

If you aren’t satisfied with the build tool and configuration choices, you can `eject` at any time. This command will remove the single build dependency from your project.

Instead, it will copy all the configuration files and the transitive dependencies (Webpack, Babel, ESLint, etc) right into your project so you have full control over them. All of the commands except `eject` will still work, but they will point to the copied scripts so you can tweak them. At this point you’re on your own.

You don’t have to ever use `eject`. The curated feature set is suitable for small and middle deployments, and you shouldn’t feel obligated to use this feature. However we understand that this tool wouldn’t be useful if you couldn’t customize it when you are ready for it.

## Learn More

You can learn more in the [Create React App documentation](https://facebook.github.io/create-react-app/docs/getting-started).

To learn React, check out the [React documentation](https://reactjs.org/).
