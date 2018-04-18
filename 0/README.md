# Dive 0 – Getting our hair wet

1. Look at this file.  Package? What's that?  Import? *cooing noise*
2. What's a fmt?  "format".  Take a peek inside.
  2.1  Interesting how it's able to add a space between hello and world automatically.  How does it manage that?
  2.2  Looks like the argument to Println is an `...interface{}`.  What's that?
3. Cool, we have empty interfaces and set enumerators.
4. We also see that os.Stdout is used as a "Writer" interface, whatever that is.  Let's look at Stdout.write()...
5. Cool, file structs and whatnot.  Let's find a syscall... would you look at that! syscall.Write()
6. Let's see what syscall package looks like, shouldn't be anything too surprising... right?

    func Write(fd int, p []byte) (n int, err error) {
      if race.Enabled {
        race.ReleaseMerge(unsafe.Pointer(&ioSync))
      }
      n, err = write(fd, p)
      if race.Enabled && n > 0 {
        race.ReadRange(unsafe.Pointer(&p[0]), n)
      }
      if msanenabled && n > 0 {
        msanRead(unsafe.Pointer(&p[0]), n)
      }
      return
    }

7. Okay, so `write` here points to `zsyscall_darwin_386.go`, I'm going to guess that's specific to my machine – might be different on yours.
8. But what is this funny business with race.Enabled and msaenabled?  What do those mean?  I step into where enabled is defined, and I see a bunch of empty functions.  What in the world?
9. Seems there's a race.go and a norace.go.  Strange.  Looked it up, seems to be that they're there on purpose so that if you compile without the -race flag, they're just no-ops and not included.  Fascinating.
10. If you look in race.go, there are a bunch of references to a 'runtime' package.  Sounds interesting, let's explore.
11. Oh no, seems that my editor doesn't even know where the `runtime` package is! We must be very deep in the weeds now.  Open up the old file explorer and search.
12. Ha! Found it.  Next to a million other files – some of them assembly, some of them Go.  Ladies and gentlemen, welcome to the runtime.  Says at the top that this file is present in compiled code IFF compiled with `-race`.  Cool.

That's deep enough, don't you think?  We've finally hit the code where variable names begin with multiple underscores.  Let's keep in mind how deep we had to go to find that.

When you're debugging a problem, 99.9% of the time, you will not end up looking at this code; but this code is the heart of the language's power.  The API made here (in the runtime) are leveraged everywhere else, yielding code like that in the in the `os` or `fmt` package.  That is, code that is very human-readable.  The runtime died for your sins!  And you get cool things like race detectors as a free lunch because of it :)

