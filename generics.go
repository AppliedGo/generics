/*
<!--
Copyright (c) 2016 Christoph Berger. Some rights reserved.
Use of this text is governed by a Creative Commons Attribution Non-Commercial
Share-Alike License that can be found in the LICENSE.txt file.

The source code contained in this file may import third-party source code
whose licenses are provided in the respective license files.
-->

<!--
NOTE: The comments in this file are NOT godoc compliant. This is not an oversight.

Comments and code in this file are used for describing and explaining a particular topic to the reader. While this file is a syntactically valid Go source file, its main purpose is to get converted into a blog article. The comments were created for learning and not for code documentation.
-->

+++
title = "Who needs generics? Use ... instead!"
description = "The Go language has no generics. This article is a survey of techniques that can be used instead."
author = "Christoph Berger"
email = "chris@appliedgo.net"
date = "2016-07-07"
publishdate = "2016-07-07"
domains = ["Patterns and Paradigms"]
tags = ["generics", "interface", "reflection", "code generation"]
categories = ["Essay"]
+++

What are generics? Why are they considered useful? Why does Go have no generics? What shall Gophers use instead?

This article examines the nature of generics, and surveys various techniques that can be used to work around the absence of this programming paradigm.

<!--more-->

## First, an important note

*The question about generics in Go is years old, and has been discussed up and down and forth and back across the Go forums, newsgroups, and email lists. It is NOT the goal of this article to spark this discussion once again. I think that all that can be said, has been said.*

*Fact is, the concept of generics does not exist in Go.*

*This article solely focuses on alternate ways of achieving some of the goals that other languages try to solve with generics.*

Let's start by briefly looking at the motivation for generics as well as the possible downsides.


## The problem

Many data structures and algorithms are applicable to a range of data types. A sorted tree, for example, could be defined to hold elements of type `int`, or `map[string]string`, or some `struct` type. A sorting algorithm is able to sort any type whose elements are comparable to each other. So if you need, say, a sorted tree for strings, you could sit down and write one. Easy enough!

But what if you need a sorted tree for many different types? The tree datatype comes with a couple of methods like `Insert`, `Find`, `Delete`, etc. If there are *N* tree methods, and *M* element types to implement the tree for, you would end up with *N x M* methods to implement! Sounds like days of tedious, repetitive work.

Can't we let the compiler do that?


## Enter generics

To address this problem, many programming languages have a concept called *'generics'*. The point is to write the code only once, in an abstract form, with opaque type placeholders instead of real types, as in this Java example:

```Java
// Java pseudo code

public class SortedTree<E implements Comparable<E>> {
	void insert(E comparableSortTreeElement) {
		//...
	}
}
```

`E` is a type placeholder; it can appear in class or interface definitions, as well as in function parameter lists. `SortedTree` is used by substituting a real type for the placeholder:

```Java
// Java pseudo code (again)

SortedTree<String> sortedTreeOfStrings = new SortedTree<String>();
sortedTreeOfStrings.insert("abcde");
```

These lines instantiate a sorted tree with string elements. The `insert` method then only accepts strings, and the sort algorithm uses string comparison methods. Had we used `Integer` instead, like in `SortedTree<Integer>`, the tree would be a sorted tree of integers instead.

To summarize, generics allow to define classes and functions with type parameters. The type parameters can be restricted to a certain subset (e.g., parameter T must be comparable), but otherwise they are unspecific. The result is a code template that can be applied to different types as the need arises.


## The downsides

While generics may come in handy, they also have some strings attached.

1. **Performance:** Turning a generic code template into actual code takes time, either at compile time or at runtime. Or as [Russ Cox stated](http://research.swtch.com/generic) in 2009:

    > The generic dilemma is this: *do you want slow programmers, slow compilers and bloated binaries, or slow execution times?*

	(The 'slow programmers' part refers to having no generics at all, nor any suitable substitute.)

2. **Complexity**: Generics are not complex *per se*, but they can become complex when integrated with other language features such as inheritance, or when devlopers start creating nested generics like,

	```C#
	// C#
	List<Dictionary<string<IEnumerable<HttpRequest>>>>
	```

	(Thanks to [Jonathan Oliver](http://blog.jonathanoliver.com/golang-has-generics/) for this nice example), or even seemingly recursive inheritance like,

	```Java
	// Java
	public abstract class Enum<E extends Enum<E>>
	```

	(taken from [this Java Generics FAQ](http://www.angelikalanger.com/GenericsFAQ/FAQSections/TypeParameters.html#FAQ106) by Angelika Langer).


## Go has no generics

As we all know, Go has deliberately been designed with simplicity in mind, and generics are considered to add complexity to a language (see the previous section). So along with inheritance, polymorphism, and some other features of the 'state of the art' languages at that time, generics were also left off from the list.


## But what if a project just seems to cry for generics?

At times you might come across a problem where generics would seem to be very handy, if not downright indispensable. What can you do then?

Luckily, there are a couple of options.


### 1. Review the requirements

Step back and revisit the requirements. Review the technical or functional specification (you should have one). Do the specs really demand the use of generics? Consider that while other languages may support a design that is based on type systems, Go's philosophy is different:

> If C++ and Java are about type hierarchies and the taxonomy of types, Go is about composition. [(Rob Pike)](https://commandcenter.blogspot.de/2012/06/less-is-exponentially-more.html)

So think of the paradigms that come with Go--most notably composition and embedding--and verify if any of these would help approaching the problem in a way more natural to Go.



### 2. Consider copy & paste

This may sound like a foolish advice (and if applied improperly, it surely is), but do not hastily reject it.

Here is the point:

If you plan to implement some generic library, then of course there is no way around generics. **But in all other cases, you might not necessarily need any generic abstractions at all.**

Every time you think that you need to create a generic object, do a quick Litmus test: Ask yourself, "How many times would I ever have to instantiate this generic object?" In other words, is it worth to construct a generic object when there may only be one or two actual implementations of it? In this case, the generic object is nothing but a result from over-abstracting things.

(At this point, I cannot help but thinking of Joel Spolsky's witty article on [Architecture Astronauts](http://www.joelonsoftware.com/articles/fog0000000018.html). Warning: the article is from 2001. Expect a couple of references to outdated software concepts of which you never may have heard if you are young enough.)

An excellent real-life example can be found right in the standard library. The two packages `strings` and `bytes` have almost identical API's, yet no generics were used in the making of these packages (or so I strongly guess).

Remember that this article tries to remain neutral and takes no stand on generics. I don't want to indicate that copy&paste is better than generics. I only want to encourage you to seek out other options even if they look strange at a first glance.


### 3. Use interfaces

Interfaces define behavior without requiring any implementation details. This is ideal for defining 'generic' behavior:

* Find a set of basic operations that your generic algorithm or data container can use to process the data.
* Define an `interface` containing these operations.
* To 'instantiate' your generic entity on a given data type, implement the interface methods for that type.

The `sort` package is an example for this technique. It contains a sort interface (appropriately called `sort.Interface`) that declares three methods: `Len()`, `Less()`, and `Swap()`.

```go
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
```

By just implementing these three methods for a data container (say, a slice of structs), `sort.Sort()` can be
applied to any kind of data, as long as the data has a reasonable definition of 'less than'.

The code of `sort.Sort()` does not know anything about the data it sorts, and actually it does not have to. It simply relies on the three interface methods `Len`, `Less`, and `Swap`.


*/

