# gosha256

  A sha256 implementation written in Go.

# **NOT FOR PRODUCTION USE**

  You probably want to use the standard Go crypto/sha256 package instead.

  Written by @jwatson0 for my own exploration of the SHA2 algorithm. Primarily the [FIPS 180-2 spec]( https://csrc.nist.gov/csrc/media/publications/fips/180/2/archive/2002-08-01/documents/fips180-2.pdf) was used with very little examination of other implementations or comparison to the standard Go crypto/sha256 package (yet).

### Installing

  ```
  import "github.com/jwatson0/go/gosha256/sha2"
  ```

## Running the tests

  ```
  go get github.com/jwatson0/go/gosha256
  cd $GOPATH/github.com/jwatson0/go/gosha256/sha2
  go test
  ```

## Caveats

  - There are currently a few test cases that are failing, but I haven't tracked down the issue.

  - No optimizations for speed or benchmarks have been done (yet).

  - This implementation works on whole byte boundaries.  Lengths of bits not divisible by 8 are supported in the spec, but not this code (yet).

## Authors

  * **Jason Watson** - [jwatson0](https://github.com/jwatson0)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
