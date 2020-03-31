## jvm-go

Simple JVM written in Golang

Mainly referenced from [zxh0/jvm.go](https://github.com/zxh0/jvm.go) , and added many Chinese comments

This project is used to familiarize myself with Golang and JVM. Actually, it works well

## Todo

- [x] command line startup

<details><summary>supported arguments</summary><pre><code>
done:
	version
  	verbose
  	verbose:class
  	verbose:inst
  	classpath
  	cp
  	Xjre
  	class
  	args
todo:
	Xms
	Xmx
	Xss
</code></pre></details>

- [x] load and parse class file
- [x] runtime data
- [x] JVM instructions ( 195 instructions completed )
- [x] method call
- [x] array support
- [x] string support
- [x] native method call
- [x] exception handling
- [ ] implement native methods...
- [ ] multifaceted optimization...