// Imports and globals (also for the examples from the sections further down)
package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
)

// point is a 2-dimensional point with Cartesian coordinates.
type point struct {
	x, y float64
}

// Dist is the point's distance to (0,0).
func (p point) Dist() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y) // good old Pythagoras
}

// points is our point container. We want to implement the generic sort.Sort() for this.
type points []point

// Len is just point's len().
func (p points) Len() int {
	return len(p)
}

// Less returns true if the point at index i is closer to (0,0) than the point at index j.
func (p points) Less(i, j int) bool {
	return p[i].Dist() < p[j].Dist()
}

// Swap is a simple slice swap.
func (p points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Let's sort some points.
func sortExample() {
	p := points{{7, 9}, {-3, 3}, {5, -8}, {0, 4}}
	sort.Sort(p)
	fmt.Println("sortExample: ", p)
}

/*

### 4. Use type assertions

Generic container types usually do not need to care much about the actual type of their contents. (Except maybe for basic operations like sorting, but we already know a solution for that.) Hence the value can be stored in the container as a 'type without properties'. Such a type is already built into Go: It is the empty interface, declared as `interface{}`. This interface has no particular behavior, hence objects with any behavior satisfy this interface.

It is quite easy to build a container type based on `interface{}`. We only need a way to recover the actual data type when reading elements from the container. For that purpose, Go has type assertions. Here is an example:

*/

// list is a generic container, accepting anything.
type list []interface{}

// Put adds an element to the container.
func (l *list) Put(elem interface{}) {
	*l = append(*l, elem)
}

// Get gets an element from the container.
func (l *list) Get() interface{} {
	elem := (*l)[0]
	*l = (*l)[1:]
	return elem
}

// The calling code does the type assertion when retrieving an element.
func assertExample() {
	intContainer := &list{}
	intContainer.Put(7)
	intContainer.Put(42)
	elem, ok := intContainer.Get().(int) // assert that the actual type is int
	if !ok {
		fmt.Println("Unable to read an int from intContainer")
	}
	fmt.Printf("assertExample: %d (%T)\n", elem, elem)
}

/*

This looks so straightforward that we might easily forget the cons of this technique. We give up compile-time type checking here, exposing the application to increased risk of type-related failure at runtime. Also, conversions to and from interfaces come with a cost.


### 5. Use reflection

Remember the *generic dilemma* from the 'downsides' section? One option was "slow execution times". The technique discussed below is a manifestation of this option. Yes, I am talking about reflection.

Reflection allows a program to work with objects whose types are not known at compile time. This sounds like a good fit for solving the generics problem, so let's give it a try.

*/

// Container can hold anything that satisfies the empty interface.
type Container []interface{}

func NewContainer(t reflect.Type) *Container {
	return nil
}

// Set sets the element at index i to the passed-in value.
func (c *Container) Set(i int, val interface{}) {
}

// Get gets the element at index i.
func (c *Container) Get(i int) interface{} {
	return (*c)[i]
}

//
func reflectExample() {
	//c := NewContainer(float64)
	//c.Put(0, 1.2345)
	//f := c.Get(0).ValueOf(float64)
	//fmt.Println(f)
}

/*
### 6. Use a code generator



*/

func main() {
	// The interface approach:
	sortExample()

	// The type assertion approach:
	assertExample()

	// The reflection approach:

	// The code generator approach:
}

/*
## Conclusion





# A brief announcement about this blog

While this blog is only a few weeks old, early readers may already have noticed that the articles were always published quite regularly, once a week on Thursdays or Fridays. And from the beginning I strived to provide quality content that has a real benefit to the reader--even if it is just some sort of aha experience.

Now life is usually filled with many duties, and right now they start piling up again, forcing me to either publish smaller articles, or to publish less often. As I don't want to compromise quality, I decided to change the schedule to bi-weekly. So starting this Thursday, new articles will appear every fortnight, (Although I would not rule out that once in a while I may return to a weekly schedule; for example, when posting an article series.)

*/
