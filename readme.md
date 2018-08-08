# GoLang vs Python for pie estimation using Basel

Probability that two random numbers are coprime is (6/π^2). One can estimate the value of Pi by using random integers.

I wrote a program in python to estimate the value of Pie using this method, mostly for fun. This is not the easiest nor the best way to estimate π but is just an interesting way to do so.

```
$ time python pi.py
0 3.14331776888
1 3.14068121226
2 3.13908183774
3 3.14396499755
4 3.13856643148
5 3.14614261795
6 3.14355072515
7 3.13594180281
8 3.14590908502
9 3.15330351514
python pi.py  82.47s user 0.46s system 98% cpu 1:24.04 total
```

Once I did that I was curious to check the same thing using GoLang

```
$ time go run pi.go
3.1369859687785224
3.1369087972027527
3.1430747045726997
3.1394069003233436
3.1366258834777594
3.1416525514018883
3.1455616426154176
3.1482364956210604
3.1452504524902403
3.148210493362175
go run pi.go  0.51s user 0.17s system 85% cpu 0.784 total
```

Turns out GoLang is insanely faster compared to Python as far as this problem is concerned. Almost 100 times faster!

In this repository, you will find both the programs.
