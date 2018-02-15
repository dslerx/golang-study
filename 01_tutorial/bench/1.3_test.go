package main

import (
    "testing"
    "os"
    "strings"
)

func BenchmarkCaseA(b *testing.B) {
    b.ResetTimer()
    // Nはコマンド引数から与えられたベンチマーク時間から自動で計算される
    for i := 0; i < b.N; i++ {
        var s, sep string
        for i := 1; i < len(os.Args); i++ {
          s += sep + os.Args[i]
          sep = " "
        }
    }
}

func BenchmarkCaseB(b *testing.B) {
    b.ResetTimer()
    // Nはコマンド引数から与えられたベンチマーク時間から自動で計算される
    for i := 0; i < b.N; i++ {
        strings.Join(os.Args[1:], " ")
    }
}