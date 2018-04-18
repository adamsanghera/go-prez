# Error-handling with defer, panic, recover

A few takeaways here:

1. defer() is your friend.  You can use it for a lot of things.  Here, we use it to recover from panics.  It's good to understand what exactly a panic is first.

Let's say we have a callstack that looks like this:

     ____
    | f3 |  panic begins
    |____|
    | f2 |
    |____|
    | f1 |  deferred recover
    |____|

In this scenario, f3 encounters some error that the developer of f3 decided was beyond his pay grade.

f1 is paid a lot though, so it's her job to handle the kinds of things that f3 panics about.  Accordingly, the developer of f1 has deferred a recover statement, which handles the error and restarts f1, which again calls f2 (who calls f3, and so on...).

In reality, the call stack looks something like this:

     ________
    | f3     |  panic begins
    |________|
    | f2     |
    |________|
    | f1     |
    |________|
    |defer fn|
    |________|

The panic blows through the call stack, `aborting` every function until `recover` is called.  It blows through f3, f2, f1, and is caught by the the `defer fn`.  `defer fn` `recover`s the panic, and rebuilds the call-stack.  This pattern is called `defer, panic, recover`.  You can read more about it <a href='https://blog.golang.org/defer-panic-and-recover'>here</a>.

The idea is that this model is more flexible and legible than try, catch recover.  It is idiomatic in Go to return error values that you should expect the user (another developer) of your package to handle.  However, some errors are so rare or difficult to recover from, that panicking is appropriate.  Thus, we have "classes" of errors, which are "caught" and responded to separately (in the return values, and recover() respectively).  This is useful.

Sometimes, it is also useful to panic and defer a recover, because it makes the code more elegant.  For example, if you're downloading a large file or performing some long-running computation, it is useful to panic on failure, and then have a manager-function induce panic in all of the workers before saving their results in a deferred function.  Neat-o.