# FAQ

> 免责声明：仅供参考。

## Origins(起源)
### What is the purpose of the project?
At the time of Go's inception, only a decade ago, the programming world was different from today. Production software was usually written in C++ or Java, GitHub did not exist, most computers were not yet multiprocessors, and other than Visual Studio and Eclipse there were few IDEs or other high-level tools available at all, let alone for free on the Internet.

Meanwhile, we had become frustrated by the undue complexity required to use the languages we worked with to develop server software. Computers had become enormously quicker since languages such as C, C++ and Java were first developed but the act of programming had not itself advanced nearly as much. Also, it was clear that multiprocessors were becoming universal but most languages offered little help to program them efficiently and safely.

We decided to take a step back and think about what major issues were going to dominate software engineering in the years ahead as technology developed, and how a new language might help address them. For instance, the rise of multicore CPUs argued that a language should provide first-class support for some sort of concurrency or parallelism. And to make resource management tractable in a large concurrent program, garbage collection, or at least some sort of safe automatic memory management was required.

These considerations led to a series of discussions from which Go arose, first as a set of ideas and desiderata, then as a language. An overarching goal was that Go do more to help the working programmer by enabling tooling, automating mundane tasks such as code formatting, and removing obstacles to working on large code bases.

A much more expansive description of the goals of Go and how they are met, or at least approached, is available in the article, Go at Google: Language Design in the Service of Software Engineering.


### 这个项目的目的是什么？

在 `Go` 诞生的时候，也就是十年前，编程世界与今天不同。生产软件通常是用 `C++` 或 `Java` 编写的，`GitHub` 还不存在，大多数计算机还没有多处理器，除了 `Visual Studio` 和 `Eclipse`，几乎没有IDE或其他高级工具，更不用说在互联网上免费使用了。

同时，我们对使用我们所使用的语言来开发服务器软件所需要的过度复杂性感到沮丧。自从 `C`、`C++` 和 `Java` 等语言被开发出来后，计算机的速度已经大大加快，但编程行为本身却没有进步这么多。此外，很明显，多处理器正在变得普遍，但大多数语言几乎没有提供什么帮助来有效和安全地对它们进行编程。

我们决定退一步思考，随着技术的发展，在未来的几年里，哪些主要问题将主导软件工程，以及一种新的语言如何帮助解决这些问题。例如，多核 `CPU` 的兴起，证明了语言应该为某种并发或并行提供一流的支持。而为了使大型并发程序中的资源管理变得易懂，需要垃圾收集，或者至少是某种安全的自动内存管理。

这些考虑导致了一系列的讨论，而 `Go` 就是从这些讨论中产生的，首先是作为一套想法和愿望，然后是作为一种语言。一个首要的目标是，`Go` 要做更多的事情来帮助工作中的程序员，实现工具化，自动化代码格式化等普通任务，并消除在大型代码库中工作的障碍。

