Setlocal EnableDelayedExpansion
set GOPATH=%~dp0

IF %1%==atom(
  echo "atom"
)

if %1%==install(
  :: Install pin
  pushd %GOPATH%src\github.com\xolan\pin
  echo "go install"
  popd
)

if %1%==run(
  pushd %GOPATH%bin
  echo "pin" 
  popd
)
