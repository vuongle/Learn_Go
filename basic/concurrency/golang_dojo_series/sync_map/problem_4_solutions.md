Periodically, when you are developing a Golang application, there is a need to store some data in memory using a map data structure.
At the same time, the application can concurrently process a large number of requests and use the map in several goroutines concurrently.

If you do not use any mechanisms of synchronization, the code fails with panic in runtime in that case.
It is important to think of concurrent access to the map and data races.

In this article, I want to consider several ways to solve this problem. Golang has built-in synchronization primitives and channels that can be used for that. Depending on the specific case, you can use different approaches:

```
solution1_mutex
solution2_rwmutex
solution3_syncmap
solution4_channels
```

I am going to create the call counter which is a simple component that counts the number of times a method is called and the counters will be stored in a map. To begin with, let's define a simple interface CallCounter that will have two methods.

The method Call will update the counter when the method is called, and the Count method will return the current value for the given method name.
