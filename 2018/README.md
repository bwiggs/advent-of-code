# Day 1: Chronal Calibration

**GoLang has GOTO!**: Fun to use ``GOTO`` to solve a problem. I don't recommend it and it's probably not much cleaner than rewriting using while style loop and manual index management.

### Part 1: Starting with a frequency of zero, what is the resulting frequency after all of the changes in frequency have been applied?

Read Integers from a file and find the sum.

### Part 2: What is the first frequency your device reaches twice?

Continuously run through the numbers from the file until you hit a repeat number.

# Day 2: Inventory Management System

### Part 1: What is the checksum for your list of box IDs?

### Part 2: What letters are common between the two correct box IDs?

# Day 3: No Matter How You Slice It

### Part 1: How many square inches of fabric are within two or more claims?

### Part 2: What is the ID of the only claim that doesn't overlap?

# Day 4: Repose Record

### Part 1: Strategy 1: What is the ID of the guard you chose multiplied by the minute you chose?

### Part 2: Strategy 2: which guard is most frequently asleep on the same minute? What is the ID of the guard you chose multiplied by the minute you chose?

# Day 5: Alchemical Reduction

**Text Processing and Stacks**: How many times am I not going to reach for a stack when solving a problem related to text processing?! I solved this originally using loops with some nasty index tracking. This worked but wasn't anywhere near as clean as using a simple stack. **Whenever I see text processing problem, I should think of reaching for a stack based solution.**

**Use XOR to invert case:** I also used some fancy bitwise operation to compare uppercase to lower case ascii letters. Turns out if you XOR 32, you can convert case. Pretty brilliant on behalf of the creators of ASCII.

### Part 1: How many units remain after fully reacting the polymer you scanned?

Throw each char on a stack and wait for the sibling character to arrive.

### Part 2: What is the length of the shortest polymer you can produce by removing all units of exactly one type and fully reacting the result?



# Day 6: Chronal Coordinates

>  Part 1: What is the size of the largest area that isn't infinite?

>  Part 2: What is the size of the region containing all locations which have a total distance to all given coordinates of less than 10000?


# Day 7 - The Sum of Its Parts (Topological Sort)

### Part 1: In what order should the steps in your instructions be completed?

### Part 2: With 5 workers and the 60+ second step durations described above, how long will it take to complete all of the steps?

Day 7 messed me up. My brain really struggled to come up with a clean solution, which then lead to a lot of mental frustration. The way I was going to solve it felt really dirty, so I cheated and looked at how some others were solving it. That's when I learned about Topological Sorts, something I'd never come across before. I was on the right track with a graph solution, but if you used a Topological sort with Kahns Algorithm things a bit easier to manage.

I also found out that I had misread the requirements! I thought that all child steps had to be completed first. If you reread the instructions: "Then, even though F was available earlier, steps B and D are now also available, and B is the first alphabetically of the three." it turns out you just need to process the next alphabetical step in the list. THAT MADE THINGS WAY EASIER. Read instructions carefully!

After researching topological sorting and Kahns Algorithm. I was able to put together a working solution.

Thoughts on GoLang:

1. **GoLang is verbose**. Setting up maps and filling them requires a lot of nil checking. Nothing is free.
2. **List processesing isnt as easy** as it is in functional languages. Being able to slice and dice arrays is really convinent, and probably one of the reasons why a good amount of AoC participants are using Python to solve these problems.

# Day 8: Memory Maneuver

### Part 1: What is the sum of all metadata entries?

