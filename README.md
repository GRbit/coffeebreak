### Discliamer

The thing is just for fun, it's useless and I take no responsibility for anything.

# coffebreak

Hello fellow gophers!

This package is aimed to give all of you some time to drink a cup of coffee while your code is... compiling!

No more mockery from developers of other compiled languages! Now you too can take a well-deserved break to get your thoughts in order.

## How to

Just import the package like this:

```
import _ "github.com/grbit/coffebreak/break/drink1024Kcups"
```

The number correspond to number of function that will be in the imported package. Unfortunately, go comipler is so fast, that even compiling 1 million of functions takes a few minutes.

### Pros

* Binaries have the same size
* No side effects
* You have time to drink coffee

### Cons

* It's a one time trick. Since go build caches packages, you'll need to run `go clean -cache`
* It weights ~600mb. So, be cautious, don't use it in CICD
