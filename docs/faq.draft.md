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
