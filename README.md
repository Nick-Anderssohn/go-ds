# go-ds
Useful data structures for go. First one is just a heap package that is
more convenient and intuitive than go's stdlib heap package. As long as the value
you are storing in the heap implements `comparable`, you are good and
dont need to do anything else. Plus this leverages generics instead of just
using empty interfaces.