| WARNING: This tool is created for entertainment and fun purposes only. It should not be used to crack a real life locks! |
| --- |

[Combination locks](https://en.wikipedia.org/wiki/Combination_lock) usually have C = 10^N (where N is a number of rotating discs) combinations.
In case of average 3 disks lock one would have to try up to 1000 combinations to open it.

But what if we take human factor into account and **assume** when owner scrambles the lock, the **wrong** (or closed) combinations are usually within ± range.  
With this assumption we can potentially learn a code by M numbers of the **wrong** combinations where M is much smaller than C.

Example:
Suppose we have a lock with the passcode 309.
We open/close it multiple times and write down resulting combinations:
![2020-10-03 14 21 55](https://user-images.githubusercontent.com/4749052/94992735-de198080-0583-11eb-8855-58e9e09df0bc.jpg)

Now let's feed it to the tool:
```
$ go run github.com/leoleovich/glock 527 117 581 117
Number 1 is probably 3 (4/4). If not try: 6 (2/4) or 4 (2/4)
Number 2 is probably 0 (4/4). If not try: 9 (3/4) or 2 (3/4)
Number 3 is probably 9 (4/4). If not try: 8 (3/4) or 6 (3/4)
```

`Glock` operates on intersections of sets which means the more diverse numbers we get the faster and more reliable the result is.
