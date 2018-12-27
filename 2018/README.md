# Day 1: Chronal Calibration

> Part 1: Starting with a frequency of zero, what is the resulting frequency after all of the changes in frequency have been applied?

Read Integers from a file and find the sum.

> Part 2: What is the first frequency your device reaches twice?

Continuously run through the numbers from the file until you hit a repeat number.

This was fun because I got to use a goto. *I don't recommend using this.*

# Day 2: Inventory Management System

> Part 1: What is the checksum for your list of box IDs?

> Part 2: What letters are common between the two correct box IDs?

# Day 3: No Matter How You Slice It

> Part 1: How many square inches of fabric are within two or more claims?

> Part 2: What is the ID of the only claim that doesn't overlap?

# Day 4: Repose Record

> Part 1: What is the ID of the guard you chose multiplied by the minute you chose?

> Part 2: What is the ID of the guard you chose multiplied by the minute you chose?

# Day 5: Alchemical Reduction

> Part 1: How many units remain after fully reacting the polymer you scanned?

> Part 2: What is the length of the shortest polymer you can produce by removing all units of exactly one type and fully reacting the result?

# Day 6: Chronal Coordinates

>  Part 1: What is the size of the largest area that isn't infinite?

>  Part 2: What is the size of the region containing all locations which have a total distance to all given coordinates of less than 10000?


# Day 7 - The Sum of Its Parts (Topological Sort)

> Part 1: In what order should the steps in your instructions be completed?

> Part 2: With 5 workers and the 60+ second step durations described above, how long will it take to complete all of the steps?

Day 7 messed me up. My brain really struggled to solve this and it lead to a lot of mental frustration that I couldn't come up with a clean solution. Not that I couldn't code up a solution, but jsut that the way I was going to solve it felt really dirty.

So I cheated and looked at how some others were solving it. That's when I learned about Topological Sorts. Topo sorts was new to me.

After lookup up Topo sorts and reading other's code I found out that I had misread the requirements. I thought that all child steps had to be completed first. Well friends, if you reread the instructions, you'll see that in reality you just need to process the next alphabetical step in the list, regardless of it's location in the graph. THAT MAKES THINGS WAY EASIER.

Thoughts on GoLang:

1. **GoLang is verbose**. Setting up maps and filling them requires a lot of nil checking. Nothing is free.
2. **List processesing isnt as easy** as it is in functional languages. Being able to slice and dice arrays is really convinent, and probably one of the reasons why a good amount of AoC participants are using Python to solve these problems.

# Day 8: Memory Maneuver

> Part 1: What is the sum of all metadata entries?