关于 `Go` 的目标以及如何实现这些目标，或者至少是如何接近这些目标，可以在[《Go at Google: Language Design in the Service of Software Engineering》](https://talks.golang.org/2012/splash.article)一文中得到更广泛的描述。语言设计为软件工程服务。


### What is the history of the project?
Robert Griesemer, Rob Pike and Ken Thompson started sketching the goals for a new language on the white board on September 21, 2007. Within a few days the goals had settled into a plan to do something and a fair idea of what it would be. Design continued part-time in parallel with unrelated work. By January 2008, Ken had started work on a compiler with which to explore ideas; it generated C code as its output. By mid-year the language had become a full-time project and had settled enough to attempt a production compiler. In May 2008, Ian Taylor independently started on a GCC front end for Go using the draft specification. Russ Cox joined in late 2008 and helped move the language and libraries from prototype to reality.

Go became a public open source project on November 10, 2009. Countless people from the community have contributed ideas, discussions, and code.

There are now millions of Go programmers—gophers—around the world, and there are more every day. Go's success has far exceeded our expectations.

### 这个项目的历史是什么？
2007年9月21日，`Robert Griesemer`、`Rob Pike` 和 `Ken Thompson` 开始在白板上勾画新语言的目标。几天之内，这些目标就变成了一个计划，并对它的内容有了一个合理的想法。设计与无关的工作同时继续兼职。到2008年1月，Ken开始了一个编译器的工作，用它来探索各种想法；它的输出是 `C` 代码。到了年中，这门语言已经成为一个全职项目，并且已经稳定下来，可以尝试生产编译器。2008年5月，`Ian Taylor` 使用规范草案独立地开始了 `Go` 的 `GCC` 前端开发工作。`Russ Cox` 在2008年末加入，帮助将语言和库从原型变为现实。

2009年11月10日，`Go` 成为一个公开的开源项目。无数来自社区的人贡献了想法、讨论和代码。

现在全世界有数百万的 `Go` 程序员--`gophers`，而且每天都在增加。`Go` 的成功远远超出了我们的预期。

### What's the origin of the gopher mascot?
The mascot and logo were designed by Renée French, who also designed Glenda, the Plan 9 bunny. A blog post about the gopher explains how it was derived from one she used for a WFMU T-shirt design some years ago. The logo and mascot are covered by the Creative Commons Attribution 3.0 license.

The gopher has a model sheet illustrating his characteristics and how to represent them correctly. The model sheet was first shown in a talk by Renée at Gophercon in 2016. He has unique features; he's the Go gopher, not just any old gopher.

### 地鼠吉祥物的由来是什么？
吉祥物和标志是由 `Renée French` 设计的，她还设计了 `Plan 9` 的兔子 `Glenda`。在一篇关于地鼠的博客文章中，她解释了地鼠是如何从她几年前用于WFMU T恤设计的图案中衍生出来的。这个标志和吉祥物采用了知识共享署名3.0许可。

地鼠有一个模型表，说明了他的特征和如何正确地表现它们。该模型表最早是在2016年 `Renée` 在 `Gophercon` 的一次演讲中展示的。他有独特的特征，他是围棋地鼠，不是普通的地鼠。

### Is the language called Go or Golang?
The language is called Go. The "golang" moniker arose because the web site is golang.org, not go.org, which was not available to us. Many use the golang name, though, and it is handy as a label. For instance, the Twitter tag for the language is "#golang". The language's name is just plain Go, regardless.

A side note: Although the official logo has two capital letters, the language name is written Go, not GO.

### 这门语言是叫Go还是Golang？
这门语言叫做Go。产生 "golang "这个专名是因为网站是golang.org，而不是go.org，因为我们无法使用。不过很多人都在使用golang这个名字，它作为一个标签很方便。例如，该语言的 `Twitter` 标签是 "#golang"。不管怎么说，这种语言的名字就是普通的围棋。

附带说明：虽然官方标志有两个大写字母，但语言名称写的是 `Go` ，而不是 `GO` 。

### Why did you create a new language?
Go was born out of frustration with existing languages and environments for the work we were doing at Google. Programming had become too difficult and the choice of languages was partly to blame. One had to choose either efficient compilation, efficient execution, or ease of programming; all three were not available in the same mainstream language. Programmers who could were choosing ease over safety and efficiency by moving to dynamically typed languages such as Python and JavaScript rather than C++ or, to a lesser extent, Java.

We were not alone in our concerns. After many years with a pretty quiet landscape for programming languages, Go was among the first of several new languages—Rust, Elixir, Swift, and more—that have made programming language development an active, almost mainstream field again.

Go addressed these issues by attempting to combine the ease of programming of an interpreted, dynamically typed language with the efficiency and safety of a statically typed, compiled language. It also aimed to be modern, with support for networked and multicore computing. Finally, working with Go is intended to be fast: it should take at most a few seconds to build a large executable on a single computer. To meet these goals required addressing a number of linguistic issues: an expressive but lightweight type system; concurrency and garbage collection; rigid dependency specification; and so on. These cannot be addressed well by libraries or tools; a new language was called for.

The article Go at Google discusses the background and motivation behind the design of the Go language, as well as providing more detail about many of the answers presented in this FAQ.

### 你为什么要创造一种新的语言？
`Go` 的诞生源于我们对 `Google` 现有工作语言和环境的失望。编程变得太难了，语言的选择也是部分原因。人们必须选择高效的编译、高效的执行或编程的简易性；这三者在同一种主流语言中都不具备。有能力的程序员选择了简单而不是安全和效率，他们转向动态类型语言，如 `Python` 和 `JavaScript`，而不是 `C++`，或者在较小程度上选择 `Java`。

我们并不是一个人在担心。在经历了多年编程语言相当安静的景象之后，`Go` 是几个新语言--`Rust`、`Elixir`、`Swift`等等--中的第一批，这使得编程语言开发再次成为一个活跃的、几乎是主流的领域。

`Go` 解决了这些问题，它试图将解释型、动态类型语言的编程便利性与静态类型、编译语言的效率和安全性结合起来。它还旨在实现现代化，支持网络化和多核计算。最后，使用 `Go` 的目的是快速：在一台计算机上建立一个大型可执行文件最多只需要几秒钟。为了达到这些目标，需要解决一些语言问题：一个表达式但轻量级的类型系统；并发和垃圾收集；严格的依赖规范；等等。这些都是库或工具不能很好解决的，需要一种新的语言。

[《Go at Google》](https://talks.golang.org/2012/splash.article)这篇文章讨论了 `Go` 语言设计背后的背景和动机，并对本FAQ中提出的许多答案提供了更多细节。

### What are Go's ancestors?
Go is mostly in the C family (basic syntax), with significant input from the Pascal/Modula/Oberon family (declarations, packages), plus some ideas from languages inspired by Tony Hoare's CSP, such as Newsqueak and Limbo (concurrency). However, it is a new language across the board. In every respect the language was designed by thinking about what programmers do and how to make programming, at least the kind of programming we do, more effective, which means more fun.

### Go的祖先是什么？
`Go` 主要属于 `C` 家族（基本语法），从 `Pascal/Modula/Oberon` 家族（声明、包）中获得了大量的输入，再加上受 `Tony Hoare` 的 `CSP` 启发的语言的一些想法，如 `Newsqueak` 和 `Limbo`（并发）。然而，它是一门全面的新语言。在每一个方面，这门语言都是通过思考程序员的工作，以及如何让编程，至少是我们所做的那种编程，变得更有效，这意味着更有趣。

### What are the guiding principles in the design?
When Go was designed, Java and C++ were the most commonly used languages for writing servers, at least at Google. We felt that these languages required too much bookkeeping and repetition. Some programmers reacted by moving towards more dynamic, fluid languages like Python, at the cost of efficiency and type safety. We felt it should be possible to have the efficiency, the safety, and the fluidity in a single language.

Go attempts to reduce the amount of typing in both senses of the word. Throughout its design, we have tried to reduce clutter and complexity. There are no forward declarations and no header files; everything is declared exactly once. Initialization is expressive, automatic, and easy to use. Syntax is clean and light on keywords. Stuttering (foo.Foo* myFoo = new(foo.Foo)) is reduced by simple type derivation using the := declare-and-initialize construct. And perhaps most radically, there is no type hierarchy: types just are, they don't have to announce their relationships. These simplifications allow Go to be expressive yet comprehensible without sacrificing, well, sophistication.

Another important principle is to keep the concepts orthogonal. Methods can be implemented for any type; structures represent data while interfaces represent abstraction; and so on. Orthogonality makes it easier to understand what happens when things combine.

### 设计中的指导原则是什么？
在设计 `Go` 的时候，`Java` 和 `C++` 是编写服务器最常用的语言，至少在 `Google` 是这样。我们觉得这些语言需要太多的记账和重复。一些程序员的反应是转向 `Python` 等更动态、更流畅的语言，但代价是效率和类型安全。我们认为应该可以在一种语言中拥有效率、安全和流动性。

`Go` 试图减少两种意义上的类型量。在整个设计过程中，我们一直试图减少混乱和复杂性。没有前向声明和头文件，所有的东西都只声明一次。初始化是富有表现力的，自动的，易于使用的。语法是干净的，关键词也很少。通过使用 `:=` 声明和初始化结构进行简单的类型推导，减少了啰嗦表达式  `(foo.Foo* myFoo = new(foo.Foo))`。也许最根本的是，没有类型层次结构：类型就是这样，它们不需要公布它们的关系。这些简化使 Go 的表现力和可理解性都得到了提升，同时也没有牺牲复杂度。

另一个重要的原则是保持概念的正交性。方法可以为任何类型实现；结构代表数据而接口代表抽象；等等。正交性使我们更容易理解事物结合后的情况。

## Usage（使用）
### Is Google using Go internally?
Yes. Go is used widely in production inside Google. One easy example is the server behind golang.org. It's just the godoc document server running in a production configuration on Google App Engine.

A more significant instance is Google's download server, dl.google.com, which delivers Chrome binaries and other large installables such as apt-get packages.

Go is not the only language used at Google, far from it, but it is a key language for a number of areas including [site reliability engineering (SRE)](https://talks.golang.org/2013/go-sreops.slide) and large-scale data processing.

### 谷歌内部使用 Go 吗？
是的，谷歌内部在生产中广泛使用 `Go`。`Go` 在 `Google` 内部的生产中被广泛使用。一个简单的例子就是 `golang.org` 背后的服务器。它只是在 `Google App Engine` 上以生产配置运行的 `godoc` 文档服务器。

更重要的例子是 `Google` 的下载服务器 `dl.google.com`，它提供 `Chrome` 二进制文件和其他大型可安装文件，如 `apt-get` 包。

`Go` 并不是 `Google` 唯一使用的语言，远非如此，但它是网站可靠性工程（SRE）和大规模数据处理等多个领域的关键语言。

### What other companies use Go?
Go usage is growing worldwide, especially but by no means exclusively in the cloud computing space. A couple of major cloud infrastructure projects written in Go are Docker and Kubernetes, but there are many more.

It's not just cloud, though. The Go Wiki includes a [page](https://github.com/golang/go/wiki/GoUsers), updated regularly, that lists some of the many companies using Go.

The Wiki also has a page with links to success stories about companies and projects that are using the language.

### 还有哪些公司使用Go？
`Go` 的使用量在全球范围内不断增长，尤其是但绝不是只在云计算领域。用 `Go` 编写的几个主要云基础设施项目是 `Docker` 和 `Kubernetes`，但还有更多。

不过，这不仅仅是云计算。`Go Wiki` 包含一个定期更新的页面，列出了一些使用 `Go` 的公司。

维基也有一个页面，链接到使用该语言的公司和项目的成功故事。

### Do Go programs link with C/C++ programs?
It is possible to use C and Go together in the same address space, but it is not a natural fit and can require special interface software. Also, linking C with Go code gives up the memory safety and stack management properties that Go provides. Sometimes it's absolutely necessary to use C libraries to solve a problem, but doing so always introduces an element of risk not present with pure Go code, so do so with care.

If you do need to use C with Go, how to proceed depends on the Go compiler implementation. There are three Go compiler implementations supported by the Go team. These are gc, the default compiler, gccgo, which uses the GCC back end, and a somewhat less mature gollvm, which uses the LLVM infrastructure.

Gc uses a different calling convention and linker from C and therefore cannot be called directly from C programs, or vice versa. The cgo program provides the mechanism for a “foreign function interface” to allow safe calling of C libraries from Go code. SWIG extends this capability to C++ libraries.

You can also use cgo and SWIG with Gccgo and gollvm. Since they use a traditional API, it's also possible, with great care, to link code from these compilers directly with GCC/LLVM-compiled C or C++ programs. However, doing so safely requires an understanding of the calling conventions for all languages concerned, as well as concern for stack limits when calling C or C++ from Go.

### Go程序可以和C/C++程序链接吗？
在同一个地址空间中，`C` 和 `Go` 是可以一起使用的，但这不是一个自然的配合，可能需要特殊的接口软件。而且，将 `C`与 `Go` 代码链接起来，会放弃 `Go`提供的内存安全和堆栈管理特性。有时候，使用 `C` 库来解决问题是绝对必要的，但这样做总是会引入纯 `Go` 代码所不具备的风险因素，所以要谨慎行事。

如果你确实需要在 `Go` 中使用 `C` 库，如何进行取决于 `Go` 编译器的实现。 `Go` 团队支持的 `Go` 编译器实现有三种。它们是 `gc`，默认的编译器，`gccgo`，使用 `GCC` 后端，以及一个不太成熟的 `gollvm`，使用 `LLVM` 基础设施。

`Gc` 使用了与 `C` 不同的调用约定和链接器，因此不能直接从 `C` 程序中调用，反之亦然。 `cgo` 程序提供了 "外函数接口 "的机制，允许从 `Go` 代码中安全调用 `C` 库。`SWIG` 将这种功能扩展到 `C++` 库。

你也可以将 `cgo` 和 `SWIG` 与 `Gccgo` 和 `gollvm` 一起使用。由于它们使用的是传统的 `API`，所以也可以很小心地将这些编译器的代码直接与 `GCC/LLVM` 编译的 `C` 或 `C++` 程序链接。然而，安全地这样做需要了解所有相关语言的调用惯例，以及从 `Go` 调用 `C` 或 `C++` 时对堆栈限制的关注。

### What IDEs does Go support?
The Go project does not include a custom IDE, but the language and libraries have been designed to make it easy to analyze source code. As a consequence, most well-known editors and IDEs support Go well, either directly or through a plugin.

The list of well-known IDEs and editors that have good Go support available includes Emacs, Vim, VSCode, Atom, Eclipse, Sublime, IntelliJ (through a custom variant called Goland), and many more. Chances are your favorite environment is a productive one for programming in Go.

### Go支持哪些IDE？
`Go` 项目不包括自定义的 `IDE`，但语言和库的设计使其易于分析源代码。因此，大多数著名的编辑器和 `IDE` 都能直接或通过插件很好地支持 `Go`。

知名的 `IDE` 和编辑器列表中都有很好的 `Go` 支持，包括 `Emacs`、`Vim`、`VSCode`、`Atom`、`Eclipse`、`Sublime`、`IntelliJ`（通过一个名为Goland的定制变体），以及更多。您最喜欢的环境很有可能是在 `Go` 中编程的高效环境。

### Does Go support Google's protocol buffers?
A separate open source project provides the necessary compiler plugin and library. It is available at github.com/golang/protobuf/.

### Go是否支持Google的 protocol buffers？
一个单独的开源项目提供了必要的编译器插件和库。它可以在[github.com/golang/protobuf/](https://github.com/golang/protobuf/)获得。

### Can I translate the Go home page into another language?
Absolutely. We encourage developers to make Go Language sites in their own languages. However, if you choose to add the Google logo or branding to your site (it does not appear on golang.org), you will need to abide by the guidelines at www.google.com/permissions/guidelines.html

### 我可以将Go主页翻译成其他语言吗？
当然可以。我们鼓励开发者用自己的语言制作 `Go` 语言网站。但是，如果您选择在您的网站上添加 `Google` 标志或品牌（它不会出现在golang.org上），您将需要遵守www.google.com/permissions/guidelines.html。

## Design（设计）
### Does Go have a runtime?
Go does have an extensive library, called the runtime, that is part of every Go program. The runtime library implements garbage collection, concurrency, stack management, and other critical features of the Go language. Although it is more central to the language, Go's runtime is analogous to libc, the C library.

It is important to understand, however, that Go's runtime does not include a virtual machine, such as is provided by the Java runtime. Go programs are compiled ahead of time to native machine code (or JavaScript or WebAssembly, for some variant implementations). Thus, although the term is often used to describe the virtual environment in which a program runs, in Go the word “runtime” is just the name given to the library providing critical language services.

### Go有运行时吗？
`Go` 确实有一个扩展库叫做 *runtime* （运行时），它是每个 `Go` 程序的一部分。运行时库实现了垃圾收集、并发、堆栈管理和其他 `Go` 语言的关键功能。虽然它对 `Go` 语言更为核心，但 `Go` 的运行时类似于 `C` 库 `libc`。

然而，重要的是要明白，`Go` 的运行时不包括虚拟机，如 `Java` 运行时提供的虚拟机。`Go` 程序是提前编译成本地机器代码（或JavaScript或WebAssembly，对于某些变体实现）。因此，虽然这个词经常被用来描述程序运行的虚拟环境，但在 `Go` 中，"运行时 "这个词只是提供关键语言服务的库的名称。

### What's up with Unicode identifiers?
When designing Go, we wanted to make sure that it was not overly ASCII-centric, which meant extending the space of identifiers from the confines of 7-bit ASCII. Go's rule—identifier characters must be letters or digits as defined by Unicode—is simple to understand and to implement but has restrictions. Combining characters are excluded by design, for instance, and that excludes some languages such as Devanagari.

This rule has one other unfortunate consequence. Since an exported identifier must begin with an upper-case letter, identifiers created from characters in some languages can, by definition, not be exported. For now the only solution is to use something like X日本語, which is clearly unsatisfactory.

Since the earliest version of the language, there has been considerable thought into how best to expand the identifier space to accommodate programmers using other native languages. Exactly what to do remains an active topic of discussion, and a future version of the language may be more liberal in its definition of an identifier. For instance, it might adopt some of the ideas from the Unicode organization's recommendations for identifiers. Whatever happens, it must be done compatibly while preserving (or perhaps expanding) the way letter case determines visibility of identifiers, which remains one of our favorite features of Go.

For the time being, we have a simple rule that can be expanded later without breaking programs, one that avoids bugs that would surely arise from a rule that admits ambiguous identifiers.

### Unicode标识符是怎么回事？
当设计 `Go` 时，我们想确保它不会过度以 `ASCII` 为中心，这意味着将标识符的空间从7位 `ASCII` 的限制中扩展出来。`Go` 的规则--标识符字符必须是 `Unicod`定义的字母或数字--很容易理解和实现，但也有限制。例如，组合字符被设计排除在外，这就排除了一些语言，如 `Devanagari`。

这个规则还有一个不幸的后果。由于导出的标识符必须以大写字母开头，所以根据定义，由某些语言中的字符创建的标识符不能导出。目前唯一的解决办法是使用类似X日本语这样的语言，但这显然不能令人满意。

从最早的语言版本开始，人们就在考虑如何最好地扩展标识符空间，以适应使用其他母语的程序员。具体怎么做仍然是一个活跃的讨论话题，未来的语言版本可能会在标识符的定义上更加自由。例如，它可能会采用 `Unicode` 组织对标识符建议中的一些想法。无论发生什么，都必须兼容，同时保留（或扩大）字母大小写决定标识符可见性的方式，这仍然是我们最喜欢的围棋特性之一。

目前，我们有一个简单的规则，可以在不破坏程序的情况下进行扩展，这个规则可以避免一个接纳模糊标识符的规则肯定会产生的错误。

### Why does Go not have feature X?
Every language contains novel features and omits someone's favorite feature. Go was designed with an eye on felicity of programming, speed of compilation, orthogonality of concepts, and the need to support features such as concurrency and garbage collection. Your favorite feature may be missing because it doesn't fit, because it affects compilation speed or clarity of design, or because it would make the fundamental system model too difficult.

If it bothers you that Go is missing feature X, please forgive us and investigate the features that Go does have. You might find that they compensate in interesting ways for the lack of X.

### 为什么Go没有特征X？
每一种语言都包含新颖的功能，也会省略某人最喜欢的功能。`Go` 的设计着眼于编程的简洁性、编译速度、概念的正交性，以及支持并发和垃圾回收等特性的需要。你最喜欢的特性可能会因为不适合而缺失，因为它影响编译速度或设计的清晰度，或者因为它会使基本的系统模型变得太难。

如果你因为 `Go` 缺少了特性X而感到困扰，请原谅我们，调查一下 `Go` 确实有的特性。您可能会发现，它们以有趣的方式弥补了X的缺失。

### Why does Go not have generic types?
Generics may well be added at some point. We don't feel an urgency for them, although we understand some programmers do.

Go was intended as a language for writing server programs that would be easy to maintain over time. (See [this article](https://talks.golang.org/2012/splash.article) for more background.) The design concentrated on things like scalability, readability, and concurrency. Polymorphic programming did not seem essential to the language's goals at the time, and so was left out for simplicity.

The language is more mature now, and there is scope to consider some form of generic programming. However, there remain some caveats.

Generics are convenient but they come at a cost in complexity in the type system and run-time. We haven't yet found a design that gives value proportionate to the complexity, although we continue to think about it. Meanwhile, Go's built-in maps and slices, plus the ability to use the empty interface to construct containers (with explicit unboxing) mean in many cases it is possible to write code that does what generics would enable, if less smoothly.

The topic remains open. For a look at several previous unsuccessful attempts to design a good generics solution for Go, see this proposal.

### 为什么Go没有泛型？
泛型很可能会在某些时候被添加。我们并不觉得它们的迫切性，尽管我们知道有些程序员会这么做。

`Go` 的目的是作为一种用于编写服务器程序的语言，这种语言很容易长期维护。(更多的背景请看这篇文章。)设计集中在可扩展性、可读性和并发性等方面。多态编程在当时看来对该语言的目标并不重要，所以为了简单起见，就没有采用。

现在，该语言已经比较成熟了，而且还有考虑某种形式的通用编程的余地。但是，仍然有一些注意事项。

泛型是方便的，但是在类型系统和运行时的复杂性上是有代价的。我们还没有找到一种能让价值与复杂度成正比的设计，尽管我们还在继续思考。同时，`Go` 内置的字典和切片，再加上使用空接口构造容器的能力（有显式的开箱），意味着在许多情况下，可以写出能实现通用性的代码，即使不那么流畅。

这个话题仍然是开放的。要想了解之前为 `Go` 设计一个好的 `generics` 解决方案的几个不成功的尝试，请看这个建议。

### Why does Go not have exceptions?
We believe that coupling exceptions to a control structure, as in the try-catch-finally idiom, results in convoluted code. It also tends to encourage programmers to label too many ordinary errors, such as failing to open a file, as exceptional.

Go takes a different approach. For plain error handling, Go's multi-value returns make it easy to report an error without overloading the return value. A canonical error type, coupled with Go's other features, makes error handling pleasant but quite different from that in other languages.

Go also has a couple of built-in functions to signal and recover from truly exceptional conditions. The recovery mechanism is executed only as part of a function's state being torn down after an error, which is sufficient to handle catastrophe but requires no extra control structures and, when used well, can result in clean error-handling code.

See the [Defer, Panic, and Recover](https://golang.org/doc/articles/defer_panic_recover.html) article for details. Also, the [Errors are values](https://blog.golang.org/errors-are-values) blog post describes one approach to handling errors cleanly in Go by demonstrating that, since errors are just values, the full power of Go can be deployed in error handling.

### 为什么Go没有异常？
我们认为，将异常与控制结构耦合在一起，就像 `try-catch-finally` 习语一样，会导致代码复杂化。它还倾向于鼓励程序员将太多的普通错误（如未能打开文件）标记为异常。

`Go` 采用了一种不同的方法。对于普通的错误处理，`Go` 的多值返回使报告错误变得容易，而无需重载返回值。一个规范的错误类型，加上 `Go` 的其他特性，使得错误处理变得愉快，但与其他语言的错误处理完全不同。

`Go` 还有几个内置函数，用于从真正的特殊情况下发出信号和恢复。恢复机制只在出错后作为函数的状态被拆掉的一部分来执行，这足以处理灾难，但不需要额外的控制结构，用得好的话，可以得到干净的错误处理代码。

详见[ `Defer`、`Panic` 和 `Recover`](https://golang.org/doc/articles/defer_panic_recover.html) 文章。另外，`Errors are values` 这篇博文描述了一种在 `Go` 中干净地处理错误的方法，它展示了由于错误只是值，所以可以在错误处理中部署Go的全部能力。

### Why does Go not have assertions?
Go doesn't provide assertions. They are undeniably convenient, but our experience has been that programmers use them as a crutch to avoid thinking about proper error handling and reporting. Proper error handling means that servers continue to operate instead of crashing after a non-fatal error. Proper error reporting means that errors are direct and to the point, saving the programmer from interpreting a large crash trace. Precise errors are particularly important when the programmer seeing the errors is not familiar with the code.

We understand that this is a point of contention. There are many things in the Go language and libraries that differ from modern practices, simply because we feel it's sometimes worth trying a different approach.

### 为什么Go没有断言？
Go没有提供断言。不可否认，它们很方便，但我们的经验是，程序员把它们当作一个拐杖，以避免考虑适当的错误处理和报告。正确的错误处理意味着服务器继续运行，而不是在发生非致命错误后崩溃。正确的错误报告意味着错误是直接的和重点的，使程序员免于解释一个大的崩溃跟踪。当看到错误的程序员不熟悉代码时，精确的错误尤为重要。

我们理解这是一个争论点。围棋语言和库中有很多与现代实践不同的地方，只是因为我们觉得有时值得尝试不同的方法。

### Why build concurrency on the ideas of CSP?
Concurrency and multi-threaded programming have over time developed a reputation for difficulty. We believe this is due partly to complex designs such as pthreads and partly to overemphasis on low-level details such as mutexes, condition variables, and memory barriers. Higher-level interfaces enable much simpler code, even if there are still mutexes and such under the covers.

One of the most successful models for providing high-level linguistic support for concurrency comes from Hoare's Communicating Sequential Processes, or CSP. Occam and Erlang are two well known languages that stem from CSP. Go's concurrency primitives derive from a different part of the family tree whose main contribution is the powerful notion of channels as first class objects. Experience with several earlier languages has shown that the CSP model fits well into a procedural language framework.

### 为什么要把并发建立在CSP的思想上？
并发和多线程编程随着时间的推移已经形成了困难的声誉。我们认为这部分是由于复杂的设计，如 `pthreads`，部分是由于过度强调低级细节，如 `mutexes`，条件变量和内存屏障。更高层次的接口可以使代码简单得多，即使在掩盖之下仍然有 `mutexes` 之类的东西。

为并发性提供高级语言支持的最成功的模型之一来自 `Hoare` 的 Communicating Sequential Processes`，即 `CSP`。`Occam` 和 `Erlang` 是源于 `CSP` 的两种著名语言。`Go` 的并发基元来源于家族树的另一部分，其主要贡献是通道作为一级对象的强大概念。几个早期语言的经验表明， `CSP` 模型很适合程序化语言框架。

### Why goroutines instead of threads?
Goroutines are part of making concurrency easy to use. The idea, which has been around for a while, is to multiplex independently executing functions—coroutines—onto a set of threads. When a coroutine blocks, such as by calling a blocking system call, the run-time automatically moves other coroutines on the same operating system thread to a different, runnable thread so they won't be blocked. The programmer sees none of this, which is the point. The result, which we call goroutines, can be very cheap: they have little overhead beyond the memory for the stack, which is just a few kilobytes.

To make the stacks small, Go's run-time uses resizable, bounded stacks. A newly minted goroutine is given a few kilobytes, which is almost always enough. When it isn't, the run-time grows (and shrinks) the memory for storing the stack automatically, allowing many goroutines to live in a modest amount of memory. The CPU overhead averages about three cheap instructions per function call. It is practical to create hundreds of thousands of goroutines in the same address space. If goroutines were just threads, system resources would run out at a much smaller number.

### 为什么用goroutines而不是线程？
`Goroutines` 是使并发变得简单易用的一部分。这个想法已经存在了一段时间，就是将独立执行的函数-- `goroutine`--复用到一组线程中。当一个 `coroutine` 阻塞时，比如调用阻塞系统调用，运行时就会自动将同一操作系统线程上的其他 `coroutine` 移动到另一个可运行的线程上，这样它们就不会被阻塞。程序员看不到这些，这就是问题的关键。其结果，我们称之为 `Goroutines`，可以非常便宜：它们除了堆栈的内存外，几乎没有什么开销，而堆栈的内存只有几千字节。

为了使堆栈变小，`Go` 的运行时使用了可调整大小、有边界的堆栈。一个新诞生的 `goroutine` 被赋予了几千字节的内存，这几乎总是足够的。当它不够时，运行时就会自动增长（和收缩）用于存储堆栈的内存，允许许多 `goroutine` 生活在一个适度的内存中。 `CPU` 的开销平均为每次函数调用约三条廉价指令。在同一个地址空间中创建几十万个 `goroutine` 是很实用的。如果 `goroutine` 只是线程，系统资源的耗费将大大减少。

### Why are map operations not defined to be atomic?
After long discussion it was decided that the typical use of maps did not require safe access from multiple goroutines, and in those cases where it did, the map was probably part of some larger data structure or computation that was already synchronized. Therefore requiring that all map operations grab a mutex would slow down most programs and add safety to few. This was not an easy decision, however, since it means uncontrolled map access can crash the program.

The language does not preclude atomic map updates. When required, such as when hosting an untrusted program, the implementation could interlock map access.

Map access is unsafe only when updates are occurring. As long as all goroutines are only reading—looking up elements in the map, including iterating through it using a for range loop—and not changing the map by assigning to elements or doing deletions, it is safe for them to access the map concurrently without synchronization.

As an aid to correct map use, some implementations of the language contain a special check that automatically reports at run time when a map is modified unsafely by concurrent execution.

### 为什么不将map操作定义为原子操作？
经过长时间的讨论，我们决定地图的典型使用不需要多个 `goroutine` 的安全访问，在那些需要安全访问的情况下，地图很可能是一些更大的数据结构的一部分，或者是已经同步的计算。因此，要求所有的 `map` 操作都要抓取一个 `mutex`，会减慢大多数程序的速度，而增加少数程序的安全性。然而，这并不是一个容易的决定，因为这意味着不受控制的映射访问会使程序崩溃。

该语言不排除原子图更新。当需要的时候，比如托管一个不受信任的程序时，实现可以联锁地图访问。

只有当更新发生时，地图访问才是不安全的。只要所有的 `goroutine` 都只是读取地图中的元素，包括使用 `for range` 循环进行迭代，而不是通过分配给元素或进行删除来改变地图，那么它们在没有同步的情况下并发访问地图是安全的。

作为对正确使用地图的帮助，一些语言的实现包含了一个特殊的检查，当一个地图被并发执行的不安全修改时，它会在运行时自动报告。

### Will you accept my language change?
People often suggest improvements to the language—the mailing list contains a rich history of such discussions—but very few of these changes have been accepted.

Although Go is an open source project, the language and libraries are protected by a compatibility promise that prevents changes that break existing programs, at least at the source code level (programs may need to be recompiled occasionally to stay current). If your proposal violates the Go 1 specification we cannot even entertain the idea, regardless of its merit. A future major release of Go may be incompatible with Go 1, but discussions on that topic have only just begun and one thing is certain: there will be very few such incompatibilities introduced in the process. Moreover, the compatibility promise encourages us to provide an automatic path forward for old programs to adapt should that situation arise.

Even if your proposal is compatible with the Go 1 spec, it might not be in the spirit of Go's design goals. The article Go at Google: Language Design in the Service of Software Engineering explains Go's origins and the motivation behind its design.

### 你会接受我对语言的修改吗？
人们经常建议对语言进行改进--邮件列表中包含了丰富的此类讨论历史--但很少有这些修改被接受。

虽然围棋是一个开源项目，但语言和库受到兼容性承诺的保护，至少在源代码层面上防止了破坏现有程序的更改（程序可能需要偶尔重新编译以保持最新性）。如果你的建议违反了 `Go 1` 的规范，我们甚至不能接受这个想法，无论它的优点如何。未来的围棋主要版本可能会与 `Go 1` 不兼容，但关于这个话题的讨论才刚刚开始，有一点是肯定的：在这个过程中引入的这种不兼容的情况将非常少。此外，兼容性承诺鼓励我们为旧程序提供一条自动前进的道路，以便在出现这种情况时进行调整。

即使你的建议与 `Go 1` 规范兼容，也可能不符合Go的设计目标的精神。 `Go` 在 `Google` 的文章。语言设计为软件工程服务》一文解释了 `Go` 的起源及其设计背后的动机。

## Types（类型）
### Is Go an object-oriented language?
Yes and no. Although Go has types and methods and allows an object-oriented style of programming, there is no type hierarchy. The concept of “interface” in Go provides a different approach that we believe is easy to use and in some ways more general. There are also ways to embed types in other types to provide something analogous—but not identical—to subclassing. Moreover, methods in Go are more general than in C++ or Java: they can be defined for any sort of data, even built-in types such as plain, “unboxed” integers. They are not restricted to structs (classes).

Also, the lack of a type hierarchy makes “objects” in Go feel much more lightweight than in languages such as C++ or Java.

### Go是一种面向对象的语言吗？
是，也不是。虽然 `Go` 有类型和方法，并允许面向对象的编程风格，但没有类型层次结构。`Go` 中的 "接口 "概念提供了一种不同的方法，我们相信这种方法很容易使用，并且在某些方面更加通用。还有一些方法可以将类型嵌入到其他类型中，以提供类似于子类的东西--但不完全相同。此外，`Go` 中的方法比 `C++` 或 `Java` 中的方法更通用：它们可以为任何类型的数据定义，甚至是内置类型，如普通的 "未装箱的 "整数。它们不限于结构（类）。

另外，由于没有类型层次结构，`Go` 中的 "对象 "比 `C++` 或 `Java` 等语言更轻量级。

### How do I get dynamic dispatch of methods?
The only way to have dynamically dispatched methods is through an interface. Methods on a struct or any other concrete type are always resolved statically.

### 如何获得方法的动态调度？
拥有动态调度方法的唯一方法是通过一个接口。结构体或任何其他具体类型的方法总是静态解析的。

### Why is there no type inheritance?
Object-oriented programming, at least in the best-known languages, involves too much discussion of the relationships between types, relationships that often could be derived automatically. Go takes a different approach.

Rather than requiring the programmer to declare ahead of time that two types are related, in Go a type automatically satisfies any interface that specifies a subset of its methods. Besides reducing the bookkeeping, this approach has real advantages. Types can satisfy many interfaces at once, without the complexities of traditional multiple inheritance. Interfaces can be very lightweight—an interface with one or even zero methods can express a useful concept. Interfaces can be added after the fact if a new idea comes along or for testing—without annotating the original types. Because there are no explicit relationships between types and interfaces, there is no type hierarchy to manage or discuss.

It's possible to use these ideas to construct something analogous to type-safe Unix pipes. For instance, see how fmt.Fprintf enables formatted printing to any output, not just a file, or how the bufio package can be completely separate from file I/O, or how the image packages generate compressed image files. All these ideas stem from a single interface (io.Writer) representing a single method (Write). And that's only scratching the surface. Go's interfaces have a profound influence on how programs are structured.

It takes some getting used to but this implicit style of type dependency is one of the most productive things about Go.

### 为什么没有类型继承？
面向对象的编程，至少在最著名的语言中，涉及到太多关于类型之间关系的讨论，而这些关系通常是可以自动导出的。`Go` 采用了一种不同的方法。

在 `Go` 中，一个类型自动满足任何指定其方法子集的接口，而不是要求程序员提前声明两个类型是相关的。除了减少记账，这种方法还有真正的优势。类型可以同时满足许多接口，而没有传统的多重继承的复杂性。接口可以是非常轻量级的--一个只有一个甚至零个方法的接口可以表达一个有用的概念。如果有了新的想法，或者为了测试，可以在事后添加接口--而无需对原始类型进行注释。因为类型和接口之间没有明确的关系，所以没有类型层次结构需要管理或讨论。

我们可以使用这些想法来构建类似于类型安全的 `Unix` 管道。例如，看看 `fmt.Fprintf` 如何实现格式化打印到任何输出，而不仅仅是一个文件，或者 `bufio` 包如何与文件 `I/O` 完全分离，或者 `image` 包如何生成压缩的图像文件。所有这些想法都源于一个接口（io.Writer）代表一个方法（Write）。而这仅仅是表面现象。Go的接口对程序的结构方式有着深远的影响。

它需要一些时间来适应，但这种隐式的类型依赖风格是 `Go` 最有成效的地方之一。

### Why is len a function and not a method?
We debated this issue but decided implementing len and friends as functions was fine in practice and didn't complicate questions about the interface (in the Go type sense) of basic types.

### 为什么len是一个函数而不是方法？
我们对这个问题进行了讨论，但决定将 `len` 和 `friends` 实现为函数在实践中是很好的，而且不会使基本类型的接口问题复杂化（在 `Go` 类型意义上）。


### Why does Go not support overloading of methods and operators?
Method dispatch is simplified if it doesn't need to do type matching as well. Experience with other languages told us that having a variety of methods with the same name but different signatures was occasionally useful but that it could also be confusing and fragile in practice. Matching only by name and requiring consistency in the types was a major simplifying decision in Go's type system.

Regarding operator overloading, it seems more a convenience than an absolute requirement. Again, things are simpler without it.

### 为什么Go不支持方法和操作符的重载？
如果不需要也做类型匹配的话，方法调度就简化了。其他语言的经验告诉我们，拥有多种名称相同但签名不同的方法偶尔是有用的，但在实践中也可能是混乱和脆弱的。只通过名称进行匹配，并要求类型的一致性，是 `Go` 类型系统中一个重要的简化决定。

关于操作符重载，它似乎更多的是一种方便，而不是绝对的要求。同样，没有它，事情就简单了。

### Why doesn't Go have "implements" declarations?
A Go type satisfies an interface by implementing the methods of that interface, nothing more. This property allows interfaces to be defined and used without needing to modify existing code. It enables a kind of structural typing that promotes separation of concerns and improves code re-use, and makes it easier to build on patterns that emerge as the code develops. The semantics of interfaces is one of the main reasons for Go's nimble, lightweight feel.

See the question on type inheritance for more detail.

### 为什么Go没有 "实现 "声明？
一个 `Go` 类型通过实现该接口的方法来满足一个接口，仅此而已。这个属性允许在不需要修改现有代码的情况下定义和使用接口。它实现了一种结构性的类型化，促进了关注点的分离，提高了代码的重用性，也使得在代码开发过程中出现的模式更容易建立。接口的语义是 `Go` 灵活、轻量级感觉的主要原因之一。

更多细节请参见关于类型继承的问题。

### How can I guarantee my type satisfies an interface?
You can ask the compiler to check that the type T implements the interface I by attempting an assignment using the zero value for T or pointer to T, as appropriate:

```
type T struct{}
var _ I = T{}       // Verify that T implements I.
var _ I = (*T)(nil) // Verify that *T implements I.
```

If T (or *T, accordingly) doesn't implement I, the mistake will be caught at compile time.

If you wish the users of an interface to explicitly declare that they implement it, you can add a method with a descriptive name to the interface's method set. For example:

```
type Fooer interface {
    Foo()
    ImplementsFooer()
}
```

A type must then implement the ImplementsFooer method to be a Fooer, clearly documenting the fact and announcing it in go doc's output.

```
type Bar struct{}
func (b Bar) ImplementsFooer() {}
func (b Bar) Foo() {}
```

Most code doesn't make use of such constraints, since they limit the utility of the interface idea. Sometimes, though, they're necessary to resolve ambiguities among similar interfaces.

### 我如何保证我的类型满足接口？
你可以要求编译器检查类型T是否实现了接口I，方法是尝试使用T的零值或T的指针进行赋值，视情况而定。

```
type T struct{}
var _ I = T{} // 验证T是否实现了I。
var _ I = (*T)(nil) // 验证*T是否实现了I。
```

如果T（或*T，相应地）没有实现I，那么在编译时就会发现这个错误。

如果你希望一个接口的用户明确声明他们实现了它，你可以在接口的方法集中添加一个带有描述性名称的方法。例如

```
type Fooer interface {
    Foo()
    ImplementsFooer()
}
```

然后，一个类型必须实现 `ImplementsFooer` 方法才能成为一个 `Fooer`，清楚地记录这一事实，并在 `go doc` 的输出中宣布它。

```
type Bar struct{}
func (b Bar) ImplementsFooer() {}
func (b Bar) Foo() {}
```

大多数代码都不会使用这种约束，因为它们限制了接口概念的效用。但有时，它们对于解决相似接口之间的歧义是必要的。


### Why doesn't type T satisfy the Equal interface?
Consider this simple interface to represent an object that can compare itself with another value:

```
type Equaler interface {
    Equal(Equaler) bool
}
```

and this type, T:

```
type T int
func (t T) Equal(u T) bool { return t == u } // does not satisfy Equaler
```

Unlike the analogous situation in some polymorphic type systems, T does not implement Equaler. The argument type of T.Equal is T, not literally the required type Equaler.

In Go, the type system does not promote the argument of Equal; that is the programmer's responsibility, as illustrated by the type T2, which does implement Equaler:

```
type T2 int
func (t T2) Equal(u Equaler) bool { return t == u.(T2) }  // satisfies Equaler
```

Even this isn't like other type systems, though, because in Go any type that satisfies Equaler could be passed as the argument to T2.Equal, and at run time we must check that the argument is of type T2. Some languages arrange to make that guarantee at compile time.

A related example goes the other way:

```
type Opener interface {
   Open() Reader
}

func (t T3) Open() *os.File
```

In Go, T3 does not satisfy Opener, although it might in another language.

While it is true that Go's type system does less for the programmer in such cases, the lack of subtyping makes the rules about interface satisfaction very easy to state: are the function's names and signatures exactly those of the interface? Go's rule is also easy to implement efficiently. We feel these benefits offset the lack of automatic type promotion. Should Go one day adopt some form of polymorphic typing, we expect there would be a way to express the idea of these examples and also have them be statically checked.

### 为什么T类型不满足Equal接口？
考虑用这个简单的接口来表示一个可以与另一个值进行比较的对象。

```
type Equaler interface {
    Equal(Equaler) bool
}
```

和这个类型，T。

```
type T int
func (t T) Equal(u T) bool { return t == u } // 不满足 Equaler 的要求。
```

与一些多态类型系统的类似情况不同，T并没有实现Equaler。T.Equal的参数类型是T，而不是字面上所要求的Equaler类型。

在 `Go` 中，类型系统并不提倡 `Equal` 的参数，那是程序员的责任，如类型T2所示，它确实实现了Equaler。

```
type T2 int
func (t T2) Equal(u Equaler) bool { return t == u.(T2) }  // 满足Equaler
```

不过这也和其他类型系统不一样，因为在Go中，任何满足Equaler的类型都可以作为T2.Equal的参数传递给T2.Equal，而在运行时我们必须检查参数是否为T2类型。有些语言安排在编译时做出这种保证。

一个相关的例子则反其道而行之。

```
type Opener interface {
   Open() Reader
}

func (t T3) Open() *os.File
```

在Go中，T3不满足Opener的要求，虽然在其他语言中可能会满足。

虽然在这种情况下，Go的类型系统确实对程序员的作用较小，但由于没有子类型化，关于接口满足的规则非常容易说明：函数的名称和签名是否与接口的名称和签名完全一致？Go的规则也很容易有效实现。我们觉得这些好处可以抵消自动类型推广的不足。如果有一天Go采用了某种形式的多态类型，我们期望有一种方法来表达这些例子的想法，并且还能对它们进行静态检查。

### Can I convert a []T to an []interface{}?
Not directly. It is disallowed by the language specification because the two types do not have the same representation in memory. It is necessary to copy the elements individually to the destination slice. This example converts a slice of int to a slice of interface{}:

```
t := []int{1, 2, 3, 4}
s := make([]interface{}, len(t))
for i, v := range t {
    s[i] = v
}
```



### #我可以将[]T转换为[]interface{}吗？
不能直接转换。语言规范不允许这样做，因为这两种类型在内存中没有相同的表示。必须将元素分别复制到目标分片中。这个例子将 `int` 的分片转换为 `interface{}` 的分片。

```
t := []int{1, 2, 3, 4}。
s := make([]interface{}, len(t))
for i, v := range t {
    s[i] = v
}
```



### Can I convert []T1 to []T2 if T1 and T2 have the same underlying type?This last line of this code sample does not compile.
```
type T1 int
type T2 int
var t1 T1
var x = T2(t1) // OK
var st1 []T1
var sx = ([]T2)(st1) // NOT OK
```

In Go, types are closely tied to methods, in that every named type has a (possibly empty) method set. The general rule is that you can change the name of the type being converted (and thus possibly change its method set) but you can't change the name (and method set) of elements of a composite type. Go requires you to be explicit about type conversions.

### 我是否可以将 [] T1 转换为 [] T2 如果 T1 和 T2 有相同的基础类型？

```
type T1 int
type T2 int
var t1 T1
var x = T2(t1) // OK
var st1 []T1
var sx = ([]T2)(st1) // NOT OK
```

在 `Go` 中，类型与方法密切相关，因为每个命名的类型都有一个（可能是空的）方法集。一般的规则是，你可以改变被转换的类型的名称（从而可能改变它的方法集），但你不能改变复合类型元素的名称（和方法集）。`Go` 要求你明确类型转换。


### Why is my nil error value not equal to nil?
Under the covers, interfaces are implemented as two elements, a type T and a value V. V is a concrete value such as an int, struct or pointer, never an interface itself, and has type T. For instance, if we store the int value 3 in an interface, the resulting interface value has, schematically, (T=int, V=3). The value V is also known as the interface's dynamic value, since a given interface variable might hold different values V (and corresponding types T) during the execution of the program.

An interface value is nil only if the V and T are both unset, (T=nil, V is not set), In particular, a nil interface will always hold a nil type. If we store a nil pointer of type *int inside an interface value, the inner type will be *int regardless of the value of the pointer: (T=*int, V=nil). Such an interface value will therefore be non-nil even when the pointer value V inside is nil.

This situation can be confusing, and arises when a nil value is stored inside an interface value such as an error return:

```
func returnsError() error {
	var p *MyError = nil
	if bad() {
		p = ErrBad
	}
	return p // Will always return a non-nil error.
}
```

If all goes well, the function returns a nil p, so the return value is an error interface value holding (T=*MyError, V=nil). This means that if the caller compares the returned error to nil, it will always look as if there was an error even if nothing bad happened. To return a proper nil error to the caller, the function must return an explicit nil:

```
func returnsError() error {
	if bad() {
		return ErrBad
	}
	return nil
}
```

It's a good idea for functions that return errors always to use the error type in their signature (as we did above) rather than a concrete type such as *MyError, to help guarantee the error is created correctly. As an example, os.Open returns an error even though, if not nil, it's always of concrete type *os.PathError.

Similar situations to those described here can arise whenever interfaces are used. Just keep in mind that if any concrete value has been stored in the interface, the interface will not be nil. For more information, see The Laws of Reflection.

### 为什么我的nil错误值不等于nil？
在底层，接口实现为两个元素，一个是类型T，一个是值V，V是一个具体的值，比如int、struct或指针，绝不是接口本身，其类型为T。例如，如果我们将int值3存储在一个接口中，那么所产生的接口值的示意为：（T=int，V=3）。V值也被称为接口的动态值，因为在程序执行过程中，一个给定的接口变量可能会持有不同的值V（以及相应的类型T）。

只有当V和T都未设置时，一个接口值才是 `nil`，（T=nil，V未设置），特别是，一个 `nil` 接口将始终持有一个 `nil` 类型。如果我们在一个接口值里面存储一个类型为 `*int` 的 `nil` 指针，无论指针的值是多少，其内部类型都将是 `*int`。(T=*int, V=nil). 因此，即使里面的指针值V是 `nil`，这样的接口值也将是非 `nil`。

这种情况可能会让人感到困惑，当一个 `nil` 值存储在一个接口值里面时，比如错误返回，就会出现这种情况。

```
func returnsError() error {
    var p *MyError = nil
    if bad() {
        p = ErrBad
    }
    return p // 总是返回一个非空类型的错误
}
```

如果一切顺利，函数返回一个 `nil p`，所以返回值是一个错误接口值保持 `（T=*MyError，V=nil）`。这意味着，如果调用者将返回的错误与nil进行比较，即使没有发生任何坏事，它也总是看起来像有一个错误。为了向调用者返回一个合适的nil错误，函数必须返回一个显式的nil。

````
func returnsError() error {
    if bad() {
        return ErrBad
    }
    return nil
}
````

对于那些返回错误的函数来说，最好在其签名中使用错误类型(就像我们上面所做的那样)，而不是像*MyError这样的具体类型，以帮助保证错误的正确创建。举个例子，`os.Open` 返回一个错误，即使它不是 `nil` ，也总是使用具体类型  ` *os.PathError` 。

只要使用了接口，就会出现与这里描述的类似情况。只要记住，如果在接口中存储了任何具体的值，接口就不会是nil。更多信息，请参见《反射法则》。

### Why are there no untagged unions, as in C?
Untagged unions would violate Go's memory safety guarantees.

### 为什么没有像C语言那样的无标记联合？
未标记的联合体会违反 `Go` 的内存安全保证。

### Why does Go not have variant types?
Variant types, also known as algebraic types, provide a way to specify that a value might take one of a set of other types, but only those types. A common example in systems programming would specify that an error is, say, a network error, a security error or an application error and allow the caller to discriminate the source of the problem by examining the type of the error. Another example is a syntax tree in which each node can be a different type: declaration, statement, assignment and so on.

We considered adding variant types to Go, but after discussion decided to leave them out because they overlap in confusing ways with interfaces. What would happen if the elements of a variant type were themselves interfaces?

Also, some of what variant types address is already covered by the language. The error example is easy to express using an interface value to hold the error and a type switch to discriminate cases. The syntax tree example is also doable, although not as elegantly.

### 为什么 Go 没有变体类型？
变量类型，也被称为代数类型，它提供了一种方法来指定一个值可以取一组其他类型中的一个，但只能取那些类型。在系统编程中，一个常见的例子是指定一个错误是比如网络错误、安全错误或应用程序错误，并允许调用者通过检查错误的类型来分辨问题的来源。另一个例子是语法树，其中每个节点可以是不同的类型：声明、声明、赋值等等。

我们考虑过在 `Go` 中加入变体类型，但经过讨论后决定不加入，因为它们与接口有重叠的地方。如果变体类型的元素本身就是接口，会发生什么？

另外，变体类型所解决的一些问题已经被语言所覆盖。错误的例子很容易表达，用一个接口值来存放错误，用一个类型开关来区分情况。语法树的例子也是可以做的，虽然没有那么优雅。

### Why does Go not have covariant result types?
Covariant result types would mean that an interface like

```
type Copyable interface {
	Copy() interface{}
}
```

would be satisfied by the method

```
func (v Value) Copy() Value
```

because Value implements the empty interface. In Go method types must match exactly, so Value does not implement Copyable. Go separates the notion of what a type does—its methods—from the type's implementation. If two methods return different types, they are not doing the same thing. Programmers who want covariant result types are often trying to express a type hierarchy through interfaces. In Go it's more natural to have a clean separation between interface and implementation.

### 为什么Go没有共变结果类型？
共变的结果类型意味着一个接口，如

```
type Copyable interface {
	Copy() interface{}
}
```

你对这个方法满意吗

```
func (v Value) Copy() Value
```

因为 `Value` 实现了空接口。在 `Go` 中，方法类型必须完全匹配，所以 `Value` 没有实现 `Copyable` 。`Go` 将一个类型做什么的概念--它的方法--与类型的实现分开。如果两个方法返回不同的类型，它们做的不是同一件事。想要共变结果类型的程序员通常会试图通过接口来表达类型层次。在 `Go` 中，接口和实现之间的分离是比较自然的。

## Values（值）
### Why does Go not provide implicit numeric conversions?
The convenience of automatic conversion between numeric types in C is outweighed by the confusion it causes. When is an expression unsigned? How big is the value? Does it overflow? Is the result portable, independent of the machine on which it executes? It also complicates the compiler; “the usual arithmetic conversions” are not easy to implement and inconsistent across architectures. For reasons of portability, we decided to make things clear and straightforward at the cost of some explicit conversions in the code. The definition of constants in Go—arbitrary precision values free of signedness and size annotations—ameliorates matters considerably, though.

A related detail is that, unlike in C, int and int64 are distinct types even if int is a 64-bit type. The int type is generic; if you care about how many bits an integer holds, Go encourages you to be explicit.

### 为什么Go不提供隐式数值转换？
在 `C` 语言中，数值类型之间自动转换的便利性被它所造成的混乱所抵消。什么时候表达式是无符号的？值有多大？它是否会溢出？结果是否可移植，与执行的机器无关？它还使编译器复杂化；"通常的算术转换 "不容易实现，而且在不同的架构上不一致。出于可移植性的考虑，我们决定以在代码中进行一些显式转换为代价，让事情变得清晰明了。不过，`Go` 中的常量定义--不含符号和大小注释的任意精度值--大大改善了问题。

一个相关的细节是，与 `C` 中不同，即使 `int` 是 `64` 位类型，`int` 和 `int64` 也是不同的类型。`int` 类型是通用的；如果你关心一个整数拥有多少位，`Go` 鼓励你使用显式。

### How do constants work in Go?
Although Go is strict about conversion between variables of different numeric types, constants in the language are much more flexible. Literal constants such as 23, 3.14159 and [math.Pi](https://golang.org/pkg/math/#pkg-constants) occupy a sort of ideal number space, with arbitrary precision and no overflow or underflow. For instance, the value of math.Pi is specified to 63 places in the source code, and constant expressions involving the value keep precision beyond what a float64 could hold. Only when the constant or constant expression is assigned to a variable—a memory location in the program—does it become a "computer" number with the usual floating-point properties and precision.

Also, because they are just numbers, not typed values, constants in Go can be used more freely than variables, thereby softening some of the awkwardness around the strict conversion rules. One can write expressions such as

```
sqrt2 := math.Sqrt(2)
```

without complaint from the compiler because the ideal number 2 can be converted safely and accurately to a float64 for the call to math.Sqrt.

A blog post titled [Constants](https://blog.golang.org/constants) explores this topic in more detail.

### 常量在Go中是如何工作的？
虽然 `Go` 对不同数值类型的变量之间的转换很严格，但语言中的常量却要灵活得多。字面常量，如 23、3.14159 和 math.Pi，占据了一种理想的数字空间，具有任意的精度，没有溢出或下溢。例如，math.Pi的值在源代码中被指定为63位，涉及该值的常量表达式保持的精度超出了 `float64` 所能容纳的范围。只有当常量或常量表达式被分配到一个变量--程序中的内存位置--时，它才会成为一个具有通常浮点属性和精度的 "计算机 "数字。

此外，由于常量只是数字，而不是键入值，因此在围棋中，常量可以比变量更自由地使用，从而缓解了严格的转换规则带来的一些尴尬。我们可以写一些表达式，比如

```
sqrt2 :=math.Sqrt(2)
```


一篇名为Constants的博客文章详细探讨了这个话题。

### Why are maps built in?
The same reason strings are: they are such a powerful and important data structure that providing one excellent implementation with syntactic support makes programming more pleasant. We believe that Go's implementation of maps is strong enough that it will serve for the vast majority of uses. If a specific application can benefit from a custom implementation, it's possible to write one but it will not be as convenient syntactically; this seems a reasonable tradeoff.

### 为什么内置 map ？
和字符串的原因是一样的：它们是如此强大和重要的数据结构，提供一个优秀的、具有语法支持的实现会使编程更加愉快。我们相信，`Go` 的字典实现足够强大，它将为绝大多数用途服务。如果一个特定的应用可以从定制的实现中受益，那么可以写一个，但在语法上不会那么方便；这似乎是一个合理的折衷。

### Why don't maps allow slices as keys?
Map lookup requires an equality operator, which slices do not implement. They don't implement equality because equality is not well defined on such types; there are multiple considerations involving shallow vs. deep comparison, pointer vs. value comparison, how to deal with recursive types, and so on. We may revisit this issue—and implementing equality for slices will not invalidate any existing programs—but without a clear idea of what equality of slices should mean, it was simpler to leave it out for now.

In Go 1, unlike prior releases, equality is defined for structs and arrays, so such types can be used as map keys. Slices still do not have a definition of equality, though.

### 为什么 map 不允许 slice 作为键？
`map` 查找需要一个平等操作符，而切片没有实现。它们不实现平等是因为平等在这类类型上没有很好的定义；有多种考虑因素，包括浅层与深层比较，指针与值比较，如何处理递归类型，等等。我们可能会重新讨论这个问题--而且实现分片的平等不会使任何现有的程序失效--但在没有明确分片的平等应该意味着什么的情况下，暂时不讨论这个问题比较简单。

在 `Go 1` 中，与之前的版本不同，平等性是为结构和数组定义的，所以这类类型可以作为映射键使用。不过，切片仍然没有平等的定义。

### Why are maps, slices, and channels references while arrays are values?
There's a lot of history on that topic. Early on, maps and channels were syntactically pointers and it was impossible to declare or use a non-pointer instance. Also, we struggled with how arrays should work. Eventually we decided that the strict separation of pointers and values made the language harder to use. Changing these types to act as references to the associated, shared data structures resolved these issues. This change added some regrettable complexity to the language but had a large effect on usability: Go became a more productive, comfortable language when it was introduced.

### 为什么 map、切片和通道是引用，而数组是值？
关于这个话题有很多历史。早期，字典和通道在语法上是指针，不可能声明或使用一个非指针实例。另外，我们也在为数组应该如何工作而苦恼。最终我们决定，指针和值的严格分离使语言更难使用。将这些类型改变为作为相关的、共享的数据结构的引用，解决了这些问题。这个变化给语言增加了一些令人遗憾的复杂性，但对可用性有很大的影响。当 `Go` 被引入时，它成为一种更有成效、更舒适的语言。


## Writing Code（编码）
### How are libraries documented?
There is a program, godoc, written in Go, that extracts package documentation from the source code and serves it as a web page with links to declarations, files, and so on. An instance is running at golang.org/pkg/. In fact, godoc implements the full site at golang.org/.

A godoc instance may be configured to provide rich, interactive static analyses of symbols in the programs it displays; details are listed here.

For access to documentation from the command line, the go tool has a doc subcommand that provides a textual interface to the same information.

### 如何对库进行文档化？
有一个用 `Go` 编写的程序，`godoc`，可以从源代码中提取包的文档，并以网页的形式提供给用户，其中有声明、文件等的链接。一个实例正在 `golang.org/pkg/`上运行。事实上，`godoc` 在 `golang.org` 实现了完整的站点。

一个 `godoc` 实例可以被配置为对它所显示的程序中的符号进行丰富的、交互式的静态分析；细节在这里列出。

对于从命令行访问文档，`go` 工具有一个 `doc` 子命令，它提供了一个文本界面来获取相同的信息。

### Is there a Go programming style guide?
There is no explicit style guide, although there is certainly a recognizable "Go style".

Go has established conventions to guide decisions around naming, layout, and file organization. The document [Effective Go](https://golang.org/doc/effective_go.html) contains some advice on these topics. More directly, the program gofmt is a pretty-printer whose purpose is to enforce layout rules; it replaces the usual compendium of do's and don'ts that allows interpretation. All the Go code in the repository, and the vast majority in the open source world, has been run through gofmt.

The document titled Go Code Review Comments is a collection of very short essays about details of Go idiom that are often missed by programmers. It is a handy reference for people doing code reviews for Go projects.

### 有Go编程风格指南吗？
没有明确的风格指南，尽管肯定有一个可识别的 "Go风格"。

`Go` 已经建立了一些惯例来指导命名、布局和文件组织方面的决定。文件 `Effective Go` 包含了一些关于这些主题的建议。更直接地说，程序 `gofmt` 是一个 `pretty-printer`，它的目的是执行布局规则；它取代了通常的do's和don's的汇编，允许解释。仓库中的所有围棋代码，以及开源世界中的绝大多数围棋代码，都是通过 `gofmt` 运行的。

这份名为《Go Code Review Comments》的文档，是一个非常简短的文章集，内容是关于 `Go` 习语中经常被程序员忽略的细节。对于为 `Go` 项目做代码审查的人来说，它是一个方便的参考。

### How do I submit patches to the Go libraries?
The library sources are in the src directory of the repository. If you want to make a significant change, please discuss on the mailing list before embarking.

See the document Contributing to the Go project for more information about how to proceed.

### 我如何向Go库提交补丁？
库的源代码在版本库的 `src` 目录下。如果你想做一个重大的改动，请在开始之前在邮件列表中讨论。

参见文档 Contributing to the Go project 获取更多关于如何进行的信息。

### Why does "go get" use HTTPS when cloning a repository?
Companies often permit outgoing traffic only on the standard TCP ports 80 (HTTP) and 443 (HTTPS), blocking outgoing traffic on other ports, including TCP port 9418 (git) and TCP port 22 (SSH). When using HTTPS instead of HTTP, git enforces certificate validation by default, providing protection against man-in-the-middle, eavesdropping and tampering attacks. The go get command therefore uses HTTPS for safety.

Git can be configured to authenticate over HTTPS or to use SSH in place of HTTPS. To authenticate over HTTPS, you can add a line to the $HOME/.netrc file that git consults:

machine github.com login USERNAME password APIKEY
For GitHub accounts, the password can be a personal access token.

Git can also be configured to use SSH in place of HTTPS for URLs matching a given prefix. For example, to use SSH for all GitHub access, add these lines to your ~/.gitconfig:

```
[url "ssh://git@github.com/"]
	insteadOf = https://github.com/
```

### 为什么 "Go get "在克隆版本库时使用HTTPS？
公司通常只允许在标准的 TCP 端口 80 (HTTP) 和 443 (HTTPS) 上发送流量，而阻止其他端口的发送流量，包括 TCP 端口 9418 (git) 和 TCP 端口 22 (SSH)。当使用HTTPS而不是HTTP时，`git` 会默认执行证书验证，以防止中间人、窃听和篡改攻击。因此，go get 命令使用 HTTPS 以确保安全。

Git 可以被配置为通过 HTTPS 进行身份验证，或者使用 SSH 来代替 HTTPS。要通过 HTTPS 进行身份验证，可以在 git 使用的 $HOME/.netrc 文件中添加一行。

机器 github.com 登录 USERNAME 密码 APIKEY
对于 GitHub 账户，密码可以是个人访问令牌。

Git 也可以配置为使用 SSH 代替 HTTPS 访问指定前缀的 URL。例如，如果要在所有 GitHub 的访问中使用 SSH，请在 ~/.gitconfig 中添加以下内容。

```
[url "ssh://git@github.com/"]
	insteadOf = https://github.com/
```

### How should I manage package versions using "go get"?
Since the inception of the project, Go has had no explicit concept of package versions, but that is changing. Versioning is a source of significant complexity, especially in large code bases, and it has taken some time to develop an approach that works well at scale in a large enough variety of situations to be appropriate to supply to all Go users.

The Go 1.11 release adds new, experimental support for package versioning to the go command, in the form of Go modules. For more information, see the Go 1.11 release notes and the go command documentation.

Regardless of the actual package management technology, "go get" and the larger Go toolchain does provide isolation of packages with different import paths. For example, the standard library's html/template and text/template coexist even though both are "package template". This observation leads to some advice for package authors and package users.

Packages intended for public use should try to maintain backwards compatibility as they evolve. The Go 1 compatibility guidelines are a good reference here: don't remove exported names, encourage tagged composite literals, and so on. If different functionality is required, add a new name instead of changing an old one. If a complete break is required, create a new package with a new import path.

If you're using an externally supplied package and worry that it might change in unexpected ways, but are not yet using Go modules, the simplest solution is to copy it to your local repository. This is the approach Google takes internally and is supported by the go command through a technique called "vendoring". This involves storing a copy of the dependency under a new import path that identifies it as a local copy. See the design document for details.

### 我应该如何使用 "go get "管理包的版本？
从项目开始，Go就没有明确的包版本概念，但这正在改变。版本管理是一个非常复杂的源头，尤其是在大型代码库中，而且花了一些时间来开发一种在足够多的情况下能够很好地大规模运行的方法，以适合提供给所有Go用户。

Go 1.11版本以Go模块的形式，为go命令增加了新的、实验性的包版本支持。更多信息，请参阅Go 1.11版本说明和go命令文档。

无论实际的包管理技术如何，"go get "和更大的Go工具链确实提供了不同导入路径的包的隔离。例如，标准库的html/template和text/template共存，尽管都是 "包模板"。这个观察结果导致了对包作者和包用户的一些建议。

面向公众使用的软件包在发展过程中应尽量保持向后兼容。Go 1的兼容性指南在这里是一个很好的参考：不要删除导出的名称，鼓励使用标记的复合字元，等等。如果需要不同的功能，添加一个新的名字，而不是改变一个旧的名字。如果需要完全中断，用新的导入路径创建一个新的包。

如果你使用的是外部提供的包，担心它可能会发生意想不到的变化，但还没有使用Go模块，最简单的解决方案是将其复制到本地仓库。这是Google内部采用的方法，go命令通过一种叫做 "vendoring "的技术来支持。这涉及到在一个新的导入路径下存储依赖的副本，并将其标识为本地副本。详情请看设计文档。

## Pointers and Allocation（指针和内存分配）
### When are function parameters passed by value?
As in all languages in the C family, everything in Go is passed by value. That is, a function always gets a copy of the thing being passed, as if there were an assignment statement assigning the value to the parameter. For instance, passing an int value to a function makes a copy of the int, and passing a pointer value makes a copy of the pointer, but not the data it points to. (See a later section for a discussion of how this affects method receivers.)

Map and slice values behave like pointers: they are descriptors that contain pointers to the underlying map or slice data. Copying a map or slice value doesn't copy the data it points to. Copying an interface value makes a copy of the thing stored in the interface value. If the interface value holds a struct, copying the interface value makes a copy of the struct. If the interface value holds a pointer, copying the interface value makes a copy of the pointer, but again not the data it points to.

Note that this discussion is about the semantics of the operations. Actual implementations may apply optimizations to avoid copying as long as the optimizations do not change the semantics.

### 什么时候函数参数是通过值传递的？
和 `C` 语言家族的所有语言一样，`Go` 中的所有东西都是通过值传递的。也就是说，一个函数总是得到一个被传递的东西的副本，就像有一个赋值语句把值分配给参数一样。例如，将一个 `int` 值传递给函数，就会得到 `int` 的副本，而传递一个指针值，就会得到指针的副本，但不会得到它所指向的数据。(关于这如何影响方法接收者，请参见后面的章节。)

映射和切片值的行为就像指针一样：它们是包含指向底层映射或切片数据的指针的描述符。复制一个 `map` 或 `slice` 值并不会复制它所指向的数据。复制一个接口值会对存储在接口值中的东西进行复制。如果接口值持有一个结构体，复制接口值就会复制该结构体。如果接口值持有一个指针，复制接口值会复制一个指针，但同样不会复制它所指向的数据。

注意，这个讨论是关于操作的语义。实际实现可以应用优化来避免复制，只要优化不改变语义即可。

### When should I use a pointer to an interface?
Almost never. Pointers to interface values arise only in rare, tricky situations involving disguising an interface value's type for delayed evaluation.

It is a common mistake to pass a pointer to an interface value to a function expecting an interface. The compiler will complain about this error but the situation can still be confusing, because sometimes a pointer is necessary to satisfy an interface. The insight is that although a pointer to a concrete type can satisfy an interface, with one exception a pointer to an interface can never satisfy an interface.

Consider the variable declaration,

```
var w io.Writer
```

The printing function fmt.Fprintf takes as its first argument a value that satisfies io.Writer—something that implements the canonical Write method. Thus we can write

```
fmt.Fprintf(w, "hello, world\n")
```

If however we pass the address of w, the program will not compile.

```
fmt.Fprintf(&w, "hello, world\n") // Compile-time error.
```

The one exception is that any value, even a pointer to an interface, can be assigned to a variable of empty interface type (interface{}). Even so, it's almost certainly a mistake if the value is a pointer to an interface; the result can be confusing.

### 什么时候应该使用指向接口的指针？
几乎不用。指针指向接口值的情况只有在极少数的、棘手的情况下才会出现，包括为了延迟评估而掩盖接口值的类型。

将一个指向接口值的指针传递给一个期望接口的函数是一个常见的错误。编译器会抛这个错误，但这种情况仍然会让人感到困惑，因为有时指针是满足接口的必要条件。我们的见解是，虽然指向具体类型的指针可以满足一个接口，但有一个例外，指向接口的指针永远不能满足一个接口。

考虑一下变量声明。

```
var w io.Writer
```

打印函数fmt.Fprintf的第一个参数是一个满足 `io.Writer` 的值，这个值实现了规范的 `Write` 方法。因此我们可以写

```
fmt.Fprintf(w, "hello, world/n")
```

但是如果我们传递了w的地址，程序将无法编译。

```
fmt.Fprintf(&w, "hello, world/n") // 编译时错误。
```

一个例外是，任何值，甚至是一个接口的指针，都可以分配给一个空接口类型的变量（`interface{}`）。即便如此，如果值是一个指向接口的指针，几乎可以肯定是一个错误；结果可能会让人困惑。

### Should I define methods on values or pointers?
```
func (s *MyStruct) pointerMethod() { } // method on pointer
func (s MyStruct)  valueMethod()   { } // method on value
```

For programmers unaccustomed to pointers, the distinction between these two examples can be confusing, but the situation is actually very simple. When defining a method on a type, the receiver (s in the above examples) behaves exactly as if it were an argument to the method. Whether to define the receiver as a value or as a pointer is the same question, then, as whether a function argument should be a value or a pointer. There are several considerations.

First, and most important, does the method need to modify the receiver? If it does, the receiver must be a pointer. (Slices and maps act as references, so their story is a little more subtle, but for instance to change the length of a slice in a method the receiver must still be a pointer.) In the examples above, if pointerMethod modifies the fields of s, the caller will see those changes, but valueMethod is called with a copy of the caller's argument (that's the definition of passing a value), so changes it makes will be invisible to the caller.

By the way, in Java method receivers are always pointers, although their pointer nature is somewhat disguised (and there is a proposal to add value receivers to the language). It is the value receivers in Go that are unusual.

Second is the consideration of efficiency. If the receiver is large, a big struct for instance, it will be much cheaper to use a pointer receiver.

Next is consistency. If some of the methods of the type must have pointer receivers, the rest should too, so the method set is consistent regardless of how the type is used. See the section on method sets for details.

For types such as basic types, slices, and small structs, a value receiver is very cheap so unless the semantics of the method requires a pointer, a value receiver is efficient and clear.

### 我应该在值或指针上定义方法？
```
func (s *MyStruct) pointerMethod() { } // 在指针上定义方法。
func (s MyStruct) valueMethod() { } // method on value
```

对于不习惯使用指针的程序员来说，这两个例子之间的区别可能会让人感到困惑，但实际上情况非常简单。当在一个类型上定义一个方法时，接收器（在上面的例子中是s）的行为就像它是方法的一个参数一样。那么，是将接收器定义为值还是指针，这与函数参数应该是值还是指针是同一个问题。有几个考虑因素。

首先，也是最重要的，方法是否需要修改接收器？如果需要，接收器必须是一个指针。(切片和字典作为引用，所以它们的故事更微妙一些，但例如要在方法中改变切片的长度，接收器仍然必须是一个指针。) 在上面的例子中，如果 `pointerMethod` 修改了 `s` 的字段，调用者会看到这些变化，但是 `valueMethod` 是用调用者参数的副本来调用的(这就是传值的定义)，所以它所做的变化对调用者来说是看不见的。

顺便说一下，在 `Java` 中，方法接收器总是指针，虽然它们的指针性质有些被掩盖了（有人提议在语言中增加值接收器）。`Go` 中的值接受器才是与众不同的。

其次是对效率的考虑。如果接收器很大，比如一个大的结构，使用指针接收器会便宜很多。

其次是一致性。如果类型的某些方法必须有指针接收器，那么其余的方法也应该有指针接收器，所以无论类型如何使用，方法集都是一致的。详见方法集一节。

对于基本类型、切片和小结构等类型，值接收器是非常廉价的，所以除非方法的语义需要指针，否则值接收器是高效和清晰的。

### What's the difference between new and make?
In short: new allocates memory, while make initializes the slice, map, and channel types.

See the relevant section of Effective Go for more details.

### new和make的区别是什么？
简而言之：new分配内存，而make初始化slice、map和channel类型。

更多细节请参考Effective Go的相关章节。

### What is the size of an int on a 64 bit machine?
The sizes of int and uint are implementation-specific but the same as each other on a given platform. For portability, code that relies on a particular size of value should use an explicitly sized type, like int64. On 32-bit machines the compilers use 32-bit integers by default, while on 64-bit machines integers have 64 bits. (Historically, this was not always true.)

On the other hand, floating-point scalars and complex types are always sized (there are no float or complex basic types), because programmers should be aware of precision when using floating-point numbers. The default type used for an (untyped) floating-point constant is float64. Thus foo := 3.0 declares a variable foo of type float64. For a float32 variable initialized by an (untyped) constant, the variable type must be specified explicitly in the variable declaration:

```
var foo float32 = 3.0
```

Alternatively, the constant must be given a type with a conversion as in foo := float32(3.0).

### 在64位机器上int的大小是多少？
`int` 和 `uint` 的大小是特定的实现，但在给定的平台上是一样的。为了便于移植，依赖于特定大小的代码应该使用一个显式大小的类型，比如 `int64`。在 `32` 位机器上，编译器默认使用 `32` 位整数，而在 `64` 位机器上，整数有 `64` 位。(从历史上看，这并不总是正确的。)

另一方面，浮点标量和复杂类型总是有大小的（没有浮点或复杂的基本类型），因为程序员在使用浮点数时应该注意精度。(无类型的)浮点常量的默认类型是float64。因此 `foo :=3.0` 声明了一个类型为 `float64` 的变量。对于一个由（无类型）常量初始化的 `float32` 变量，必须在变量声明中明确指定变量类型。

```
var foo float32 = 3.0
```

或者，常量必须被赋予一个带有转换的类型，如foo := float32(3.0)。

### How do I know whether a variable is allocated on the heap or the stack?
From a correctness standpoint, you don't need to know. Each variable in Go exists as long as there are references to it. The storage location chosen by the implementation is irrelevant to the semantics of the language.

The storage location does have an effect on writing efficient programs. When possible, the Go compilers will allocate variables that are local to a function in that function's stack frame. However, if the compiler cannot prove that the variable is not referenced after the function returns, then the compiler must allocate the variable on the garbage-collected heap to avoid dangling pointer errors. Also, if a local variable is very large, it might make more sense to store it on the heap rather than the stack.

In the current compilers, if a variable has its address taken, that variable is a candidate for allocation on the heap. However, a basic escape analysis recognizes some cases when such variables will not live past the return from the function and can reside on the stack.

### 我如何知道一个变量是在堆上还是在栈上分配？
从正确性的角度来看，你不需要知道。在 Go 中，只要有对它的引用，每个变量就会存在。实现所选择的存储位置与语言的语义无关。

存储位置对编写高效程序确实有影响。在可能的情况下，Go 编译器会在一个函数的堆栈框架中分配该函数的局部变量。但是，如果编译器不能证明该变量在函数返回后没有被引用，那么编译器就必须在垃圾收集堆上分配该变量，以避免悬空指针错误。另外，如果一个局部变量非常大，那么将其存储在堆上而不是堆上可能更有意义。

在目前的编译器中，如果一个变量的地址被取走了，那么这个变量就是堆上分配的候选变量。然而，一个基本的转义分析可以识别出一些情况，当这样的变量不会活过函数的返回，而是可以驻留在堆栈上。

### Why does my Go process use so much virtual memory?
The Go memory allocator reserves a large region of virtual memory as an arena for allocations. This virtual memory is local to the specific Go process; the reservation does not deprive other processes of memory.

To find the amount of actual memory allocated to a Go process, use the Unix top command and consult the RES (Linux) or RSIZE (macOS) columns.

### 为什么我的Go进程使用了这么多虚拟内存？
`Go` 内存分配器保留了一个很大的虚拟内存区域作为分配的舞台。这个虚拟内存是特定的围棋进程的本地内存；保留并没有剥夺其他进程的内存。

要想知道实际分配给Go进程的内存量，可以使用Unix顶级命令，并查阅 `RES`（Linux）或 `RSIZE`（macOS）列。

## Concurrency（并发）
### What operations are atomic? What about mutexes?
A description of the atomicity of operations in Go can be found in the Go Memory Model document.

Low-level synchronization and atomic primitives are available in the sync and sync/atomic packages. These packages are good for simple tasks such as incrementing reference counts or guaranteeing small-scale mutual exclusion.

For higher-level operations, such as coordination among concurrent servers, higher-level techniques can lead to nicer programs, and Go supports this approach through its goroutines and channels. For instance, you can structure your program so that only one goroutine at a time is ever responsible for a particular piece of data. That approach is summarized by the original Go proverb,

Do not communicate by sharing memory. Instead, share memory by communicating.

See the [Share Memory By Communicating](https://golang.org/doc/codewalk/sharemem/) code walk and its associated article for a detailed discussion of this concept.

Large concurrent programs are likely to borrow from both these toolkits.

### 什么操作是原子的？互斥呢？
关于 Go 中操作的原子性的描述可以在 Go 内存模型文档中找到。

sync 和 sync/atomic 包中提供了低级同步和原子基元。这些包适用于简单的任务，如增量引用计数或保证小规模的互斥。

对于更高层次的操作，例如并发服务器之间的协调，更高层次的技术可以带来更好的程序，Go 通过其 goroutines 和通道支持这种方法。例如，你可以将你的程序结构化，使每次只有一个 goroutine 负责某项数据。这种方法是由最初的 Go 谚语总结出来的。

不要通过分享记忆来沟通。相反，通过通信来共享内存。

关于这个概念的详细讨论，请参见 Share Memory By Communicating 代码走势及其相关文章。

大型并发程序很可能会借鉴这两个工具箱。

### Why doesn't my program run faster with more CPUs?
Whether a program runs faster with more CPUs depends on the problem it is solving. The Go language provides concurrency primitives, such as goroutines and channels, but concurrency only enables parallelism when the underlying problem is intrinsically parallel. Problems that are intrinsically sequential cannot be sped up by adding more CPUs, while those that can be broken into pieces that can execute in parallel can be sped up, sometimes dramatically.

Sometimes adding more CPUs can slow a program down. In practical terms, programs that spend more time synchronizing or communicating than doing useful computation may experience performance degradation when using multiple OS threads. This is because passing data between threads involves switching contexts, which has significant cost, and that cost can increase with more CPUs. For instance, the prime sieve example from the Go specification has no significant parallelism although it launches many goroutines; increasing the number of threads (CPUs) is more likely to slow it down than to speed it up.

For more detail on this topic see the talk entitled [Concurrency is not Parallelism](https://blog.golang.org/2013/01/concurrency-is-not-parallelism.html).

### 为什么我的程序在更多CPU的情况下运行速度不快？
一个程序是否在更多CPU的情况下运行得更快，取决于它要解决的问题。Go 语言提供了并发基元，如 goroutines 和通道，但只有当底层问题本质上是并行的时候，并发才能实现并行。本质上是顺序的问题不能通过增加更多的CPU来加速，而那些可以分解成可以并行执行的碎片的问题则可以加速，有时甚至是极大的加速。

有时，增加更多的CPU会使程序的速度变慢。从实际情况来看，当使用多个操作系统线程时，那些花在同步或通信上的时间比做有用计算的时间多的程序可能会出现性能下降。这是因为线程之间传递数据涉及到切换上下文，这有很大的成本，而且这种成本会随着CPU的增加而增加。例如，Go 规范中的素筛示例虽然启动了许多 goroutine，但没有显著的并行性；增加线程（CPU）数量更有可能使其速度减慢而不是加快。

关于这个话题的更多细节，请看题为《并发不是并行》的演讲。

### How can I control the number of CPUs?
The number of CPUs available simultaneously to executing goroutines is controlled by the GOMAXPROCS shell environment variable, whose default value is the number of CPU cores available. Programs with the potential for parallel execution should therefore achieve it by default on a multiple-CPU machine. To change the number of parallel CPUs to use, set the environment variable or use the similarly-named function of the runtime package to configure the run-time support to utilize a different number of threads. Setting it to 1 eliminates the possibility of true parallelism, forcing independent goroutines to take turns executing.

The runtime can allocate more threads than the value of GOMAXPROCS to service multiple outstanding I/O requests. GOMAXPROCS only affects how many goroutines can actually execute at once; arbitrarily more may be blocked in system calls.

Go's goroutine scheduler is not as good as it needs to be, although it has improved over time. In the future, it may better optimize its use of OS threads. For now, if there are performance issues, setting GOMAXPROCS on a per-application basis may help.

### 如何控制CPU的数量？
同时执行 goroutine 的CPU数量由 `GOMAXPROCS` shell 环境变量控制，其默认值是可用的CPU核数。因此，具有并行执行潜力的程序应该在多CPU机器上默认实现。要改变要使用的并行CPU数，可以设置环境变量或使用运行时包的类似名称的函数来配置运行时支持利用不同数量的线程。将其设置为1，则消除了真正并行的可能性，迫使独立的 goroutine 轮流执行。

运行时可以分配比 `GOMAXPROCS` 值更多的线程来服务多个未完成的I/O请求。`GOMAXPROCS` 只影响一次实际可以执行多少个 goroutine；任意更多的线程可能在系统调用中被阻塞。

Go 的 goroutine 调度器并不尽如人意，尽管随着时间的推移，它已经有所改进。在未来，它可能会更好地优化其对操作系统线程的使用。目前，如果有性能问题，在每个应用程序的基础上设置 `GOMAXPROCS` 可能会有帮助。

### Why is there no goroutine ID?
Goroutines do not have names; they are just anonymous workers. They expose no unique identifier, name, or data structure to the programmer. Some people are surprised by this, expecting the go statement to return some item that can be used to access and control the goroutine later.

The fundamental reason goroutines are anonymous is so that the full Go language is available when programming concurrent code. By contrast, the usage patterns that develop when threads and goroutines are named can restrict what a library using them can do.

Here is an illustration of the difficulties. Once one names a goroutine and constructs a model around it, it becomes special, and one is tempted to associate all computation with that goroutine, ignoring the possibility of using multiple, possibly shared goroutines for the processing. If the net/http package associated per-request state with a goroutine, clients would be unable to use more goroutines when serving a request.

Moreover, experience with libraries such as those for graphics systems that require all processing to occur on the "main thread" has shown how awkward and limiting the approach can be when deployed in a concurrent language. The very existence of a special thread or goroutine forces the programmer to distort the program to avoid crashes and other problems caused by inadvertently operating on the wrong thread.

For those cases where a particular goroutine is truly special, the language provides features such as channels that can be used in flexible ways to interact with it.

### 为什么没有goroutine ID？
Goroutine 没有名字，它们只是一个匿名的工作者。它们没有向程序员暴露任何独特的标识符、名称或数据结构。有些人对此感到惊讶，期望 `go` 语句能够返回一些项目，以便以后可以用来访问和控制 goroutine。

goroutine 之所以是匿名的，根本原因是为了在并发代码编程时可以使用完整的Go语言。相比之下，当线程和 goroutine 被命名后形成的使用模式会限制使用它们的库的功能。

下面是一个困难的例子。一旦我们命名了一个 goroutine 并围绕它构建了一个模型，它就变得特殊了，人们就会将所有的计算与这个 goroutine 联系起来，而忽略了使用多个可能共享的 goroutine 进行处理的可能性。如果 `net/http` 包将每个请求的状态与 goroutine 相关联，那么客户端在服务一个请求时就无法使用更多的goroutine。

此外，诸如图形系统的库的经验表明，当部署在并发语言中时，这种方法是多么的笨拙和受限。特殊线程或 goroutine 的存在本身就迫使程序员扭曲程序，以避免无意中在错误的线程上操作而导致的崩溃和其他问题。

对于那些真正特殊的 goroutine 的情况，语言提供了通道等功能，可以用灵活的方式与之交互。

## Functions and Methods（函数和方法）
### Why do T and *T have different method sets?
As the Go specification says, the method set of a type T consists of all methods with receiver type T, while that of the corresponding pointer type *T consists of all methods with receiver *T or T. That means the method set of *T includes that of T, but not the reverse.

This distinction arises because if an interface value contains a pointer *T, a method call can obtain a value by dereferencing the pointer, but if an interface value contains a value T, there is no safe way for a method call to obtain a pointer. (Doing so would allow a method to modify the contents of the value inside the interface, which is not permitted by the language specification.)

Even in cases where the compiler could take the address of a value to pass to the method, if the method modifies the value the changes will be lost in the caller. As an example, if the Write method of bytes.Buffer used a value receiver rather than a pointer, this code:

```
var buf bytes.Buffer
io.Copy(buf, os.Stdin)
```

would copy standard input into a copy of buf, not into buf itself. This is almost never the desired behavior.

### 为什么T和*T有不同的方法集？
正如 Go 规范所说，类型 `T` 的方法集由所有接收方类型为 `T` 的方法组成，而对应的指针类型 `*T` 的方法集由所有接收方为 `*T` 或 `T` 的方法组成，这意味着 `*T` 的方法集包括 `T` 的方法集，而不是相反。

这种区别产生的原因是，如果一个接口值包含一个指针 `*T` ，方法调用可以通过去引用该指针来获得一个值，但是如果一个接口值包含一个值 `T`，方法调用就没有安全的方法来获得一个指针。(这样做会让方法修改接口内部的值的内容，这是语言规范所不允许的)。

即使在编译器可以取值的地址传递给方法的情况下，如果方法修改了值，那么修改的内容也会在调用者中丢失。举个例子，如果 bytes.Buffer` 的 `Write` 方法使用一个值接收器而不是指针，这段代码。

```
var buf bytes.Buffer
io.Copy(buf, os.Stdin)
```

会将标准输入复制到 buf 的副本中，而不是复制到 buf 本身。这几乎不是我们想要的行为。


### What happens with closures running as goroutines?
Some confusion may arise when using closures with concurrency. Consider the following program:

    func main() {
        done := make(chan bool)
        values := []string{"a", "b", "c"}
        for _, v := range values {
            go func() {
                fmt.Println(v)
                done <- true
            }()
        }
    
        // wait for all goroutines to complete before exiting
        for _ = range values {
            <-done
        }
    }

One might mistakenly expect to see a, b, c as the output. What you'll probably see instead is c, c, c. This is because each iteration of the loop uses the same instance of the variable v, so each closure shares that single variable. When the closure runs, it prints the value of v at the time fmt.Println is executed, but v may have been modified since the goroutine was launched. To help detect this and other problems before they happen, run go vet.

To bind the current value of v to each closure as it is launched, one must modify the inner loop to create a new variable each iteration. One way is to pass the variable as an argument to the closure:

    for _, v := range values {
        go func(u string) {
            fmt.Println(u)
            done <- true
        }(v)
    }
In this example, the value of v is passed as an argument to the anonymous function. That value is then accessible inside the function as the variable u.

Even easier is just to create a new variable, using a declaration style that may seem odd but works fine in Go:

    for _, v := range values {
        v := v // create a new 'v'.
        go func() {
            fmt.Println(v)
            done <- true
        }()
    }
This behavior of the language, not defining a new variable for each iteration, may have been a mistake in retrospect. It may be addressed in a later version but, for compatibility, cannot change in Go version 1.

### 以goroutine的形式运行闭包会怎样？
当使用并发的闭包时，可能会出现一些混淆。考虑一下下面的程序。

    func main() {
        done := make(chan bool)
        values := []string{"a", "b", "c"}。
        for _, v := range values {
            go func() {
                fmt.Println(v)
                done <- true
            }()
        }
    
        // 等待所有的goroutines完成后再退出。
        for _ = range values {
            <完成
        }
    }

人们可能会误以为输出的是a, b, c。这是因为循环的每一次迭代都使用同一个变量v的实例，所以每个闭包都共享这个单一变量。当闭包运行时，它会在执行 `fmt.Println` 时打印 `v` 的值，但 `v` 可能在 goroutine 启动后被修改过。为了帮助在问题发生之前检测到这个问题和其他问题，请运行 `go vet`。

为了在启动时将 `v` 的当前值绑定到每个闭包中，必须修改内部循环以在每次迭代时创建一个新的变量。一种方法是将变量作为一个参数传递给闭包。

    for _, v := range values {
        go func(u string) {
            fmt.Println(u)
            done <- true
        }(v)
    }
在这个例子中，v的值被作为参数传递给匿名函数。然后这个值就可以在函数中作为变量 u 访问。

更简单的方法是创建一个新的变量，使用一种可能看起来很奇怪但在Go中很好用的声明风格。

    for _, v := range values {
        v := v // 创建一个新的 "v"。
        go func() {
            fmt.Println(v)
            done <- true
        }()
    }
语言的这种行为，没有为每次迭代定义一个新的变量，回想起来可能是个错误。在以后的版本中可能会解决这个问题，但为了兼容性，在Go版本1中不能改变。

## Control flow（流程控制）
### Why does Go not have the ?: operator?
There is no ternary testing operation in Go. You may use the following to achieve the same result:

```
if expr {
    n = trueVal
} else {
    n = falseVal
}
```

The reason ?: is absent from Go is that the language's designers had seen the operation used too often to create impenetrably complex expressions. The if-else form, although longer, is unquestionably clearer. A language needs only one conditional control flow construct.

Go中没有三元测试操作。你可以用下面的方法来实现同样的结果。

```
if expr {
    n = trueVal
} else {
    n = falseVal
}
```

Go中没有使用?:的原因是，语言的设计者们看到这种操作使用得太频繁了，以至于产生了难以捉摸的复杂表达式。if-else形式虽然较长，但无疑更清晰。一门语言只需要一个条件控制流结构。

## Packages and Testing（包和测试）
### How do I create a multifile package?
Put all the source files for the package in a directory by themselves. Source files can refer to items from different files at will; there is no need for forward declarations or a header file.

Other than being split into multiple files, the package will compile and test just like a single-file package.

### 如何创建一个多文件包？
把包的所有源文件都单独放在一个目录中。源文件可以随意引用不同文件中的项目，不需要前向声明或头文件。

除了被分割成多个文件外，包将像单文件包一样进行编译和测试。

### How do I write a unit test?
Create a new file ending in _test.go in the same directory as your package sources. Inside that file, import "testing" and write functions of the form

```
func TestFoo(t *testing.T) {
    ...
}
```

Run go test in that directory. That script finds the Test functions, builds a test binary, and runs it.

See the How to Write Go Code document, the testing package and the go test subcommand for more details.

### 如何编写单元测试？
在你的包源文件的同一目录下创建一个以 _test.go 结尾的新文件。在该文件中，导入 "testing"，并写下如下形式的函数。

```
func TestFoo(t *testing.T) {
    ...
}
```

在该目录下运行go test。该脚本找到测试函数，构建一个测试二进制文件，并运行它。

更多细节请参见 How to Write Go Code 文档，测试包和 go test 子命令。

### Where is my favorite helper function for testing?
Go's standard testing package makes it easy to write unit tests, but it lacks features provided in other language's testing frameworks such as assertion functions. An earlier section of this document explained why Go doesn't have assertions, and the same arguments apply to the use of assert in tests. Proper error handling means letting other tests run after one has failed, so that the person debugging the failure gets a complete picture of what is wrong. It is more useful for a test to report that isPrime gives the wrong answer for 2, 3, 5, and 7 (or for 2, 4, 8, and 16) than to report that isPrime gives the wrong answer for 2 and therefore no more tests were run. The programmer who triggers the test failure may not be familiar with the code that fails. Time invested writing a good error message now pays off later when the test breaks.

A related point is that testing frameworks tend to develop into mini-languages of their own, with conditionals and controls and printing mechanisms, but Go already has all those capabilities; why recreate them? We'd rather write tests in Go; it's one fewer language to learn and the approach keeps the tests straightforward and easy to understand.

If the amount of extra code required to write good errors seems repetitive and overwhelming, the test might work better if table-driven, iterating over a list of inputs and outputs defined in a data structure (Go has excellent support for data structure literals). The work to write a good test and good error messages will then be amortized over many test cases. The standard Go library is full of illustrative examples, such as in the formatting tests for the fmt package.

### 我最喜欢的测试辅助函数在哪里？
Go的标准测试包让编写单元测试变得很容易，但它缺乏其他语言的测试框架所提供的功能，比如断言函数。本文档前面的一节解释了为什么Go没有断言，同样的论点也适用于测试中使用断言。适当的错误处理意味着在一个测试失败后让其他测试运行，这样调试失败的人就能完整地了解问题所在。对于一个测试来说，报告 isPrime 给出了 2、3、5 和 7 的错误答案（或 2、4、8 和 16）比报告 isPrime 给出了 2 的错误答案，因此没有运行更多的测试更有用。触发测试失败的程序员可能不熟悉失败的代码。现在投入的时间写一个好的错误信息，在以后测试失败时就会得到回报。

相关的一点是，测试框架往往会发展成自己的迷你语言，有条件、控制和打印机制，但Go已经拥有所有这些功能；为什么要重新创建它们呢？我们宁可用Go来写测试，这样就少了一门需要学习的语言，而且这种方法能让测试保持简单明了，易于理解。

如果写好错误所需的额外代码量看起来很重复，让人不知所措，那么如果用表驱动，在数据结构中定义的输入和输出列表上迭代，测试可能会更好（Go对数据结构字面的支持很好）。然后，编写一个好的测试和好的错误信息的工作将被摊派到许多测试用例上。标准的 Go 库中有很多说明性的例子，比如在 fmt 包的格式化测试中。

### Why isn't X in the standard library?
The standard library's purpose is to support the runtime, connect to the operating system, and provide key functionality that many Go programs require, such as formatted I/O and networking. It also contains elements important for web programming, including cryptography and support for standards like HTTP, JSON, and XML.

There is no clear criterion that defines what is included because for a long time, this was the only Go library. There are criteria that define what gets added today, however.

New additions to the standard library are rare and the bar for inclusion is high. Code included in the standard library bears a large ongoing maintenance cost (often borne by those other than the original author), is subject to the Go 1 compatibility promise (blocking fixes to any flaws in the API), and is subject to the Go release schedule, preventing bug fixes from being available to users quickly.

Most new code should live outside of the standard library and be accessible via the go tool's go get command. Such code can have its own maintainers, release cycle, and compatibility guarantees. Users can find packages and read their documentation at godoc.org.

Although there are pieces in the standard library that don't really belong, such as log/syslog, we continue to maintain everything in the library because of the Go 1 compatibility promise. But we encourage most new code to live elsewhere.

### 为什么标准库中没有X？
标准库的目的是支持运行时，连接到操作系统，并提供许多Go程序所需的关键功能，如格式化I/O和网络。它还包含了对网络编程很重要的元素，包括密码学和对HTTP、JSON和XML等标准的支持。

没有明确的标准来定义包括哪些内容，因为在很长一段时间内，这是唯一的Go库。不过，今天有一些标准定义了哪些内容会被添加进来。

标准库中新加入的代码很少，而且收录的门槛很高。包含在标准库中的代码要承担很大的持续维护成本（通常由原作者以外的人承担），受制于Go 1的兼容性承诺（阻止对API中任何缺陷的修复），并且受制于Go的发布时间表，阻止错误修复快速提供给用户。

大多数新代码应该活在标准库之外，并通过go工具的go get命令来访问。这样的代码可以有自己的维护者、发布周期和兼容性保证。用户可以在godoc.org上找到包并阅读其文档。

虽然在标准库中有一些并不真正属于自己的部分，比如log/syslog，但由于Go 1的兼容性承诺，我们会继续维护库中的所有内容。但我们鼓励大多数新代码放在其他地方。

## Implementation（执行）
### What compiler technology is used to build the compilers?
There are several production compilers for Go, and a number of others in development for various platforms.

The default compiler, gc, is included with the Go distribution as part of the support for the go command. Gc was originally written in C because of the difficulties of bootstrapping—you'd need a Go compiler to set up a Go environment. But things have advanced and since the Go 1.5 release the compiler has been a Go program. The compiler was converted from C to Go using automatic translation tools, as described in this design document and talk. Thus the compiler is now "self-hosting", which means we needed to face the bootstrapping problem. The solution is to have a working Go installation already in place, just as one normally has with a working C installation. The story of how to bring up a new Go environment from source is described here and here.

Gc is written in Go with a recursive descent parser and uses a custom loader, also written in Go but based on the Plan 9 loader, to generate ELF/Mach-O/PE binaries.

At the beginning of the project we considered using LLVM for gc but decided it was too large and slow to meet our performance goals. More important in retrospect, starting with LLVM would have made it harder to introduce some of the ABI and related changes, such as stack management, that Go requires but are not part of the standard C setup. A new LLVM implementation is starting to come together now, however.

The Gccgo compiler is a front end written in C++ with a recursive descent parser coupled to the standard GCC back end.

Go turned out to be a fine language in which to implement a Go compiler, although that was not its original goal. Not being self-hosting from the beginning allowed Go's design to concentrate on its original use case, which was networked servers. Had we decided Go should compile itself early on, we might have ended up with a language targeted more for compiler construction, which is a worthy goal but not the one we had initially.

Although gc does not use them (yet?), a native lexer and parser are available in the go package and there is also a native type checker.

### 使用什么编译器技术来构建编译器？
有几个用于Go的生产编译器，还有一些用于不同平台的其他开发中的编译器。

默认的编译器，gc，作为go命令支持的一部分，包含在Go发行版中。Gc最初是用C语言编写的，因为引导困难--你需要一个Go编译器来建立一个Go环境。但事情已经进步了，自从Go 1.5版本发布后，编译器就成了Go程序。编译器是使用自动翻译工具从C语言转换为Go语言的，在这个设计文档和talk中都有描述。因此编译器现在是 "自托管"，这意味着我们需要面对引导问题。解决的办法是已经有一个工作的Go安装，就像通常有一个工作的C安装一样。关于如何从源码引入一个新的Go环境的故事，在这里和这里都有描述。

Gc是用Go编写的，它有一个递归下降解析器，并使用一个自定义的加载器，也是用Go编写的，但基于Plan 9加载器，来生成ELF/Mach-O/PE二进制文件。

在项目之初，我们考虑过使用LLVM来做gc，但认为它太大、太慢，无法满足我们的性能目标。回想起来更重要的是，从LLVM开始会使我们更难引入一些ABI和相关的变化，比如堆栈管理，这些都是Go需要的，但不是标准C设置的一部分。不过，现在一个新的LLVM实现开始出现了。

Gccgo编译器是一个用C++编写的前端，有一个递归下降解析器，耦合到标准的GCC后端。

Go原来是一种很好的语言，用它来实现Go编译器，虽然这不是它最初的目标。从一开始就不是自托管的，这使得 Go 的设计可以集中在它最初的使用案例上，即网络服务器。如果我们一开始就决定Go应该自编译，那么我们最终可能会得到一个更多针对编译器构造的语言，这是一个有价值的目标，但不是我们最初的目标。

虽然gc没有使用它们（还没有？），但go包里有一个原生的词典和解析器，还有一个原生的类型检查器。

### How is the run-time support implemented?
Again due to bootstrapping issues, the run-time code was originally written mostly in C (with a tiny bit of assembler) but it has since been translated to Go (except for some assembler bits). Gccgo's run-time support uses glibc. The gccgo compiler implements goroutines using a technique called segmented stacks, supported by recent modifications to the gold linker. Gollvm similarly is built on the corresponding LLVM infrastructure.

### 是如何实现运行时支持的？
同样由于bootstrapping的问题，运行时代码最初主要是用C语言编写的（有一小部分汇编器），但后来被翻译成Go语言（除了一些汇编器的部分）。Gccgo的运行时支持使用glibc。gccgo编译器使用一种叫做分段堆栈的技术来实现goroutines，由最近对gold linker的修改所支持。Gollvm同样是建立在相应的LLVM基础架构上。

### Why is my trivial program such a large binary?
The linker in the gc toolchain creates statically-linked binaries by default. All Go binaries therefore include the Go runtime, along with the run-time type information necessary to support dynamic type checks, reflection, and even panic-time stack traces.

A simple C "hello, world" program compiled and linked statically using gcc on Linux is around 750 kB, including an implementation of printf. An equivalent Go program using fmt.Printf weighs a couple of megabytes, but that includes more powerful run-time support and type and debugging information.

A Go program compiled with gc can be linked with the -ldflags=-w flag to disable DWARF generation, removing debugging information from the binary but with no other loss of functionality. This can reduce the binary size substantially.

### 为什么我的琐碎程序是个大二进制？
gc工具链中的链接器默认创建静态链接的二进制文件。因此，所有的 Go 二进制文件都包含了 Go 运行时，以及支持动态类型检查、反射，甚至是恐慌时堆栈跟踪所需的运行时类型信息。

一个简单的C语言 "hello, world "程序在Linux上使用gcc编译并静态链接，包括printf的实现，大约有750 kB。一个使用fmt.printf的等效Go程序重达几兆字节，但其中包括更强大的运行时支持和类型及调试信息。

用 gc 编译的 Go 程序可以用 -ldflags=-w 标志来禁用 DWARF 生成，从二进制中删除调试信息，但不会损失其他功能。这可以大大减少二进制的大小。

### Can I stop these complaints about my unused variable/import?
The presence of an unused variable may indicate a bug, while unused imports just slow down compilation, an effect that can become substantial as a program accumulates code and programmers over time. For these reasons, Go refuses to compile programs with unused variables or imports, trading short-term convenience for long-term build speed and program clarity.

Still, when developing code, it's common to create these situations temporarily and it can be annoying to have to edit them out before the program will compile.

Some have asked for a compiler option to turn those checks off or at least reduce them to warnings. Such an option has not been added, though, because compiler options should not affect the semantics of the language and because the Go compiler does not report warnings, only errors that prevent compilation.

There are two reasons for having no warnings. First, if it's worth complaining about, it's worth fixing in the code. (And if it's not worth fixing, it's not worth mentioning.) Second, having the compiler generate warnings encourages the implementation to warn about weak cases that can make compilation noisy, masking real errors that should be fixed.

It's easy to address the situation, though. Use the blank identifier to let unused things persist while you're developing.

```
import "unused"

// This declaration marks the import as used by referencing an
// item from the package.
var _ = unused.Item  // TODO: Delete before committing!

func main() {
    debugData := debug.Profile()
    _ = debugData // Used only during debugging.
    ....
}
```

Nowadays, most Go programmers use a tool, goimports, which automatically rewrites a Go source file to have the correct imports, eliminating the unused imports issue in practice. This program is easily connected to most editors to run automatically when a Go source file is written.

### 我可以停止这些关于我的未使用变量/导入的抱怨吗？
一个未使用的变量的存在可能表明是一个错误，而未使用的导入只是减慢了编译速度，当一个程序随着时间的推移积累代码和程序员时，这种影响可能会变得很大。基于这些原因，Go拒绝编译有未使用的变量或导入的程序，以短期的便利性换取长期的构建速度和程序的清晰度。

不过，在开发代码的时候，临时产生这些情况还是很常见的，在程序编译之前，不得不将这些情况编辑掉，这是很烦人的。

有些人要求提供一个编译器选项来关闭这些检查，或者至少把它们减少为警告。不过这样的选项并没有被加入，因为编译器选项不应该影响语言的语义，而且Go编译器不报告警告，只报告妨碍编译的错误。

没有警告的原因有两个。首先，如果值得抱怨，就值得在代码中修正。(如果不值得修复，那就不值得一提。)第二，让编译器产生警告，会鼓励实现对薄弱的情况发出警告，从而使编译变得嘈杂，掩盖了应该修复的真正错误。

不过要解决这种情况很容易。使用空白标识符，让未使用的东西在你开发的时候持续存在。

```
import "unused"

// This declaration marks the import as used by referencing an
// item from the package.
var _ = unused.Item  // TODO: Delete before committing!

func main() {
    debugData := debug.Profile()
    _ = debugData // Used only during debugging.
    ....
}
```

现在，大多数围棋程序员都会使用一个工具goimports，它可以自动重写一个围棋源文件，使其有正确的导入，在实践中消除了未使用的导入问题。这个程序很容易和大多数编辑器连接，在编写围棋源文件时自动运行。

### Why does my virus-scanning software think my Go distribution or compiled binary is infected?
This is a common occurrence, especially on Windows machines, and is almost always a false positive. Commercial virus scanning programs are often confused by the structure of Go binaries, which they don't see as often as those compiled from other languages.

If you've just installed the Go distribution and the system reports it is infected, that's certainly a mistake. To be really thorough, you can verify the download by comparing the checksum with those on the downloads page.

In any case, if you believe the report is in error, please report a bug to the supplier of your virus scanner. Maybe in time virus scanners can learn to understand Go programs.

### 为什么我的病毒扫描软件认为我的围棋发行版或编译的二进制文件被感染了？
这种情况很常见，特别是在Windows机器上，而且几乎总是假阳性。商业病毒扫描程序经常被围棋二进制文件的结构所迷惑，它们并不像那些从其他语言编译的二进制文件那样经常看到。

如果你刚刚安装了围棋发行版，系统却报告被感染了，那肯定是个错误。要想真正彻底，你可以通过比较下载页面上的校验和来验证下载。

无论如何，如果你认为报告有误，请向病毒扫描器的供应商报告错误。也许到时候病毒扫描器可以学会理解围棋程序。

## Performance（性能）
### Why does Go perform badly on benchmark X?
One of Go's design goals is to approach the performance of C for comparable programs, yet on some benchmarks it does quite poorly, including several in golang.org/x/exp/shootout. The slowest depend on libraries for which versions of comparable performance are not available in Go. For instance, pidigits.go depends on a multi-precision math package, and the C versions, unlike Go's, use GMP (which is written in optimized assembler). Benchmarks that depend on regular expressions (regex-dna.go, for instance) are essentially comparing Go's native regexp package to mature, highly optimized regular expression libraries like PCRE.

Benchmark games are won by extensive tuning and the Go versions of most of the benchmarks need attention. If you measure comparable C and Go programs (reverse-complement.go is one example), you'll see the two languages are much closer in raw performance than this suite would indicate.

Still, there is room for improvement. The compilers are good but could be better, many libraries need major performance work, and the garbage collector isn't fast enough yet. (Even if it were, taking care not to generate unnecessary garbage can have a huge effect.)

In any case, Go can often be very competitive. There has been significant improvement in the performance of many programs as the language and tools have developed. See the blog post about profiling Go programs for an informative example.

### 为什么Go在基准X上表现不好？
Go的设计目标之一是在可比程序上接近C语言的性能，然而在一些基准上，它的表现相当糟糕，包括golang.org/x/exp/shootout中的几个基准。最慢的取决于一些库，对于这些库，Go中没有性能相当的版本。例如，pidigits.go依赖于一个多精度数学包，而C版本与Go的不同，使用GMP（用优化的汇编器编写）。依赖于正则表达式的基准（比如regex-dna.go），本质上是将Go的原生regexp包和成熟的、高度优化的正则表达式库（比如PCRE）进行比较。

基准游戏是靠大量的调优取胜的，大部分基准的Go版本都需要注意。如果你测量可比较的C和Go程序（reverse-complement.go就是一个例子），你会发现这两种语言的原始性能比这个套件所显示的要接近得多。

不过，还是有改进的空间。编译器是好的，但还可以更好，许多库需要进行重大的性能改进，垃圾收集器还不够快。即使是这样，注意不要产生不必要的垃圾也会有很大的影响）。

在任何情况下，围棋的竞争往往会非常激烈。随着语言和工具的发展，许多程序的性能有了显著的提高。请参阅关于剖析Go程序的博文，了解一个翔实的例子。

## Changes from C（C的变化）
### Why is the syntax so different from C?
Other than declaration syntax, the differences are not major and stem from two desires. First, the syntax should feel light, without too many mandatory keywords, repetition, or arcana. Second, the language has been designed to be easy to analyze and can be parsed without a symbol table. This makes it much easier to build tools such as debuggers, dependency analyzers, automated documentation extractors, IDE plug-ins, and so on. C and its descendants are notoriously difficult in this regard.

## #为什么语法与C语言如此不同？
除了声明语法之外，其他的差异并不大，而是来自于两个愿望：第一，语法应该让人感觉轻盈，没有太多强制性的关键字、重复或奥义。首先，语法应该让人感觉轻巧，没有太多强制性的关键字、重复或奥义。第二，该语言被设计成易于分析，可以在没有符号表的情况下进行解析。这使得构建调试器、依赖分析器、自动文档提取器、IDE插件等工具更加容易。C和它的后代在这方面是出了名的困难。


### Why are declarations backwards?
They're only backwards if you're used to C. In C, the notion is that a variable is declared like an expression denoting its type, which is a nice idea, but the type and expression grammars don't mix very well and the results can be confusing; consider function pointers. Go mostly separates expression and type syntax and that simplifies things (using prefix * for pointers is an exception that proves the rule). In C, the declaration

    int* a, b;
declares a to be a pointer but not b; in Go

    var a, b *int
declares both to be pointers. This is clearer and more regular. Also, the := short declaration form argues that a full variable declaration should present the same order as := so

    var a uint64 = 1
has the same effect as

    a := uint64(1)
Parsing is also simplified by having a distinct grammar for types that is not just the expression grammar; keywords such as func and chan keep things clear.

See the article about Go's Declaration Syntax for more details.

### 为什么声明是落后的？
只有当你习惯于C语言时，它们才是落后的。在C语言中，概念是一个变量的声明就像一个表示其类型的表达式一样，这是一个很好的想法，但是类型和表达式语法并没有很好的混合，结果可能会很混乱；考虑一下函数指针。Go大多将表达式和类型语法分开，这就简化了事情（对指针使用前缀*是一个证明规则的例外）。在C语言中，声明

    int* a, b.声明a是一个指针，但不是b。
声明a是一个指针，而不是b；在Go中。

    var a, b *int
声明两者都是指针。这样更清晰、更规范。另外，:=短声明形式认为，完整的变量声明应该呈现与:=相同的顺序，所以

    var a uint64 = 1
效果与

    a := uint64(1)
通过为类型制定一个独特的语法，而不仅仅是表达式语法，解析也得到了简化；关键字（如 func 和 chan）使事情变得清晰。

更多细节请参见关于Go的声明语法的文章。

### Why is there no pointer arithmetic?
Safety. Without pointer arithmetic it's possible to create a language that can never derive an illegal address that succeeds incorrectly. Compiler and hardware technology have advanced to the point where a loop using array indices can be as efficient as a loop using pointer arithmetic. Also, the lack of pointer arithmetic can simplify the implementation of the garbage collector.

### 为什么没有指针运算？
安全性。如果没有指针运算，就有可能创建一个永远无法派生出非法地址的语言，从而错误地成功。编译器和硬件技术已经发展到使用数组索引的循环可以和使用指针算术的循环一样高效。另外，缺少指针运算可以简化垃圾收集器的实现。

### Why are ++ and -- statements and not expressions? And why postfix, not prefix?
Without pointer arithmetic, the convenience value of pre- and postfix increment operators drops. By removing them from the expression hierarchy altogether, expression syntax is simplified and the messy issues around order of evaluation of ++ and -- (consider f(i++) and p[i] = q[++i]) are eliminated as well. The simplification is significant. As for postfix vs. prefix, either would work fine but the postfix version is more traditional; insistence on prefix arose with the STL, a library for a language whose name contains, ironically, a postfix increment.

### 为什么是++和--语句而不是表达式？还有为什么是后缀，而不是前缀？
如果没有指针运算，前缀和后缀增量运算符的便利价值就会下降。通过将它们从表达式层次结构中完全移除，表达式语法得到了简化，围绕++和--的计算顺序的混乱问题（考虑f(i++)和p[i] = q[++i]）也被消除了。简化的效果是显著的。至于postfix和prefix，两种都可以，但postfix版本更传统；坚持prefix是在STL中产生的，STL是一个语言库，它的名字包含了一个具有讽刺意味的postfix增量。

### Why are there braces but no semicolons? And why can't I put the opening brace on the next line?
Go uses brace brackets for statement grouping, a syntax familiar to programmers who have worked with any language in the C family. Semicolons, however, are for parsers, not for people, and we wanted to eliminate them as much as possible. To achieve this goal, Go borrows a trick from BCPL: the semicolons that separate statements are in the formal grammar but are injected automatically, without lookahead, by the lexer at the end of any line that could be the end of a statement. This works very well in practice but has the effect that it forces a brace style. For instance, the opening brace of a function cannot appear on a line by itself.

Some have argued that the lexer should do lookahead to permit the brace to live on the next line. We disagree. Since Go code is meant to be formatted automatically by gofmt, some style must be chosen. That style may differ from what you've used in C or Java, but Go is a different language and gofmt's style is as good as any other. More important—much more important—the advantages of a single, programmatically mandated format for all Go programs greatly outweigh any perceived disadvantages of the particular style. Note too that Go's style means that an interactive implementation of Go can use the standard syntax one line at a time without special rules.

### 为什么有大括号而没有分号？为什么我不能把开头的大括号放在下一行？
Go 使用大括号来进行语句分组，这种语法对于使用过任何 C 语言家族的程序员来说都很熟悉。然而，分号是给解析器用的，不是给人用的，我们希望尽可能地消除分号。为了实现这个目标，Go 借用了 BCPL 的一个技巧：分隔语句的分号在形式语法中，但在任何一行可能是语句末尾的地方，都会被词典器自动注入，不需要看头。这在实践中效果非常好，但有一个效果，那就是它强行采用了括号样式。例如，一个函数的开头括号不能单独出现在一行上。

有些人认为，词典应该做lookahead，以允许括号活在下一行。我们不同意这种观点。由于 Go 代码是由 gofmt 自动格式化的，所以必须选择一些样式。这种风格可能与你在C或Java中使用的不同，但Go是一种不同的语言，gofmt的风格与其他任何语言一样好。更重要的是--更重要的是--对所有的Go程序采用单一的、程序规定的格式的优势大大超过了任何认为的特定风格的缺点。还要注意的是，Go的风格意味着Go的交互式实现可以一次一行地使用标准语法，而不需要特殊的规则。

### Why do garbage collection? Won't it be too expensive?
One of the biggest sources of bookkeeping in systems programs is managing the lifetimes of allocated objects. In languages such as C in which it is done manually, it can consume a significant amount of programmer time and is often the cause of pernicious bugs. Even in languages like C++ or Rust that provide mechanisms to assist, those mechanisms can have a significant effect on the design of the software, often adding programming overhead of its own. We felt it was critical to eliminate such programmer overheads, and advances in garbage collection technology in the last few years gave us confidence that it could be implemented cheaply enough, and with low enough latency, that it could be a viable approach for networked systems.

Much of the difficulty of concurrent programming has its roots in the object lifetime problem: as objects get passed among threads it becomes cumbersome to guarantee they become freed safely. Automatic garbage collection makes concurrent code far easier to write. Of course, implementing garbage collection in a concurrent environment is itself a challenge, but meeting it once rather than in every program helps everyone.

Finally, concurrency aside, garbage collection makes interfaces simpler because they don't need to specify how memory is managed across them.

This is not to say that the recent work in languages like Rust that bring new ideas to the problem of managing resources is misguided; we encourage this work and are excited to see how it evolves. But Go takes a more traditional approach by addressing object lifetimes through garbage collection, and garbage collection alone.

The current implementation is a mark-and-sweep collector. If the machine is a multiprocessor, the collector runs on a separate CPU core in parallel with the main program. Major work on the collector in recent years has reduced pause times often to the sub-millisecond range, even for large heaps, all but eliminating one of the major objections to garbage collection in networked servers. Work continues to refine the algorithm, reduce overhead and latency further, and to explore new approaches. The 2018 ISMM keynote by Rick Hudson of the Go team describes the progress so far and suggests some future approaches.

On the topic of performance, keep in mind that Go gives the programmer considerable control over memory layout and allocation, much more than is typical in garbage-collected languages. A careful programmer can reduce the garbage collection overhead dramatically by using the language well; see the article about profiling Go programs for a worked example, including a demonstration of Go's profiling tools.

### 为什么要做垃圾回收？不会太贵吗？
系统程序中最大的记账来源之一是管理分配对象的寿命。在诸如C语言中，它是手动完成的，它可以消耗大量的程序员时间，并且经常是恶性错误的原因。即使在C++或Rust这样提供机制辅助的语言中，这些机制也会对软件的设计产生重大影响，往往会增加自身的编程开销。我们认为消除这样的程序员开销是至关重要的，而过去几年垃圾收集技术的进步让我们有信心可以足够便宜地实现垃圾收集，而且延迟也足够低，它可以成为网络系统的可行方法。

并发编程的大部分困难都源于对象寿命问题：当对象在线程之间传递时，保证它们安全地被释放就变得很麻烦。自动垃圾收集使并发代码的编写变得更加容易。当然，在并发环境中实现垃圾收集本身就是一个挑战，但满足一次而不是在每个程序中都能帮助大家。

最后，抛开并发性不谈，垃圾收集使接口变得更简单，因为它们不需要指定如何跨接口管理内存。

这并不是说最近在Rust等语言中的工作为管理资源的问题带来了新的想法是错误的；我们鼓励这项工作，并很高兴看到它的发展。但是Go采用了一种更传统的方法，通过垃圾收集，而且仅仅是垃圾收集来解决对象寿命问题。

目前的实现是一个标记和扫除的收集器。如果机器是多处理器，收集器就与主程序并行运行在一个单独的CPU核上。最近几年在收集器上的主要工作已经将暂停时间减少到了亚毫秒的范围，即使是大堆，也是如此，所有这些都消除了网络服务器中对垃圾收集的主要反对意见之一。我们还在继续努力完善算法，进一步降低开销和延迟，并探索新的方法。2018年ISMM主题演讲由Go团队的Rick Hudson描述了目前的进展，并提出了一些未来的方法。

关于性能的话题，请记住，Go给了程序员相当大的内存布局和分配控制权，比典型的垃圾收集语言要多得多。细心的程序员可以通过使用好这门语言来大幅降低垃圾收集的开销；请看关于Go程序剖析的文章，了解一个工作实例，包括Go的剖析工具的演示。