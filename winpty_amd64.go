//go:build windows && amd64
// +build windows,amd64

package console

import "embed"

// winpty 二进制直接明文内嵌。
// 原方案通过 embed-encrypt 以 AES-GCM 加密内嵌（*.enc + key.enc），
// 但 key 与密文同存于编译产物中，仅为资源混淆、无实际防护意义，
// 且 embed-encrypt 上游无许可声明，故移除该依赖。
// 已验证：明文文件与原 *.enc 的解密结果逐字节一致。
//
//go:embed winpty/amd64/winpty.dll
//go:embed winpty/amd64/winpty-agent.exe
var winpty_deps embed.FS
