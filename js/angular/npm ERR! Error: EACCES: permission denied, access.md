错误代码
```js

> node-sass@4.9.3 install /home/wwwroot/laravel_shop/node_modules/laravel-mix/node_modules/node-sass
> node scripts/install.js

Unable to save binary /home/wwwroot/laravel_shop/node_modules/laravel-mix/node_modules/node-sass/vendor/linux-x64-64 : { Error: EACCES: permission denied, mkdir '/home/wwwroot/laravel_shop/node_modules/laravel-mix/node_modules/node-sass/vendor'
    at Object.mkdirSync (fs.js:750:3)
    at sync (/home/wwwroot/laravel_shop/node_modules/mkdirp/index.js:71:13)
    at Function.sync (/home/wwwroot/laravel_shop/node_modules/mkdirp/index.js:77:24)
    at checkAndDownloadBinary (/home/wwwroot/laravel_shop/node_modules/laravel-mix/node_modules/node-sass/scripts/install.js:114:11)
    at Object.<anonymous> (/home/wwwroot/laravel_shop/node_modules/laravel-mix/node_modules/node-sass/scripts/install.js:157:1)
    at Module._compile (internal/modules/cjs/loader.js:688:30)
    at Object.Module._extensions..js (internal/modules/cjs/loader.js:699:10)
    at Module.load (internal/modules/cjs/loader.js:598:32)
    at tryModuleLoad (internal/modules/cjs/loader.js:537:12)
    at Function.Module._load (internal/modules/cjs/loader.js:529:3)
  errno: -13,
  syscall: 'mkdir',
  code: 'EACCES',
  path:
   '/home/wwwroot/laravel_shop/node_modules/laravel-mix/node_modules/node-sass/vendor' }

> node-sass@4.9.3 postinstall /home/wwwroot/laravel_shop/node_modules/laravel-mix/node_modules/node-sass
> node scripts/build.js

Building: /usr/local/bin/node /home/wwwroot/laravel_shop/node_modules/node-gyp/bin/node-gyp.js rebuild --verbose --libsass_ext= --libsass_cflags= --libsass_ldflags= --libsass_library=

```

解决方案

在命令前加上 sudo

```bash
sudo npm install --save-dev grunt 
#不过这样子可能还是不行，你需要这样：

sudo npm install --unsafe-perm=true --save-dev grunt
#$或许你还是会遇到错误，请尝试这样：

sudo npm install --unsafe-perm=true --allow-root --save-dev grunt
```</anonymous>