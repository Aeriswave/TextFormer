echo off
cls
title TextFormer
echo Start
pause
:: в Windows CMD консоли Git можно запускать, перечислив все используемые файлы
go run do.go run.go run.in.go run.of.go
echo End
pause
::exit

:: в консоли Git можно запускать так
:: go run do.go run*.go

:: или так
:: sh --login -i go.by.bat