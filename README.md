# speaker

# find only dir

```bash
find . -type d
```

# Go mobile

##### ~/.zshrc
```bash
export GOPATH=~/go
export GOBIN=$GOPATH/bin
export GO111MODULE=on
export GOPROXY=direct
export GOSUMDB=off
export PATH=$PATH:$GOBIN:/usr/local/go/bin

export ANDROID_HOME=/home/bg/Android/Sdk
export ANDROID_NDK_HOME=/home/bg/Android/Sdk/ndk/21.3.6528147
```

install

```bash
go get golang.org/x/mobile/cmd/gomobile
$GOBIN/gomobile init
$GOBIN/gomobile bind -o app/hello.aar go
gomobile bind -o tmp/gomobilelib.aar -target=android github.com/exitstop/speakerandroid/gomobilelib
gomobile install github.com/exitstop/speakerandroid/network
```
