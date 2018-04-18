# Semantics of Concurrency

We're going to build a program that constructs Nth triangular number in parallel, where the Nth triangular number is defined as follows:

    0+1+2+3+4+... + i + ... + N

To achieve this, we're going to divide N into M equal-ish (maybe one is smaller than the rest?) pieces.

Then we're going to let the M routines calculate their partial sums, and return the final result.

Things to takeaway:

1. An oft-repeated phrase in the Go community is:

    >Do not communicate by sharing memory; instead, share memory by communicating

1. Here, the wrong approach would be to share one counter, and have 10 routines atomically adding to it.  Instead, we share the 'memory' by independently communicating, and joining the results.

1. This model (`mappers` to extract data, `reducers` to join them) is how MapReduce works.  Guess where the inspiration is from?

1. This model is very powerful, and can be hard to implement in other languages.  It usually involves implementing (or importing) some type of threadsafe queue.  Building those right is really tricky, especially when the underlying language lacks low-level and efficient lock usage.

1. Channels are implemented in the run-time, and exploit the runtime's particular memory-management model to be highly efficient, even when piping very large amounts of data.

1. Mutexes are still very important.  Sometimes, communication isn't a model that makes sense, but mutexes, wait groups, condition variables are.  That is what the `sync` package is for.

Note that in this application, the speedup provided by multiple workers is probably negligible, as tightly-looped arithmetic operations are highly optimized by the compiler.  As such, this example is just a demonstration of a common principle.

However, In workloads where the routines are making disk-intensive or otherwise syscall-tied operations (like network calls), the speedup is dramatic.  Think of downloading a large file with concurrent HTTP Range Headers, making multiple concurrent SQL queries, or crawling multiple API endpoints concurrently.